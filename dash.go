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

type Flags struct {
   Address string
   Audio_Bandwidth int
   Audio_Name string
   Client_ID string
   Info bool
   Private_Key string
   Video_Bandwidth int
}

func (f Flags) DASH(base string, post widevine.Poster) error {
   res, err := client.Redirect(nil).Get(f.Address)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   var media dash.Media
   if err := xml.NewDecoder(res.Body).Decode(&media); err != nil {
      return err
   }
   reps := media.Representations()
   var str stream_DASH
   str.base = res.Request.URL
   str.basename = base
   str.flag = f
   str.post = post
   str.Representations = reps.Audio()
   if err := str.download(f.Audio_Bandwidth); err != nil {
      return err
   }
   str.Representations = reps.Video()
   return str.download(f.Video_Bandwidth)
}

type stream_DASH struct {
   base *url.URL
   basename string
   dash.Representations
   flag Flags
   post widevine.Poster
}

func (s stream_DASH) download(bandwidth int) error {
   if bandwidth <= 0 {
      return nil
   }
   rep := s.Bandwidth(bandwidth)
   if s.flag.Info {
      for _, each := range s.Representations {
         if each.Bandwidth == rep.Bandwidth {
            fmt.Print("!")
         }
         fmt.Println(each)
      }
      return nil
   }
   file, err := os.Create(s.basename + rep.Ext())
   if err != nil {
      return err
   }
   defer file.Close()
   addr, err := s.base.Parse(rep.Initialization())
   if err != nil {
      return err
   }
   res, err := client.Redirect(nil).Get(addr.String())
   if err != nil {
      return err
   }
   defer res.Body.Close()
   media := rep.Media()
   pro := os.Progress_Chunks(file, len(media))
   dec := mp4.New_Decrypt(pro)
   var key []byte
   if rep.ContentProtection != nil {
      private_key, err := os.ReadFile(s.flag.Private_Key)
      if err != nil {
         return err
      }
      client_ID, err := os.ReadFile(s.flag.Client_ID)
      if err != nil {
         return err
      }
      key_ID, err := widevine.Key_ID(rep.ContentProtection.Default_KID)
      if err != nil {
         return err
      }
      mod, err := widevine.New_Module(private_key, client_ID, key_ID)
      if err != nil {
         return err
      }
      keys, err := mod.Post(s.post)
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
      addr, err := s.base.Parse(medium)
      if err != nil {
         return err
      }
      res, err := client.Redirect(nil).Level(0).Get(addr.String())
      if err != nil {
         return err
      }
      pro.Add_Chunk(res.ContentLength)
      if rep.ContentProtection != nil {
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
