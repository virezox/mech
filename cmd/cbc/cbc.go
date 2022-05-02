package main

import (
   "fmt"
   "github.com/89z/format"
   "github.com/89z/format/hls"
   "github.com/89z/mech/cbc"
   "net/http"
   "net/url"
   "os"
   "sort"
)

func doManifest(id, address string, bandwidth int, info bool) error {
   cache, err := os.UserCacheDir()
   if err != nil {
      return err
   }
   profile, err := cbc.OpenProfile(cache, "mech/cbc.json")
   if err != nil {
      return err
   }
   if id == "" {
      id = cbc.GetID(address)
   }
   asset, err := cbc.NewAsset(id)
   if err != nil {
      return err
   }
   media, err := profile.Media(asset)
   if err != nil {
      return err
   }
   fmt.Println("GET", media.URL)
   res, err := http.Get(media.URL)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   master, err := hls.NewScanner(res.Body).Master(res.Request.URL)
   if err != nil {
      return err
   }
   if bandwidth >= 1 {
      sort.Sort(hls.Bandwidth{master, bandwidth})
   }
   if info {
      fmt.Println(asset)
   }
   for _, video := range master.Stream {
      if info {
         fmt.Println(video)
      } else {
         err := download(video.URI, asset.AppleContentID + ".video")
         if err != nil {
            return err
         }
         audio := master.Audio(video)
         return download(audio.URI, asset.AppleContentID + ".audio")
      }
   }
   return nil
}

func download(addr *url.URL, name string) error {
   seg, err := newSegment(addr.String())
   if err != nil {
      return err
   }
   fmt.Println("GET", seg.Key)
   res, err := http.Get(seg.Key.String())
   if err != nil {
      return err
   }
   defer res.Body.Close()
   block, err := hls.NewCipher(res.Body)
   if err != nil {
      return err
   }
   file, err := os.Create(name)
   if err != nil {
      return err
   }
   defer file.Close()
   pro := format.ProgressChunks(file, len(seg.Info))
   for _, info := range seg.Info {
      res, err := http.Get(info.URI.String())
      if err != nil {
         return err
      }
      pro.AddChunk(res.ContentLength)
      if _, err := block.Copy(pro, res.Body, info.IV); err != nil {
         return err
      }
      if err := res.Body.Close(); err != nil {
         return err
      }
   }
   return nil
}

func newSegment(addr string) (*hls.Segment, error) {
   fmt.Println("GET", addr)
   res, err := http.Get(addr)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   return hls.NewScanner(res.Body).Segment(res.Request.URL)
}

func doProfile(email, password string) error {
   cache, err := os.UserCacheDir()
   if err != nil {
      return err
   }
   login, err := cbc.NewLogin(email, password)
   if err != nil {
      return err
   }
   web, err := login.WebToken()
   if err != nil {
      return err
   }
   top, err := web.OverTheTop()
   if err != nil {
      return err
   }
   profile, err := top.Profile()
   if err != nil {
      return err
   }
   return profile.Create(cache, "mech/cbc.json")
}
