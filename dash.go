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
   Bandwidth_Audio int
   Bandwidth_Video int
   Client_ID string
   Info bool
   Private_Key string
   address string
}

func (f Flags) Decode(base string, post widevine.Poster) error {
   res, err := client.Redirect(nil).Get(f.address)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   var media dash.Media
   if err := xml.NewDecoder(res.Body).Decode(&media); err != nil {
      return err
   }
   reps := media.Representations()
   var str stream
   str.base = base
   str.post = post
   str.Representations = reps.Audio()
   str.bandwidth = f.Bandwidth_Audio
   if err := f.download(str); err != nil {
      return err
   }
   str.Representations = reps.Video()
   str.bandwidth = f.Bandwidth_Video
   return f.download(str)
}

func (f Flags) download(str stream) error {
   if str.bandwidth <= 0 {
      return nil
   }
   rep := str.Get_Bandwidth(str.bandwidth)
   if f.Info {
      for _, each := range str.Representations {
         if each.Bandwidth == rep.Bandwidth {
            fmt.Print("!")
         }
         fmt.Println(each)
      }
      return nil
   }
   file, err := os.Create(str.base + rep.Ext())
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
      private_key, err := os.ReadFile(f.Private_Key)
      if err != nil {
         return err
      }
      client_ID, err := os.ReadFile(f.Client_ID)
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
      keys, err := mod.Post(str.post)
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

type stream struct {
   bandwidth int
   base string
   dash.Representations
   post widevine.Poster
}
