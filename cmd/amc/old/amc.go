package main

import (
   "errors"
   "flag"
   "fmt"
   "github.com/89z/format"
   "github.com/89z/format/dash"
   "github.com/89z/mech"
   "github.com/89z/mech/amc"
   "github.com/89z/mech/widevine"
   "io"
   "net/http"
   "net/url"
   "os"
   "path/filepath"
)

func (d *downloader) download(band int64, fn dash.PeriodFunc) error {
   if band == 0 {
      return nil
   }
   reps := d.Represents(fn)
   rep := reps.Represent(band)
   ext, err := mech.ExtensionByType(rep.MimeType)
   if err != nil {
      return err
   }
   if d.info {
      for _, each := range reps {
         if each.Bandwidth == rep.Bandwidth {
            fmt.Print("!")
         }
         fmt.Println(each)
      }
      if d.key == "" {
         err := d.setKey()
         if err != nil {
            return err
         }
      }
      // github.com/edgeware/mp4ff/issues/146
      fmt.Printf("mp4decrypt --key 1:%v enc%v dec%v\n", d.key, ext, ext)
   } else {
      file, err := os.Create("enc" + ext)
      if err != nil {
         return err
      }
      defer file.Close()
      init, err := rep.Initialization(d.URL)
      if err != nil {
         return err
      }
      fmt.Println("GET", init)
      res, err := http.Get(init.String())
      if err != nil {
         return err
      }
      defer res.Body.Close()
      if res.StatusCode != http.StatusOK {
         return errors.New(res.Status)
      }
      if _, err := file.ReadFrom(res.Body); err != nil {
         return err
      }
      media, err := rep.Media(d.URL)
      if err != nil {
         return err
      }
      pro := format.ProgressChunks(file, len(media))
      for _, addr := range media {
         fmt.Println(addr)
         res, err := http.Get(addr.String())
         if err != nil {
            return err
         }
         if res.StatusCode != http.StatusOK {
            return errors.New(res.Status)
         }
         pro.AddChunk(res.ContentLength)
         if _, err := io.Copy(pro, res.Body); err != nil {
            return err
         }
         if err := res.Body.Close(); err != nil {
            return err
         }
      }
   }
   return nil
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
   d.URL = res.Request.URL
   d.Period, err = dash.NewPeriod(res.Body)
   if err != nil {
      return err
   }
   if err := d.download(audio, dash.Audio); err != nil {
      return err
   }
   return d.download(video, dash.Video)
}

func (d *downloader) setKey() error {
   privateKey, err := os.ReadFile(d.pem)
   if err != nil {
      return err
   }
   clientID, err := os.ReadFile(d.client)
   if err != nil {
      return err
   }
   kID, err := d.Protection().KID()
   if err != nil {
      return err
   }
   mod, err := widevine.NewModule(privateKey, clientID, kID)
   if err != nil {
      return err
   }
   addr := d.DASH().Key_Systems.Widevine.License_URL
   keys, err := mod.Post(addr, d.Header())
   if err != nil {
      return err
   }
   d.key = keys.Content().String()
   return nil
}

type downloader struct {
   *amc.Playback
   *dash.Period
   *url.URL
   client string
   info bool
   key string
   pem string
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

func main() {
   home, err := os.UserHomeDir()
   if err != nil {
      panic(err)
   }
   var down downloader
   // a
   var address string
   flag.StringVar(&address, "a", "", "address")
   // b
   var nid int64
   flag.Int64Var(&nid, "b", 0, "NID")
   // c
   down.client = filepath.Join(home, "mech/client_id.bin")
   flag.StringVar(&down.client, "c", down.client, "client ID")
   // e
   var email string
   flag.StringVar(&email, "e", "", "email")
   // f
   // amcplus.com/shows/orphan-black/episodes/season-1-natural-selection--1011153
   var video int64
   flag.Int64Var(&video, "f", 1_999_999, "video bandwidth")
   // g
   // amcplus.com/shows/orphan-black/episodes/season-1-natural-selection--1011153
   var audio int64
   flag.Int64Var(&audio, "g", 127_000, "audio bandwidth")
   // i
   flag.BoolVar(&down.info, "i", false, "information")
   // k
   down.pem = filepath.Join(home, "mech/private_key.pem")
   flag.StringVar(&down.pem, "k", down.pem, "private key")
   // p
   var password string
   flag.StringVar(&password, "p", "", "password")
   // v
   var verbose bool
   flag.BoolVar(&verbose, "v", false, "verbose")
   flag.Parse()
   if verbose {
      amc.LogLevel = 1
   }
   if email != "" {
      err := doLogin(email, password)
      if err != nil {
         panic(err)
      }
   } else if nid >= 1 || address != "" {
      err := down.doDASH(address, nid, video, audio)
      if err != nil {
         panic(err)
      }
   } else {
      flag.Usage()
   }
}
