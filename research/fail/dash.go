package fail

import (
   "encoding/xml"
   "fmt"
   "github.com/89z/mech/widevine"
   "github.com/89z/rosso/dash"
   "github.com/89z/rosso/http"
   "github.com/89z/rosso/mp4"
   "github.com/89z/rosso/os"
)

var client = http.Default_Client

type Stream struct {
   Client_ID string
   Info bool
   Private_Key string
   Poster widevine.Poster
   Base string
   req *http.Request
}

func (s *Stream) DASH(address string) (dash.Representations, error) {
   res, err := client.Redirect(nil).Get(address)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   var pres dash.Presentation
   if err := xml.NewDecoder(res.Body).Decode(&pres); err != nil {
      return nil, err
   }
   s.req = res.Request
   return pres.Representation(), nil
}

func (s Stream) DASH_Get(items dash.Representations, index int) error {
   if s.Info {
      for i, item := range items {
         if i == index {
            fmt.Print("!")
         }
         fmt.Println(item)
      }
      return nil
   }
   item := items[index]
   file, err := os.Create(s.Base + item.Ext())
   if err != nil {
      return err
   }
   defer file.Close()
   s.req.URL, err = s.req.URL.Parse(item.Initialization())
   if err != nil {
      return err
   }
   res, err := client.Redirect(nil).Do(s.req)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   media := item.Media()
   pro := os.Progress_Chunks(file, len(media))
   dec := mp4.New_Decrypt(pro)
   if item.ContentProtection != nil {
      private_key, err := os.ReadFile(s.Private_Key)
      if err != nil {
         return err
      }
      client_ID, err := os.ReadFile(s.Client_ID)
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
      if _, err := mod.Post(s.Poster); err != nil {
         return err
      }
      if err := dec.Init(res.Body); err != nil {
         return err
      }
   }
   return nil
}
