package mech

import (
   "fmt"
   "github.com/89z/std/hls"
   "github.com/89z/std/os"
   "io"
   "net/url"
)

type stream_HLS struct {
   base *url.URL
   basename string
   flag Flags
   hls.Streams
}

func (f Flags) HLS(base string) error {
   res, err := client.Get(f.Address)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   master, err := hls.New_Scanner(res.Body).Master()
   if err != nil {
      return err
   }
   var str stream_HLS
   str.Streams = master.Streams
   str.base = res.Request.URL
   str.basename = base
   str.flag = f
   return str.download()
}

func new_segment(addr string) (*hls.Segment, error) {
   res, err := client.Get(addr)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   return hls.New_Scanner(res.Body).Segment()
}

func (s stream_HLS) download() error {
   if s.flag.Bandwidth_Video <= 0 {
      return nil
   }
   stream := s.Get_Bandwidth(s.flag.Bandwidth_Video)
   if s.flag.Info {
      for _, each := range s.Streams {
         if each.Bandwidth == stream.Bandwidth {
            fmt.Print("!")
         }
         fmt.Println(each)
      }
      return nil
   }
   seg, err := new_segment(stream.URI)
   if err != nil {
      return err
   }
   res, err := client.Get(seg.Key)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   file, err := os.Create(s.basename + ".ts")
   if err != nil {
      return err
   }
   defer file.Close()
   pro := os.Progress_Chunks(file, len(seg.Protected))
   key, err := io.ReadAll(res.Body)
   if err != nil {
      return err
   }
   block, err := hls.New_Block(key)
   if err != nil {
      return err
   }
   for _, addr := range seg.Protected {
      res, err := client.Level(0).Get(addr)
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
