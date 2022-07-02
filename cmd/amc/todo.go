package main

import (
   "encoding/xml"
   "fmt"
   "github.com/89z/mech/amc"
   "github.com/89z/mech/widevine"
   "github.com/89z/std/dash"
   "github.com/89z/std/os"
   "github.com/89z/std/mp4"
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
   var str stream
   str.playback, err = auth.Playback(f.nid)
   if err != nil {
      return err
   }
   source := play.Source()
   res, err := amc.Client.Redirect(nil).Get(source.Src)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   var media dash.Media
   if err := xml.NewDecoder(res.Body).Decode(&media); err != nil {
      return err
   }
   reps := media.Representations()
   str.Representations = reps.Audio()
   str.bandwidth = f.bandwidth_audio
   if err := information(str); err != nil {
      return err
   }
   str.Representations = reps.Video()
   str.bandwidth = f.bandwidth_video
   return information(str)
}

type stream struct {
   bandwidth int
   dash.Representations
   playback *amc.Playback
}

func (f flags) information(str stream) error {
   if str.bandwidth >= 1 {
      rep := str.Get_Bandwidth(str.bandwidth)
      if f.info {
         for _, each := range str.Representations {
            if each.Bandwidth == rep.Bandwidth {
               fmt.Print("!")
            }
            fmt.Println(each)
         }
      } else {
         return str.download(rep)
      }
   }
   return nil
}

func (s stream) download(rep *dash.Representation) error {
   var key []byte
   if source.Key_Systems != nil {
      private_key, err := os.ReadFile(f.private_key)
      if err != nil {
         return err
      }
      client_ID, err := os.ReadFile(f.client_ID)
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
      keys, err := mod.Post(play)
      if err != nil {
         return err
      }
      key = keys.Content().Key
   }
   file, err := os.Create(base + rep.Ext())
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
   for _, addr := range media {
      res, err := amc.Client.Redirect(nil).Level(0).Get(addr)
      if err != nil {
         return err
      }
      pro.Add_Chunk(res.ContentLength)
      if key != nil {
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
