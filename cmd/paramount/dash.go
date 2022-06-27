package main

import (
   "encoding/xml"
   "fmt"
   "github.com/89z/format"
   "github.com/89z/format/dash"
   "github.com/89z/mech"
   "github.com/89z/mech/paramount"
   "github.com/89z/mech/widevine"
   "os"
   "sort"
)

func (d *downloader) set_key() error {
   private_key, err := os.ReadFile(d.pem)
   if err != nil {
      return err
   }
   client_id, err := os.ReadFile(d.client)
   if err != nil {
      return err
   }
   key_id, err := widevine.Key_ID(d.media.Protection().Default_KID)
   if err != nil {
      return err
   }
   mod, err := widevine.New_Module(private_key, client_id, key_id)
   if err != nil {
      return err
   }
   session, err := paramount.New_Session(d.Preview.GUID)
   if err != nil {
      return err
   }
   contents, err := mod.Request(session)
   if err != nil {
      return err
   }
   d.key = contents.Content().Key
   return nil
}

func (d downloader) DASH(video, audio int64) error {
   addr, err := paramount.New_Media(d.GUID).DASH()
   if err != nil {
      return err
   }
   res, err := paramount.Client.Get(addr.String())
   if err != nil {
      return err
   }
   defer res.Body.Close()
   d.url = addr
   if err := xml.NewDecoder(res.Body).Decode(&d.media); err != nil {
      return err
   }
   if err := d.download(audio, dash.Audio); err != nil {
      return err
   }
   return d.download(video, dash.Video)
}

func (d *downloader) download(band int64, fn dash.Represent_Func) error {
   if band == 0 {
      return nil
   }
   reps := d.media.Represents(fn)
   sort.Slice(reps, func(a, b int) bool {
      return reps[a].Bandwidth < reps[b].Bandwidth
   })
   rep := reps.Represent(band)
   if d.info {
      for _, each := range reps {
         if each.Bandwidth == rep.Bandwidth {
            fmt.Print("!")
         }
         fmt.Println(each)
      }
   } else {
      if d.key == nil {
         err := d.set_key()
         if err != nil {
            return err
         }
      }
      ext, err := mech.Extension_By_Type(rep.MIME_Type)
      if err != nil {
         return err
      }
      file, err := format.Create(d.Base()+ext)
      if err != nil {
         return err
      }
      defer file.Close()
      initial, err := rep.Initial(d.url)
      if err != nil {
         return err
      }
      res, err := paramount.Client.Get(initial.String())
      if err != nil {
         return err
      }
      defer res.Body.Close()
      if err := dash.Decrypt_Init(file, res.Body); err != nil {
         return err
      }
      media, err := rep.Media(d.url)
      if err != nil {
         return err
      }
      pro := format.Progress_Chunks(file, len(media))
      for _, addr := range media {
         res, err := paramount.Client.Level(0).Get(addr.String())
         if err != nil {
            return err
         }
         pro.Add_Chunk(res.ContentLength)
         if err := dash.Decrypt(pro, res.Body, d.key); err != nil {
            return err
         }
         if err := res.Body.Close(); err != nil {
            return err
         }
      }
   }
   return nil
}
