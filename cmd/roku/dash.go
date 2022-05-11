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

func doDASH(con *roku.Content, bandwidth int64, info bool) error {
   video := con.DASH()
   fmt.Println("GET", video.URL)
   res, err := http.Get(video.URL)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   adas, err := dash.NewAdaptationSet(res.Body)
   if err != nil {
      return err
   }
   if !info {
      video := adas.MimeType(dash.Video).Represent(bandwidth)
      init, err := video.Initialization(res.Request.URL)
      if err != nil {
         return err
      }
      media, err := video.Media(res.Request.URL)
      if err != nil {
         return err
      }
      return download(init, media, con)
   }
   for _, ada := range adas.MimeType(dash.Video) {
      for _, rep := range ada.Representation {
         fmt.Println(rep)
      }
   }
   return nil
}

func download(init *url.URL, media []*url.URL, con *roku.Content) error {
   ext, err := mech.ExtensionByType(dash.Video)
   if err != nil {
      return err
   }
   file, err := os.Create(con.Base() + ext)
   if err != nil {
      return err
   }
   defer file.Close()
   fmt.Println("GET", init)
   res, err := http.Get(init.String())
   if err != nil {
      return err
   }
   defer res.Body.Close()
   if _, err := file.ReadFrom(res.Body); err != nil {
      return err
   }
   key, err := getKey(con.Meta.ID)
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
      if err := dash.Decrypt(pro, res.Body, key); err != nil {
         return err
      }
      if err := res.Body.Close(); err != nil {
         return err
      }
   }
   return nil
}

func getKey(id string) ([]byte, error) {
   site, err := roku.NewCrossSite()
   if err != nil {
      return nil, err
   }
   play, err := site.Playback(id)
   if err != nil {
      return nil, err
   }
   vine, err := play.Widevine()
   if err != nil {
      return nil, err
   }
   return vine.Key()
}
