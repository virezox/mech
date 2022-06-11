package main

import (
   "fmt"
   "github.com/89z/format"
   "github.com/89z/format/dash"
   "github.com/89z/mech"
   "github.com/89z/mech/roku"
   "net/http"
   "os"
)

func (d *downloader) setKey() error {
   site, err := roku.NewCrossSite()
   if err != nil {
      return err
   }
   play, err := site.Playback(d.Meta.ID)
   if err != nil {
      return err
   }
   var client roku.Client
   client.ID, err = os.ReadFile(d.client)
   if err != nil {
      return err
   }
   client.PrivateKey, err = os.ReadFile(d.pem)
   if err != nil {
      return err
   }
   client.RawKeyID = d.period.Protection().Default_KID
   content, err := play.Content(client)
   if err != nil {
      return err
   }
   d.key = content.Key
   return nil
}

func (d downloader) DASH(video, audio int64) error {
   if d.info {
      fmt.Println(d.Content)
   }
   videoDASH := d.Content.DASH()
   fmt.Println("GET", videoDASH.URL)
   res, err := http.Get(videoDASH.URL)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   d.url = res.Request.URL
   d.period, err = dash.NewPeriod(res.Body)
   if err != nil {
      return err
   }
   if err := d.download(audio, dash.Audio); err != nil {
      return err
   }
   return d.download(video, dash.Video)
}

func (d *downloader) download(band int64, fn dash.PeriodFunc) error {
   if band == 0 {
      return nil
   }
   reps := d.period.Represents(fn)
   rep := reps.Represent(band)
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
      ext, err := mech.ExtensionByType(rep.MimeType)
      if err != nil {
         return err
      }
      file, err := os.Create(d.Content.Base()+ext)
      if err != nil {
         return err
      }
      defer file.Close()
      init, err := rep.Initialization(d.url)
      if err != nil {
         return err
      }
      fmt.Println("GET", init)
      res, err := http.Get(init.String())
      if err != nil {
         return err
      }
      defer res.Body.Close()
      if err := dash.DecryptInit(file, res.Body); err != nil {
         return err
      }
      media, err := rep.Media(d.url)
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
