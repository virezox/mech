package main

import (
   "encoding/xml"
   "fmt"
   "github.com/89z/format"
   "github.com/89z/format/dash"
   "github.com/89z/mech"
   "github.com/89z/mech/roku"
   "github.com/89z/mech/widevine"
   "os"
)

func (d *downloader) set_key() error {
   site, err := roku.New_Cross_Site()
   if err != nil {
      return err
   }
   play, err := site.Playback(d.Meta.ID)
   if err != nil {
      return err
   }
   var client widevine.Client
   client.ID, err = os.ReadFile(d.client)
   if err != nil {
      return err
   }
   client.Private_Key, err = os.ReadFile(d.pem)
   if err != nil {
      return err
   }
   client.Raw_Key_ID = d.media.Protection().Default_KID
   content, err := play.Content(client)
   if err != nil {
      return err
   }
   d.key = content.Key
   return nil
}

func (d downloader) DASH(video, audio int64) error {
   if d.info {
      fmt.Println(d.Content)
   }
   video_dash := d.Content.DASH()
   res, err := roku.Client.Get(video_dash.URL)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   d.url = res.Request.URL
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
      file, err := format.Create(d.Content.Base()+ext)
      if err != nil {
         return err
      }
      defer file.Close()
      initial, err := rep.Initial(d.url)
      if err != nil {
         return err
      }
      res, err := roku.Client.Get(initial.String())
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
         res, err := roku.Client.Level(0).Get(addr.String())
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
