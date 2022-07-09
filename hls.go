package mech

import (
   "fmt"
   "github.com/89z/rosso/hls"
   "github.com/89z/rosso/os"
   "io"
   "net/url"
)

type Flag_HLS struct {
   Address string
   Name string
   Info bool
   base *url.URL
   master *hls.Master
}

func (f *Flag_HLS) Master() error {
   res, err := client.Get(f.Address)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   f.base = res.Request.URL
   f.master, err = hls.New_Scanner(res.Body).Master()
   if err != nil {
      return err
   }
   return nil
}

func (f Flag_HLS) Stream(a hls.Filter[hls.Stream], b hls.Index[hls.Stream]) error {
   return download(f, a, b, f.master.Stream)
}

func (f Flag_HLS) Media(a hls.Filter[hls.Media], b hls.Index[hls.Media]) error {
   return download(f, a, b, f.master.Media)
}

func download[T hls.Mixed](f Flag_HLS, a hls.Filter[T], b hls.Index[T], s hls.Slice[T]) error {
   items := s.Filter(a)
   target := items.Index(b)
   if f.Info {
      for i, item := range items {
         if i == target {
            fmt.Print("!")
         }
         fmt.Println(item)
      }
      return nil
   }
   item := items[target]
   seg_addr, err := f.base.Parse(item.URI())
   if err != nil {
      return err
   }
   res, err := client.Get(seg_addr.String())
   if err != nil {
      return err
   }
   defer res.Body.Close()
   seg, err := hls.New_Scanner(res.Body).Segment()
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
   file, err := os.Create(f.Name + item.Ext())
   if err != nil {
      return err
   }
   defer file.Close()
   pro := os.Progress_Chunks(file, len(seg.URI))
   for _, raw := range seg.URI {
      addr, err := seg_addr.Parse(raw)
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
