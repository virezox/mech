package main

import (
   "fmt"
   "github.com/89z/format"
   "github.com/89z/format/hls"
   "github.com/89z/mech/paramount"
   "io"
   "net/http"
   "os"
   "sort"
)

func download(addr, base string) error {
   seg, err := new_segment(addr)
   if err != nil {
      return err
   }
   fmt.Println("GET", seg.Raw_Key)
   res, err := http.Get(seg.Raw_Key)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   file, err := os.Create(base + hls.TS)
   if err != nil {
      return err
   }
   defer file.Close()
   pro := format.Progress_Chunks(file, len(seg.Protected))
   key, err := io.ReadAll(res.Body)
   if err != nil {
      return err
   }
   block, err := hls.New_Block(key)
   if err != nil {
      return err
   }
   for _, addr := range seg.Protected {
      res, err := http.Get(addr)
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
   addr, err := paramount.New_Media(d.GUID).HLS()
   if err != nil {
      return err
   }
   fmt.Println("GET", addr)
   res, err := http.Get(addr.String())
   if err != nil {
      return err
   }
   defer res.Body.Close()
   master, err := hls.New_Scanner(res.Body).Master()
   if err != nil {
      return err
   }
   sort.Slice(master.Streams, func(a, b int) bool {
      return master.Streams[a].Bandwidth < master.Streams[b].Bandwidth
   })
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
      return download(stream.Raw_URI, d.Base())
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
