package main

import (
   "fmt"
   "github.com/89z/format"
   "github.com/89z/format/dash"
   "github.com/89z/mech"
   "github.com/89z/mech/roku"
   "github.com/89z/mech/widevine"
   "net/http"
   "os"
)

func (d downloader) DASH(bVideo, bAudio int64) error {
   if d.info {
      fmt.Println(d.Content)
   }
   video := d.Content.DASH()
   fmt.Println("GET", video.URL)
   res, err := http.Get(video.URL)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   d.URL = res.Request.URL
   d.Period, err = dash.NewPeriod(res.Body)
   if err != nil {
      return err
   }
   if err := d.download(dash.Audio, bAudio); err != nil {
      return err
   }
   return d.download(dash.Video, bVideo)
}

func (d *downloader) download(typ string, bandwidth int64) error {
   if bandwidth == 0 {
      return nil
   }
   reps := d.MimeType(typ)
   rep := reps.Represent(bandwidth)
   if d.info {
      for _, each := range reps {
         if each.Bandwidth == rep.Bandwidth {
            fmt.Print("!")
         }
         fmt.Println(each)
      }
   } else {
      if d.key == nil {
         err := d.setKey()
         if err != nil {
            return err
         }
      }
      ext, err := mech.ExtensionByType(typ)
      if err != nil {
         return err
      }
      file, err := os.Create(d.Content.Base()+ext)
      if err != nil {
         return err
      }
      defer file.Close()
      init, err := rep.Initialization(d.URL)
      if err != nil {
         return err
      }
      fmt.Println("GET", init)
      res, err := http.Get(init.String())
      if err != nil {
         return err
      }
      defer res.Body.Close()
      if _, err := file.ReadFrom(res.Body); err != nil {
         return err
      }
      media, err := rep.Media(d.URL)
      if err != nil {
         return err
      }
      pro := format.ProgressChunks(file, len(media))
      for _, addr := range media {
         res, err := http.Get(addr.String())
         if err != nil {
            return err
         }
         pro.AddChunk(res.ContentLength)
         if err := dash.Decrypt(pro, res.Body, d.key); err != nil {
            return err
         }
         if err := res.Body.Close(); err != nil {
            return err
         }
      }
   }
   return nil
}

func (d *downloader) setKey() error {
   privateKey, err := os.ReadFile(d.pem)
   if err != nil {
      return err
   }
   clientID, err := os.ReadFile(d.client)
   if err != nil {
      return err
   }
   kID, err := d.Protection().KID()
   if err != nil {
      return err
   }
   mod, err := widevine.NewModule(privateKey, clientID, kID)
   if err != nil {
      return err
   }
   site, err := roku.NewCrossSite()
   if err != nil {
      return err
   }
   play, err := site.Playback(d.Meta.ID)
   if err != nil {
      return err
   }
   keys, err := mod.Post(play.DRM.Widevine.LicenseServer, nil)
   if err != nil {
      return err
   }
   d.key = keys.Content().Key
   return nil
}
