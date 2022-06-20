package main

import (
   "encoding/xml"
   "errors"
   "fmt"
   "github.com/89z/format"
   "github.com/89z/format/dash"
   "github.com/89z/mech"
   "github.com/89z/mech/amc"
   "io"
   "net/http"
   "os"
)

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
   fmt.Println("GET", source.Src)
   res, err := http.Get(source.Src)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      return errors.New(res.Status)
   }
   d.url = res.Request.URL
   if err := xml.NewDecoder(res.Body).Decode(&d.media); err != nil {
      return err
   }
   if err := d.download(audio, dash.Audio); err != nil {
      return err
   }
   return d.download(video, dash.Video)
}

func (d *downloader) set_key() error {
   var (
      client amc.Client
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
   client.Raw_Key_ID = d.media.Protection().Default_KID
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
      file, err := os.Create(d.Base()+ext)
      if err != nil {
         return err
      }
      defer file.Close()
      init, err := rep.Initial(d.url)
      if err != nil {
         return err
      }
      fmt.Println("GET", init)
      res, err := http.Get(init.String())
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
         res, err := http.Get(addr.String())
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
