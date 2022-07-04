package mech

import (
   "fmt"
   "github.com/89z/std/hls"
   "github.com/89z/std/os"
   "io"
   "net/url"
)

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
   str.base = res.Request.URL
   str.basename = base
   str.flag = f
   str.Streams = master.Streams
   return str.download()
}

type stream_HLS struct {
   base *url.URL
   basename string
   flag Flags
   hls.Streams
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
   if s.flag.Video_Bandwidth <= 0 {
      return nil
   }
   s.Streams = s.Streams.Video()
   stream := s.Bandwidth(s.flag.Video_Bandwidth)
   if s.flag.Info {
      for _, elem := range s.Streams {
         if elem.Bandwidth == stream.Bandwidth {
            fmt.Print("!")
         }
         fmt.Println(elem)
      }
      return nil
   }
   base, err := s.base.Parse(stream.URI)
   if err != nil {
      return err
   }
   seg, err := new_segment(base.String())
   if err != nil {
      return err
   }
   var block *hls.Block
   if seg.Key != "" {
      res, err := client.Get(seg.Key)
      if err != nil {
         return err
      }
      defer res.Body.Close()
      key, err := io.ReadAll(res.Body)
      if err != nil {
         return err
      }
      block, err = hls.New_Block(key)
      if err != nil {
         return err
      }
   }
   file, err := os.Create(s.basename + ".ts")
   if err != nil {
      return err
   }
   defer file.Close()
   pro := os.Progress_Chunks(file, len(seg.URI))
   for _, raw := range seg.URI {
      addr, err := base.Parse(raw)
      if err != nil {
         return err
      }
      res, err := client.Level(0).Get(addr.String())
      if err != nil {
         return err
      }
      pro.Add_Chunk(res.ContentLength)
      if block != nil {
         text, err := io.ReadAll(res.Body)
         if err != nil {
            return err
         }
         text = block.Decrypt_Key(text)
         if _, err := pro.Write(text); err != nil {
            return err
         }
      } else {
         _, err := io.Copy(pro, res.Body)
         if err != nil {
            return err
         }
      }
      if err := res.Body.Close(); err != nil {
         return err
      }
   }
   return nil
}

