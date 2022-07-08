package mech

import (
   "fmt"
   "github.com/89z/rosso/hls"
   "github.com/89z/rosso/os"
   "io"
)

func new_segment(addr string) (*hls.Segment, error) {
   res, err := client.Get(addr)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   return hls.New_Scanner(res.Body).Segment()
}

type one[T hls.Item] struct {
   hls.Filter[T]
   hls.Reduce[T]
}

type two struct {
   Address string
   Base string
   Info bool
   media one[hls.Media]
   stream one[hls.Stream]
}

func (t two) do() error {
   res, err := client.Get(f.Address)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   master, err := hls.New_Scanner(res.Body).Master()
   if err != nil {
      return err
   }
   if err := three(master.Media, t.media); err != nil {
      return err
   }
   return three(master.Stream, t.stream)
}

func three[T hls.Item](slice hls.Slice[T], callback one[T]) error {
   items := slice.Filter(callback.Filter)
   target := items.Reduce(callback.Reduce)
   if f.Info {
      for _, item := range items {
         if item == *target {
            fmt.Print("!")
         }
         fmt.Println(item)
      }
      return nil
   }
   base, err := res.Request.URL.Parse(target.URI)
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
   file, err := os.Create(f.Base + target.Ext())
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
