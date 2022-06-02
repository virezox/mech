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


