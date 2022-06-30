package main

import (
   "encoding/xml"
   "fmt"
   "github.com/89z/format/dash"
   "github.com/89z/format/mp4"
   "github.com/89z/mech/paramount"
   "github.com/89z/mech/widevine"
   "github.com/89z/format"
   "os"
)

func (d *downloader) set_key() error {
   private_key, err := os.ReadFile(d.pem)
   if err != nil {
      return err
   }
   client_ID, err := os.ReadFile(d.client)
   if err != nil {
      return err
   }
   raw_key_id := d.media.Representations()[0].ContentProtection.Default_KID
   key_ID, err := widevine.Key_ID(raw_key_id)
   if err != nil {
      return err
   }
   mod, err := widevine.New_Module(private_key, client_ID, key_ID)
   if err != nil {
      return err
   }
   session, err := paramount.New_Session(d.Preview.GUID)
   if err != nil {
      return err
   }
   keys, err := mod.Post(session)
   if err != nil {
      return err
   }
   d.key = keys.Content().Key
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
   reps := d.media.Representations().Codecs("mp4a")
   if err := d.download(audio, reps); err != nil {
      return err
   }
   reps = d.media.Representations().Codecs("avc1")
   return d.download(video, reps)
}

func (d *downloader) download(bandwidth int64, r dash.Representations) error {
   if bandwidth == 0 {
      return nil
   }
   rep := r.Get_Bandwidth(bandwidth)
   if d.info {
      for _, each := range r {
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
      file, err := format.Create(d.Base()+rep.Ext())
      if err != nil {
         return err
      }
      defer file.Close()
      initial, err := d.url.Parse(rep.Initialization())
      if err != nil {
         return err
      }
      res, err := paramount.Client.Get(initial.String())
      if err != nil {
         return err
      }
      defer res.Body.Close()
      media := rep.Media()
      pro := format.Progress_Chunks(file, len(media))
      dec := mp4.New_Decrypt(pro)
      if err := dec.Init(res.Body); err != nil {
         return err
      }
      for _, medium := range media {
         addr, err := d.url.Parse(medium)
         if err != nil {
            return err
         }
         res, err := paramount.Client.Level(0).Get(addr.String())
         if err != nil {
            return err
         }
         pro.Add_Chunk(res.ContentLength)
         if err := dec.Segment(res.Body, d.key); err != nil {
            return err
         }
         if err := res.Body.Close(); err != nil {
            return err
         }
      }
   }
   return nil
}
