package mech

import (
   "fmt"
   "github.com/89z/rosso/hls"
   "github.com/89z/rosso/os"
   "io"
)

func (self *Stream) HLS(address string) (*hls.Master, error) {
   res, err := client.Redirect(nil).Get(address)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   self.url = res.Request.URL
   return hls.New_Scanner(res.Body).Master()
}

func (self Stream) HLS_Streams(items hls.Streams, index int) error {
   return hls_get(self, items, index)
}

func (self Stream) HLS_Media(items hls.Media, index int) error {
   return hls_get(self, items, index)
}

func hls_get[T hls.Mixed](str Stream, items []T, index int) error {
   if str.Info {
      for i, item := range items {
         if i == index {
            fmt.Print("!")
         }
         fmt.Println(item)
      }
      return nil
   }
   item := items[index]
   file, err := os.Create(str.Base + item.Ext())
   if err != nil {
      return err
   }
   defer file.Close()
   address, err := str.url.Parse(item.URI())
   if err != nil {
      return err
   }
   res, err := client.Get(address.String())
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
   pro := os.Progress_Chunks(file, len(seg.URI))
   for _, raw := range seg.URI {
      address, err := str.url.Parse(raw)
      if err != nil {
         return err
      }
      res, err := client.Level(0).Redirect(nil).Get(address.String())
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
