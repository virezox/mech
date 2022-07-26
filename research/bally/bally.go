package main

import (
   "encoding/hex"
   "encoding/xml"
   "fmt"
   "github.com/89z/mech/research/bally/dash"
   "github.com/89z/rosso/http"
   "github.com/89z/rosso/mp4"
   "github.com/89z/rosso/os"
   "strings"
   "time"
)

func (f flags) DASH() error {
   if !f.Info {
      audio, err := os.Create(f.Name + ".m4a")
      if err != nil {
         return err
      }
      defer audio.Close()
      f.audio = mp4.New_Decrypt(audio)
      video, err := os.Create(f.Name + ".m4v")
      if err != nil {
         return err
      }
      defer video.Close()
      f.video = mp4.New_Decrypt(video)
   }
   var i int
   for {
      res, err := client.Redirect(nil).Get(f.address)
      if err != nil {
         return err
      }
      defer res.Body.Close()
      var pre dash.Presentation
      if err := xml.NewDecoder(res.Body).Decode(&pre); err != nil {
         return err
      }
      reps := pre.Representation()
      if len(reps) >= 1 {
         f.base = reps[0].BaseURL
         fmt.Println("audio", i)
         audio := reps.Audio()
         index := audio.Index(func(a, b dash.Representation) bool {
            return strings.Contains(b.Codecs, f.codec)
         })
         if err := f.DASH_Get(audio, index, f.audio, &f.apos); err != nil {
            return err
         }
         fmt.Println("video", i)
         video := reps.Video()
         if err := f.DASH_Get(video, video.Bandwidth(f.bandwidth), f.video, &f.vpos); err != nil {
            return err
         }
      }
      if f.Info {
         return nil
      }
      f.init = true
      fmt.Println("Sleep", i)
      time.Sleep(9 * time.Second)
      i++
   }
}

func (f flags) DASH_Get(items dash.Representations, index int, dec mp4.Decrypt, pos *int) error {
   if f.Info {
      for i, item := range items {
         if i == index {
            fmt.Print("!")
         }
         fmt.Println(item)
      }
      return nil
   }
   item := items[index]
   if !f.init {
      req, err := http.NewRequest("GET", f.base + item.Initialization(), nil)
      if err != nil {
         return err
      }
      res, err := client.Redirect(nil).Do(req)
      if err != nil {
         return err
      }
      defer res.Body.Close()
      if err := dec.Init(res.Body); err != nil {
         return err
      }
   }
   key, err := hex.DecodeString(f.key)
   if err != nil {
      return err
   }
   media := item.Media(pos)
   for _, ref := range media {
      req, err := http.NewRequest("GET", f.base + ref, nil)
      if err != nil {
         return err
      }
      res, err := client.Redirect(nil).Do(req)
      if err != nil {
         return err
      }
      if err := dec.Segment(res.Body, key); err != nil {
         return err
      }
      if err := res.Body.Close(); err != nil {
         return err
      }
   }
   return nil
}
