package mech

import (
   "encoding/xml"
   "fmt"
   "github.com/89z/mech/widevine"
   "github.com/89z/std/dash"
   "github.com/89z/std/http"
   "github.com/89z/std/mp4"
   "github.com/89z/std/os"
   "io"
)

var client = http.Default_Client

type Flags struct {
   Address string
   Bandwidth_Audio int64
   Bandwidth_Video int64
   Client_ID string
   Info bool
   Private_Key string
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
   str.base = base
   str.flag = f
   str.post = post
   str.Representations = reps.Audio()
   if err := str.download(f.Bandwidth_Audio); err != nil {
      return err
   }
   str.Representations = reps.Video()
   return str.download(f.Bandwidth_Video)
}

type stream_DASH struct {
   base string
   dash.Representations
   flag Flags
   post widevine.Poster
}

func (s stream_DASH) download(bandwidth int64) error {
   if bandwidth <= 0 {
      return nil
   }
   rep := s.Get_Bandwidth(bandwidth)
   if s.flag.Info {
      for _, each := range s.Representations {
         if each.Bandwidth == rep.Bandwidth {
            fmt.Print("!")
         }
         fmt.Println(each)
      }
      return nil
   }
   file, err := os.Create(s.base + rep.Ext())
   if err != nil {
      return err
   }
   defer file.Close()
   res, err := client.Redirect(nil).Get(rep.Initialization())
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
   for _, addr := range media {
      res, err := client.Redirect(nil).Level(0).Get(addr)
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
