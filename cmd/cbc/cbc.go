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

func get_key(addr string) ([]byte, error) {
   fmt.Println("GET", addr)
   res, err := http.Get(addr)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   return io.ReadAll(res.Body)
}

func download(addr *url.URL, name string) error {
   seg, err := new_segment(addr.String())
   if err != nil {
      return err
   }
   key, err := get_key(seg.Raw_Key)
   if err != nil {
      return err
   }
   file, err := os.Create(name)
   if err != nil {
      return err
   }
   defer file.Close()
   pro := format.Progress_Chunks(file, len(seg.Protected))
   block, err := hls.New_Block(key)
   if err != nil {
      return err
   }
   for _, raw_addr := range seg.Protected {
      addr, err := addr.Parse(raw_addr)
      if err != nil {
         return err
      }
      res, err := http.Get(addr.String())
      if err != nil {
         return err
      }
      pro.Add_Chunk(res.ContentLength)
      if _, err := io.Copy(pro, block.Mode_Key(res.Body)); err != nil {
         return err
      }
      if err := res.Body.Close(); err != nil {
         return err
      }
   }
   return nil
}

func new_segment(addr string) (*hls.Segment, error) {
   fmt.Println("GET", addr)
   res, err := http.Get(addr)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   return hls.New_Scanner(res.Body).Segment()
}

func do_profile(email, password string) error {
   home, err := os.UserHomeDir()
   if err != nil {
      return err
   }
   login, err := cbc.New_Login(email, password)
   if err != nil {
      return err
   }
   web, err := login.Web_Token()
   if err != nil {
      return err
   }
   top, err := web.Over_The_Top()
   if err != nil {
      return err
   }
   profile, err := top.Profile()
   if err != nil {
      return err
   }
   return profile.Create(home + "/mech/cbc.json")
}

func new_master(id, address, audio string, video int64, info bool) error {
   home, err := os.UserHomeDir()
   if err != nil {
      return err
   }
   profile, err := cbc.Open_Profile(home + "/mech/cbc.json")
   if err != nil {
      return err
   }
   if id == "" {
      id = cbc.Get_ID(address)
   }
   asset, err := cbc.New_Asset(id)
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
   master, err := hls.New_Scanner(res.Body).Master()
   if err != nil {
      return err
   }
   if info {
      fmt.Println(asset)
      video := master.Streams.Get_Bandwidth(video)
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
         medium := master.Media.Get_Name(audio)
         addr, err := res.Request.URL.Parse(medium.Raw_URI)
         if err != nil {
            return err
         }
         if err := download(addr, asset.AppleContentID + hls.AAC); err != nil {
            return err
         }
      }
      if video >= 1 {
         medium := master.Streams.Get_Bandwidth(video)
         addr, err := res.Request.URL.Parse(medium.Raw_URI)
         if err != nil {
            return err
         }
         return download(addr, asset.AppleContentID + hls.TS)
      }
   }
   return nil
}
