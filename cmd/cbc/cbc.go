package main

import (
   "fmt"
   "github.com/89z/format"
   "github.com/89z/format/hls"
   "github.com/89z/mech/cbc"
   "io"
   "net/http"
   "net/url"
   "os"
)

func getKey(addr string) ([]byte, error) {
   fmt.Println("GET", addr)
   res, err := http.Get(addr)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   return io.ReadAll(res.Body)
}

func download(addr *url.URL, name string) error {
   seg, err := newSegment(addr.String())
   if err != nil {
      return err
   }
   key, err := getKey(seg.RawKey)
   if err != nil {
      return err
   }
   file, err := os.Create(name)
   if err != nil {
      return err
   }
   defer file.Close()
   pro := format.ProgressChunks(file, len(seg.Protected))
   block, err := hls.NewBlock(key)
   if err != nil {
      return err
   }
   for _, rawURL := range seg.Protected {
      addr, err := addr.Parse(rawURL)
      if err != nil {
         return err
      }
      res, err := http.Get(addr.String())
      if err != nil {
         return err
      }
      pro.AddChunk(res.ContentLength)
      if _, err := io.Copy(pro, block.ModeKey(res.Body)); err != nil {
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
   return hls.NewScanner(res.Body).Segment()
}

func doProfile(email, password string) error {
   home, err := os.UserHomeDir()
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
   return profile.Create(home + "/mech/cbc.json")
}
func newMaster(id, address, audio string, video int64, info bool) error {
   home, err := os.UserHomeDir()
   if err != nil {
      return err
   }
   profile, err := cbc.OpenProfile(home + "/mech/cbc.json")
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
   master, err := hls.NewScanner(res.Body).Master()
   if err != nil {
      return err
   }
   if info {
      fmt.Println(asset)
      video := master.Streams.GetBandwidth(video)
      for _, stream := range master.Streams {
         if stream.Bandwidth == video.Bandwidth {
            fmt.Print("!")
         }
         fmt.Println(stream)
      }
      for _, medium := range master.Media {
         fmt.Println(medium)
      }
   } else {
      if audio != "" {
         medium := master.Media.GetName(audio)
         addr, err := res.Request.URL.Parse(medium.RawURI)
         if err != nil {
            return err
         }
         if err := download(addr, asset.AppleContentID + hls.AAC); err != nil {
            return err
         }
      }
      if video >= 1 {
         medium := master.Streams.GetBandwidth(video)
         addr, err := res.Request.URL.Parse(medium.RawURI)
         if err != nil {
            return err
         }
         return download(addr, asset.AppleContentID + hls.TS)
      }
   }
   return nil
}
