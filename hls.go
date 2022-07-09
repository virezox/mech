package mech

import (
   "fmt"
   "github.com/89z/rosso/hls"
   "github.com/89z/rosso/os"
   "io"
)

func (f *Flag) Master() (*hls.Master, error) {
   res, err := client.Get(f.Address)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   f.base = res.Request.URL
   return hls.New_Scanner(res.Body).Master()
}

func (f Flag) HLS(items []hls.Mixed, index int) error {
   if f.Info {
      for i, item := range items {
         if i == index {
            fmt.Print("!")
         }
         fmt.Println(item)
      }
      return nil
   }
   item := items[index]
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
