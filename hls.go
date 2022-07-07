package mech

import (
   "fmt"
   "github.com/89z/rosso/hls"
   "github.com/89z/rosso/os"
   "io"
   "net/url"
   "strings"
)

func (s stream_HLS) download() error {
   if s.flag.Video_Bandwidth <= 0 {
      return nil
   }
   s.stream = s.stream.Filter(func(s hls.Stream) bool {
      return strings.Contains(s.Codecs, "avc1.")
   })
   distance := hls.Bandwidth(s.flag.Video_Bandwidth)
   stream := s.stream.Reduce(func(carry, item hls.Stream) bool {
      return distance(item) < distance(carry)
   })
   if s.flag.Info {
      for _, item := range s.stream {
         if item.Bandwidth == stream.Bandwidth {
            fmt.Print("!")
         }
         fmt.Println(item)
      }
      return nil
   }
   base, err := s.Parse(stream.URI)
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
   file, err := os.Create(s.base + stream.Ext())
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
   str.URL = res.Request.URL
   str.base = base
   str.flag = f
   str.stream = master.Stream
   return str.download()
}

type stream_HLS struct {
   *url.URL
   base string
   flag Flags
   stream hls.Slice[hls.Stream]
}

func new_segment(addr string) (*hls.Segment, error) {
   res, err := client.Get(addr)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   return hls.New_Scanner(res.Body).Segment()
}
