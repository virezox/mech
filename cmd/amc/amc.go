package main

import (
   "encoding/xml"
   "fmt"
   "github.com/89z/std/dash"
   "github.com/89z/std/mp4"
   "github.com/89z/std/os"
   "github.com/89z/mech/amc"
   "github.com/89z/mech/widevine"
   "io"
)

func (d *downloader) set_key() error {
   private_key, err := os.ReadFile(d.pem)
   if err != nil {
      return err
   }
   client_ID, err := os.ReadFile(d.client)
   if err != nil {
      return err
   }
   raw_key_id := d.media.Representations()[0].ContentProtection.Default_KID
   key_ID, err := widevine.Key_ID(raw_key_id)
   if err != nil {
      return err
   }
   mod, err := widevine.New_Module(private_key, client_ID, key_ID)
   if err != nil {
      return err
   }
   keys, err := mod.Post(d.Playback)
   if err != nil {
      return err
   }
   d.key = keys.Content().Key
   fmt.Println(keys.Content())
   return nil
}

func (d downloader) do_DASH(address string, nid int64, video, audio int) error {
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
   if nid == 0 {
      nid, err = amc.Get_NID(address)
      if err != nil {
         return err
      }
   }
   d.Playback, err = auth.Playback(nid)
   if err != nil {
      return err
   }
   source := d.Playback.Source()
   res, err := amc.Client.Redirect(nil).Get(source.Src)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   d.url = res.Request.URL
   if err := xml.NewDecoder(res.Body).Decode(&d.media); err != nil {
      return err
   }
   reps := d.media.Representations().Audio()
   if err := d.download(audio, reps); err != nil {
      return err
   }
   reps = d.media.Representations().Video()
   return d.download(video, reps)
}

func (d *downloader) download(bandwidth int, r dash.Representations) error {
   if bandwidth == 0 {
      return nil
   }
   rep := r.Get_Bandwidth(bandwidth)
   if d.info {
      for _, each := range r {
         if each.Bandwidth == rep.Bandwidth {
            fmt.Print("!")
         }
         fmt.Println(each)
      }
   } else {
      file, err := os.Create(d.Base()+rep.Ext())
      if err != nil {
         return err
      }
      defer file.Close()
      initial, err := d.url.Parse(rep.Initialization())
      if err != nil {
         return err
      }
      res, err := amc.Client.Redirect(nil).Get(initial.String())
      if err != nil {
         return err
      }
      defer res.Body.Close()
      if d.key == nil {
         err := d.set_key()
         if err != nil {
            return err
         }
      }
      media := rep.Media()
      pro := os.Progress_Chunks(file, len(media))
      dec := mp4.New_Decrypt(pro)
      if err := dec.Init(res.Body); err != nil {
         return err
      }
      for _, raw := range media {
         addr, err := d.url.Parse(raw)
         if err != nil {
            return err
         }
         res, err := amc.Client.Redirect(nil).Level(0).Get(addr.String())
         if err != nil {
            return err
         }
         pro.Add_Chunk(res.ContentLength)
         if d.key != nil {
            err = dec.Segment(res.Body, d.key)
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

func do_login(email, password string) error {
   auth, err := amc.Unauth()
   if err != nil {
      return err
   }
   if err := auth.Login(email, password); err != nil {
      return err
   }
   home, err := os.UserHomeDir()
   if err != nil {
      return err
   }
   return auth.Create(home + "/mech/amc.json")
}
