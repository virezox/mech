package mech

import (
   "encoding/xml"
   "fmt"
   "github.com/89z/mech/widevine"
   "github.com/89z/rosso/dash"
   "github.com/89z/rosso/http"
   "github.com/89z/rosso/mp4"
   "github.com/89z/rosso/os"
   "io"
   "net/url"
)

var client = http.Default_Client

type Stream struct {
   Client_ID string
   Info bool
   Private_Key string
   Poster widevine.Poster
   Base string
   url *url.URL
}

func (self *Stream) DASH(ref string) (dash.Representations, error) {
   res, err := client.Redirect(nil).Get(ref)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   var pres dash.Presentation
   if err := xml.NewDecoder(res.Body).Decode(&pres); err != nil {
      return nil, err
   }
   self.url = res.Request.URL
   return pres.Representation(), nil
}

func (self Stream) DASH_Get(items dash.Representations, index int) error {
   if self.Info {
      for i, item := range items {
         if i == index {
            fmt.Print("!")
         }
         fmt.Println(item)
      }
      return nil
   }
   item := items[index]
   file, err := os.Create(self.Base + item.Ext())
   if err != nil {
      return err
   }
   defer file.Close()
   ref, err := self.url.Parse(item.Initialization())
   if err != nil {
      return err
   }
   res, err := client.Redirect(nil).Get(ref.String())
   if err != nil {
      return err
   }
   defer res.Body.Close()
   media := item.Media()
   pro := os.Progress_Chunks(file, len(media))
   dec := mp4.New_Decrypt(pro)
   var key []byte
   if item.ContentProtection != nil {
      private_key, err := os.ReadFile(self.Private_Key)
      if err != nil {
         return err
      }
      client_ID, err := os.ReadFile(self.Client_ID)
      if err != nil {
         return err
      }
      key_ID, err := widevine.Key_ID(item.ContentProtection.Default_KID)
      if err != nil {
         return err
      }
      mod, err := widevine.New_Module(private_key, client_ID, key_ID)
      if err != nil {
         return err
      }
      keys, err := mod.Post(self.Poster)
      if err != nil {
         return err
      }
      key = keys.Content().Key
      if err := dec.Init(res.Body); err != nil {
         return err
      }
   } else {
      _, err := io.Copy(pro, res.Body)
      if err != nil {
         return err
      }
   }
   for _, medium := range media {
      ref, err := self.url.Parse(medium)
      if err != nil {
         return err
      }
      res, err := client.Redirect(nil).Level(0).Get(ref.String())
      if err != nil {
         return err
      }
      pro.Add_Chunk(res.ContentLength)
      if item.ContentProtection != nil {
         err = dec.Segment(res.Body, key)
      } else {
         _, err = io.Copy(pro, res.Body)
      }
      if err != nil {
         return err
      }
      if err := res.Body.Close(); err != nil {
         return err
      }
   }
   return nil
}
