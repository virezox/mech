package main

import (
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

func (d *downloader) download(band int64, fn dash.PeriodFunc) error {
   if band == 0 {
      return nil
   }
   reps := d.period.Represents(fn)
   rep := reps.Represent(band)
   if d.info {
      for _, each := range reps {
         if each.Bandwidth == rep.Bandwidth {
            fmt.Print("!")
         }
         fmt.Println(each)
      }
   } else {
      ext, err := mech.ExtensionByType(rep.MimeType)
      if err != nil {
         return err
      }
      file, err := os.Create(d.Base()+ext)
      if err != nil {
         return err
      }
      defer file.Close()
      init, err := rep.Initialization(d.url)
      if err != nil {
         return err
      }
      fmt.Println("GET", init)
      res, err := http.Get(init.String())
      if err != nil {
         return err
      }
      defer res.Body.Close()
      if err := dash.DecryptInit(file, res.Body); err != nil {
         return err
      }
      media, err := rep.Media(d.url)
      if err != nil {
         return err
      }
      if d.key == nil {
         err := d.setKey()
         if err != nil {
            return err
         }
      }
      pro := format.ProgressChunks(file, len(media))
      for _, addr := range media {
         res, err := http.Get(addr.String())
         if err != nil {
            return err
         }
         pro.AddChunk(res.ContentLength)
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

func doLogin(email, password string) error {
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
   return auth.Create(home, "mech/amc.json")
}
func (d downloader) doDASH(address string, nid, video, audio int64) error {
   home, err := os.UserHomeDir()
   if err != nil {
      return err
   }
   auth, err := amc.OpenAuth(home, "mech/amc.json")
   if err != nil {
      return err
   }
   if err := auth.Refresh(); err != nil {
      return err
   }
   if err := auth.Create(home, "mech/amc.json"); err != nil {
      return err
   }
   if nid == 0 {
      nid, err = amc.GetNID(address)
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
   d.period, err = dash.NewPeriod(res.Body)
   if err != nil {
      return err
   }
   if err := d.download(audio, dash.Audio); err != nil {
      return err
   }
   return d.download(video, dash.Video)
}

func (d *downloader) setKey() error {
   var (
      client amc.Client
      err error
   )
   client.ID, err = os.ReadFile(d.client)
   if err != nil {
      return err
   }
   client.PrivateKey, err = os.ReadFile(d.pem)
   if err != nil {
      return err
   }
   client.RawKeyID = d.period.Protection().Default_KID
   content, err := d.Playback.Content(client)
   if err != nil {
      return err
   }
   d.key = content.Key
   return nil
}
