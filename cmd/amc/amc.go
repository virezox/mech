package main

import (
   "encoding/xml"
   "fmt"
   "github.com/89z/format"
   "github.com/89z/format/dash"
   "github.com/89z/mech"
   "github.com/89z/mech/amc"
   "github.com/89z/mech/widevine"
   "io"
   "os"
)

func (d *downloader) set_key() error {
   var (
      client widevine.Client
      err error
   )
   client.ID, err = os.ReadFile(d.client)
   if err != nil {
      return err
   }
   client.Private_Key, err = os.ReadFile(d.pem)
   if err != nil {
      return err
   }
   client.Raw = d.media.Protection().Default_KID
   content, err := d.Playback.Content(client)
   if err != nil {
      return err
   }
   d.key = content.Key
   return nil
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
      res, err := amc.Client.Redirect(nil).Get(initial.String())
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
      if d.key == nil {
         err := d.set_key()
         if err != nil {
            return err
         }
      }
      pro := format.Progress_Chunks(file, len(media))
      for _, addr := range media {
         res, err := amc.Client.Redirect(nil).Level(0).Get(addr.String())
         if err != nil {
            return err
         }
         pro.Add_Chunk(res.ContentLength)
         if d.key != nil {
            err = dash.Decrypt(pro, res.Body, d.key)
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
   }
   return nil
}

func do_login(email, password string) error {
   auth, err := amc.Unauth()
   if err != nil {
      return err
   }
   if err := auth.Login(email, password); err != nil {
      return err
   }
   home, err := os.UserHomeDir()
   if err != nil {
      return err
   }
   return auth.Create(home + "/mech/amc.json")
}
func (d downloader) do_DASH(address string, nid, video, audio int64) error {
   home, err := os.UserHomeDir()
   if err != nil {
      return err
   }
   auth, err := amc.Open_Auth(home + "/mech/amc.json")
   if err != nil {
      return err
   }
   if err := auth.Refresh(); err != nil {
      return err
   }
   if err := auth.Create(home + "/mech/amc.json"); err != nil {
      return err
   }
   if nid == 0 {
      nid, err = amc.Get_NID(address)
      if err != nil {
         return err
      }
   }
   d.Playback, err = auth.Playback(nid)
   if err != nil {
      return err
   }
   source := d.Playback.DASH()
   res, err := amc.Client.Redirect(nil).Get(source.Src)
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
