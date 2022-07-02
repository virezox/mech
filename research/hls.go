package mech

import (
   "fmt"
   "github.com/89z/mech/paramount"
   "github.com/89z/std/hls"
   "github.com/89z/std/os"
   "io"
)

func download(addr, base string) error {
   seg, err := new_segment(addr)
   if err != nil {
      return err
   }
   res, err := paramount.Client.Get(seg.Key)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   file, err := os.Create(base + ".ts")
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
      res, err := paramount.Client.Level(0).Get(addr)
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

func (d downloader) HLS(bandwidth int64) error {
   addr, err := paramount.New_Media(d.guid).HLS()
   if err != nil {
      return err
   }
   res, err := paramount.Client.Get(addr.String())
   if err != nil {
      return err
   }
   defer res.Body.Close()
   master, err := hls.New_Scanner(res.Body).Master()
   if err != nil {
      return err
   }
   stream := master.Streams.Get_Bandwidth(bandwidth)
   if d.info {
      fmt.Println(d.Title)
      for _, each := range master.Streams {
         if each.Bandwidth == stream.Bandwidth {
            fmt.Print("!")
         }
         fmt.Println(each)
      }
   } else {
      return download(stream.URI, d.Base())
   }
   return nil
}

func new_segment(addr string) (*hls.Segment, error) {
   res, err := paramount.Client.Get(addr)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   return hls.New_Scanner(res.Body).Segment()
}
