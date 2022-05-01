package main

import (
   "fmt"
   "github.com/89z/format"
   "github.com/89z/format/dash"
   "github.com/89z/mech/channel4"
   "io"
   "net/http"
   "net/url"
   "os"
   "sort"
)

func doToken(token string) error {
   file, err := os.Open(token)
   if err != nil {
      return err
   }
   defer file.Close()
   payload, err := channel4.NewPayload(file)
   if err != nil {
      return err
   }
   widevine, err := payload.Widevine()
   if err != nil {
      return err
   }
   key, err := widevine.Decrypt()
   if err != nil {
      return err
   }
   fmt.Println(key)
   return nil
}

func doManifest(id, address string, bandwidth int64, info bool) error {
   if id == "" {
      id = channel4.ProgramID(address)
   }
   stream, err := channel4.NewStream(id)
   if err != nil {
      return err
   }
   widevine := stream.Widevine()
   fmt.Println("GET", widevine)
   res, err := http.Get(widevine)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   period, err := dash.NewPeriod(res.Body)
   if err != nil {
      return err
   }
   video := period.Video()
   if bandwidth >= 1 {
      sort.Sort(dash.Bandwidth{video, bandwidth})
   }
   if info {
      for _, rep := range video.Representation {
         fmt.Println(rep)
      }
   } else {
      replace := func(set *dash.AdaptationSet) error {
         for _, rep := range set.Representation {
            temp := set.SegmentTemplate.Replace(rep)
            return download(temp, res.Request.URL)
         }
         return nil
      }
      err := replace(video)
      if err != nil {
         return err
      }
      audio := period.Audio(video)
      return replace(audio)
   }
   return nil
}

func download(temp dash.Template, base *url.URL) error {
   addrs, err := temp.URLs(base)
   if err != nil {
      return err
   }
   file, err := os.Create(temp.Initialization)
   if err != nil {
      return err
   }
   defer file.Close()
   pro := format.ProgressChunks(file, len(addrs))
   for _, addr := range addrs {
      res, err := http.Get(addr.String())
      if err != nil {
         return err
      }
      pro.AddChunk(res.ContentLength)
      if _, err := io.Copy(pro, res.Body); err != nil {
         return err
      }
      if err := res.Body.Close(); err != nil {
         return err
      }
   }
   return nil
}
