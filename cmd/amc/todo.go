package main

import (
   "encoding/xml"
   "fmt"
   "github.com/89z/mech/amc"
   "github.com/89z/mech/widevine"
   "github.com/89z/std/dash"
   "github.com/89z/std/mp4"
   "github.com/89z/std/os"
   "io"
)

func download(key []byte base string) error {
   file, err := os.Create(str.base + rep.Ext())
   if err != nil {
      return err
   }
   defer file.Close()
   res, err := amc.Client.Redirect(nil).Get(rep.Initialization())
   if err != nil {
      return err
   }
   defer res.Body.Close()
   media := rep.Media()
   pro := os.Progress_Chunks(file, len(media))
   dec := mp4.New_Decrypt(pro)
   if err := dec.Init(res.Body); err != nil {
      return err
   }
   for _, addr := range media {
      res, err := amc.Client.Redirect(nil).Level(0).Get(addr.String())
      if err != nil {
         return err
      }
      pro.Add_Chunk(res.ContentLength)
      if f.key != nil {
         err = dec.Segment(res.Body, f.key)
      } else {
         _, err = io.Copy(pro, res.Body)
      }
      if err != nil {
         return err
      }
      if err := res.Body.Close(); err != nil {
         return err
      }
   }
   return nil
}

func (f *flags) key(p widevine.Poster, raw_key_id string) ([]byte, error) {
   private_key, err := os.ReadFile(f.private_key)
   if err != nil {
      return err
   }
   client_ID, err := os.ReadFile(f.client_ID)
   if err != nil {
      return err
   }
   key_ID, err := widevine.Key_ID(raw_key_id)
   if err != nil {
      return err
   }
   mod, err := widevine.New_Module(private_key, client_ID, key_ID)
   if err != nil {
      return err
   }
   keys, err := mod.Post(p)
   if err != nil {
      return err
   }
   return keys.Content().Key, nil
}
