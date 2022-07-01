package main

import (
   "encoding/xml"
   "flag"
   "fmt"
   "github.com/89z/mech/amc"
   "github.com/89z/mech/widevine"
   "github.com/89z/std/dash"
   "github.com/89z/std/mp4"
   "github.com/89z/std/os"
   "io"
   "net/url"
   "path/filepath"
)

type downloader struct {
   *amc.Playback
   client string
   info bool
   key []byte
   pem string
   url *url.URL
   media dash.Media
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
      amc.Client.Log_Level = 2
   }
   if email != "" {
      err := do_login(email, password)
      if err != nil {
         panic(err)
      }
   } else if nid >= 1 || address != "" {
      err := down.do_DASH(address, nid, video, audio)
      if err != nil {
         panic(err)
      }
   } else {
      flag.Usage()
   }
}
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
   keys, err := mod.Post(d.Playback)
   if err != nil {
      return err
   }
   d.key = keys.Content().Key
   return nil
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
      file, err := os.Create(d.Base()+rep.Ext())
      if err != nil {
         return err
      }
      defer file.Close()
      initial, err := d.url.Parse(rep.Initialization())
      if err != nil {
         return err
      }
      res, err := amc.Client.Redirect(nil).Get(initial.String())
      if err != nil {
         return err
      }
      defer res.Body.Close()
      if d.key == nil {
         err := d.set_key()
         if err != nil {
            return err
         }
      }
      media := rep.Media()
      pro := os.Progress_Chunks(file, len(media))
      dec := mp4.New_Decrypt(pro)
      if err := dec.Init(res.Body); err != nil {
         return err
      }
      for _, raw := range media {
         addr, err := d.url.Parse(raw)
         if err != nil {
            return err
         }
         res, err := amc.Client.Redirect(nil).Level(0).Get(addr.String())
         if err != nil {
            return err
         }
         pro.Add_Chunk(res.ContentLength)
         if d.key != nil {
            err = dec.Segment(res.Body, d.key)
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


