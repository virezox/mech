package main

import (
   "encoding/xml"
   "fmt"
   "github.com/89z/mech/amc"
   "github.com/89z/mech/widevine"
   "github.com/89z/std/dash"
   "github.com/89z/std/mp4"
   "github.com/89z/std/os"
   "io"
)

func (f flags) DASH() error {
   home, err := os.UserHomeDir()
   if err != nil {
      return err
   }
   auth, err := amc.Open_Auth(home + "/mech/amc.json")
   if err != nil {
      return err
   }
   if err := auth.Refresh(); err != nil {
      return err
   }
   if err := auth.Create(home + "/mech/amc.json"); err != nil {
      return err
   }
   if f.nid == 0 {
      f.nid, err = amc.Get_NID(f.address)
      if err != nil {
         return err
      }
   }
   play, err = auth.Playback(nid)
   if err != nil {
      return err
   }
   key := play.Source().Key_Systems
   if key == nil {
      // no key
      return nil
   }
   res, err := amc.Client.Redirect(nil).Get(play.DASH().Src)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   // Media
   var media dash.Media
   if err := xml.NewDecoder(res.Body).Decode(&media); err != nil {
      return err
   }
   reps := media.Representations()
   // stream
   var str stream
   str.base = play.Base()
   // audio
   str.Representations = reps.Audio()
   str.bandwidth = f.audio_bandwidth
   if err := f.download(str); err != nil {
      return err
   }
   // video
   str.Representations = reps.Video()
   str.bandwidth = f.video_bandwidth
   return f.download(str)
}

func (f *flags) set_key() error {
   private_key, err := os.ReadFile(f.private_key)
   if err != nil {
      return err
   }
   client_ID, err := os.ReadFile(f.client_ID)
   if err != nil {
      return err
   }
   raw_key_id := f.media.Representations()[0].ContentProtection.Default_KID
   key_ID, err := widevine.Key_ID(raw_key_id)
   if err != nil {
      return err
   }
   mod, err := widevine.New_Module(private_key, client_ID, key_ID)
   if err != nil {
      return err
   }
   keys, err := mod.Post(f.Playback)
   if err != nil {
      return err
   }
   f.key = keys.Content().Key
   return nil
}

func (f *flags) download(str stream) error {
   if str.bandwidth == 0 {
      return nil
   }
   rep := reps.Get_Bandwidth(str.bandwidth)
   if f.info {
      for _, each := range reps {
         if each.Bandwidth == rep.Bandwidth {
            fmt.Print("!")
         }
         fmt.Println(each)
      }
   } else {
      file, err := os.Create(str.base + rep.Ext())
      if err != nil {
         return err
      }
      defer file.Close()
      res, err := amc.Client.Redirect(nil).Get(rep.Initialization())
      if err != nil {
         return err
      }
      defer res.Body.Close()
      media := rep.Media()
      pro := os.Progress_Chunks(file, len(media))
      dec := mp4.New_Decrypt(pro)
      if err := dec.Init(res.Body); err != nil {
         return err
      }
      for _, raw := range media {
         addr, err := f.url.Parse(raw)
         if err != nil {
            return err
         }
         res, err := amc.Client.Redirect(nil).Level(0).Get(addr.String())
         if err != nil {
            return err
         }
         pro.Add_Chunk(res.ContentLength)
         if f.key != nil {
            err = dec.Segment(res.Body, f.key)
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
   }
   return nil
}

