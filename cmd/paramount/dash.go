package main

import (
   "fmt"
   "github.com/89z/format"
   "github.com/89z/format/dash"
   "github.com/89z/mech"
   "github.com/89z/mech/paramount"
   "github.com/89z/mech/widevine"
   "net/http"
   "os"
)

func (d downloader) DASH(bandwidth int64) error {
   addr, err := paramount.NewMedia(d.GUID).DASH()
   if err != nil {
      return err
   }
   fmt.Println("GET", addr)
   res, err := http.Get(addr.String())
   if err != nil {
      return err
   }
   defer res.Body.Close()
   d.URL = addr
   d.AdaptationSet, err = dash.NewAdaptationSet(res.Body)
   if err != nil {
      return err
   }
   return d.download(dash.Video, bandwidth)
}

func (d *downloader) download(typ string, bandwidth int64) error {
   if bandwidth == 0 {
      return nil
   }
   // FAIL
   adas := d.MimeType(typ)
   rep := adas.Represent(bandwidth)
   if d.info {
      for _, ada := range adas {
         for _, each := range ada.Representation {
            if each.Bandwidth == rep.Bandwidth {
               fmt.Print("!")
            }
            fmt.Println(each)
         }
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
   sess, err := paramount.NewSession(d.GUID)
   if err != nil {
      return err
   }
   keys, err := mod.Post(sess.URL, sess.Header())
   if err != nil {
      return err
   }
   d.key = keys.Content().Key
   return nil
}
