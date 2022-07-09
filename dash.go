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

type Flag struct {
   Address string
   Client_ID string
   Info bool
   Name string
   Poster widevine.Poster
   Private_Key string
   base *url.URL
}

func (f *Flag) Presentation() (*dash.Presentation, error) {
   res, err := client.Redirect(nil).Get(f.Address)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   pre := new(dash.Presentation)
   if err := xml.NewDecoder(res.Body).Decode(pre); err != nil {
      return nil, err
   }
   f.base = res.Request.URL
   return pre, nil
}

func (f Flag) DASH(items []dash.Representation, index int) error {
   if f.Info {
      for i, item := range items {
         if i == index {
            fmt.Print("!")
         }
         fmt.Println(item)
      }
      return nil
   }
   item := items[index]
   file, err := os.Create(f.Name + item.Ext())
   if err != nil {
      return err
   }
   defer file.Close()
   addr, err := f.base.Parse(item.Initialization())
   if err != nil {
      return err
   }
   res, err := client.Redirect(nil).Get(addr.String())
   if err != nil {
      return err
   }
   defer res.Body.Close()
   addrs := item.Media()
   pro := os.Progress_Chunks(file, len(addrs))
   dec := mp4.New_Decrypt(pro)
   var key []byte
   if item.ContentProtection != nil {
      private_key, err := os.ReadFile(f.Private_Key)
      if err != nil {
         return err
      }
      client_ID, err := os.ReadFile(f.Client_ID)
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
      keys, err := mod.Post(f.Poster)
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
   for _, raw := range addrs {
      addr, err := f.base.Parse(raw)
      if err != nil {
         return err
      }
      res, err := client.Redirect(nil).Level(0).Get(addr.String())
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
