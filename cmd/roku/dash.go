package main

import (
   "fmt"
   "github.com/89z/format"
   "github.com/89z/format/dash"
   "github.com/89z/mech"
   "github.com/89z/mech/roku"
   "net/http"
   "net/url"
   "os"
)

type downloader struct {
   *roku.Content
   *url.URL
   dash.AdaptationSet
   info bool
   key []byte
}

func (d downloader) DASH(bAudio, bVideo int64) error {
   if d.info {
      fmt.Println(d.Content)
   } else {
      site, err := roku.NewCrossSite()
      if err != nil {
         return err
      }
      play, err := site.Playback(d.Meta.ID)
      if err != nil {
         return err
      }
      vine, err := play.Widevine()
      if err != nil {
         return err
      }
      d.key, err = vine.Key()
      if err != nil {
         return err
      }
   }
   video := d.Content.DASH()
   fmt.Println("GET", video.URL)
   res, err := http.Get(video.URL)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   d.URL = res.Request.URL
   d.AdaptationSet, err = dash.NewAdaptationSet(res.Body)
   if err != nil {
      return err
   }
   if err := d.download(dash.Audio, bAudio); err != nil {
      return err
   }
   return d.download(dash.Video, bVideo)
}

func (d downloader) download(typ string, bandwidth int64) error {
   if bandwidth == 0 {
      return nil
   }
   rep := d.MimeType(typ).Represent(bandwidth)
   if d.info {
      for _, ada := range d.MimeType(typ) {
         for _, each := range ada.Representation {
            if each.Bandwidth == rep.Bandwidth {
               fmt.Print("!")
            }
            fmt.Println(each)
         }
      }
   } else {
      ext, err := mech.ExtensionByType(typ)
      if err != nil {
         return err
      }
      file, err := os.Create(d.Base()+ext)
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
