package main

import (
   "fmt"
   "github.com/89z/std/hls"
   "github.com/89z/std/os"
   "github.com/89z/mech/cbc"
   "io"
   "net/url"
)

func (f flags) master() error {
   home, err := os.UserHomeDir()
   if err != nil {
      return err
   }
   profile, err := cbc.Open_Profile(home + "/mech/cbc.json")
   if err != nil {
      return err
   }
   asset, err := cbc.New_Asset(f.id)
   if err != nil {
      return err
   }
   asset_media, err := profile.Media(asset)
   if err != nil {
      return err
   }
   res, err := cbc.Client.Get(*asset_media.URL)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   master, err := hls.New_Scanner(res.Body).Master()
   if err != nil {
      return err
   }
   streams := master.Streams.Video()
   stream := streams.Bandwidth(f.video_bandwidth)
   media := master.Media.Audio()
   medium := media.Get_Name(f.audio_name)
   if f.info {
      for _, elem := range streams {
         if elem.Bandwidth == stream.Bandwidth {
            fmt.Print("!")
         }
         fmt.Println(elem)
      }
      for _, elem := range media {
         if elem.Name == medium.Name {
            fmt.Print("!")
         }
         fmt.Println(elem)
      }
   } else {
      if f.audio_name != "" {
         addr, err := res.Request.URL.Parse(medium.URI)
         if err != nil {
            return err
         }
         if err := download(addr, asset.AppleContentId + medium.Ext()); err != nil {
            return err
         }
      }
      if f.video_bandwidth >= 1 {
         addr, err := res.Request.URL.Parse(stream.URI)
         if err != nil {
            return err
         }
         return download(addr, asset.AppleContentId + stream.Ext())
      }
   }
   return nil
}

func get_key(addr string) ([]byte, error) {
   res, err := cbc.Client.Get(addr)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   return io.ReadAll(res.Body)
}

func (f flags) profile() error {
   home, err := os.UserHomeDir()
   if err != nil {
      return err
   }
   login, err := cbc.New_Login(f.email, f.password)
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

func download(addr *url.URL, name string) error {
   seg, err := new_segment(addr.String())
   if err != nil {
      return err
   }
   key, err := get_key(seg.Key)
   if err != nil {
      return err
   }
   file, err := os.Create(name)
   if err != nil {
      return err
   }
   defer file.Close()
   pro := os.Progress_Chunks(file, len(seg.URI))
   block, err := hls.New_Block(key)
   if err != nil {
      return err
   }
   for _, raw_addr := range seg.URI {
      addr, err := addr.Parse(raw_addr)
      if err != nil {
         return err
      }
      res, err := cbc.Client.Level(0).Get(addr.String())
      if err != nil {
         return err
      }
      pro.Add_Chunk(res.ContentLength)
      text, err := io.ReadAll(res.Body)
      if err != nil {
         return err
      }
      text = block.Decrypt_Key(text)
      if _, err := pro.Write(text); err != nil {
         return err
      }
      if err := res.Body.Close(); err != nil {
         return err
      }
   }
   return nil
}

func new_segment(addr string) (*hls.Segment, error) {
   res, err := cbc.Client.Get(addr)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   return hls.New_Scanner(res.Body).Segment()
}

