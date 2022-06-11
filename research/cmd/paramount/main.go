package main

import (
   "flag"
   "fmt"
   "github.com/89z/format"
   "github.com/89z/format/dash"
   "github.com/89z/format/hls"
   "github.com/89z/mech"
   "github.com/89z/mech/paramount"
   "net/http"
   "net/url"
   "os"
   "path/filepath"
   "sort"
)

func newSegment(addr string) (*hls.Segment, error) {
   fmt.Println("GET", addr)
   res, err := http.Get(addr)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   return hls.NewScanner(res.Body).Segment()
}

func main() {
   home, err := os.UserHomeDir()
   if err != nil {
      panic(err)
   }
   var d download
   // b
   flag.StringVar(&d.mediaID, "b", "", "media ID")
   // c
   d.clientPath = filepath.Join(home, "mech/client_id.bin")
   flag.StringVar(&d.clientPath, "c", d.clientPath, "client ID")
   // d
   var isDASH bool
   flag.BoolVar(&isDASH, "d", false, "DASH download")
   // f
   var videoRate int64
   flag.Int64Var(&videoRate, "f", 1611000, "video bandwidth")
   // g
   var audioRate int64
   flag.Int64Var(&audioRate, "g", 999999, "audio bandwidth")
   // i
   flag.BoolVar(&d.info, "i", false, "information")
   // k
   d.keyPath = filepath.Join(home, "mech/private_key.pem")
   flag.StringVar(&d.keyPath, "k", d.keyPath, "private key")
   // v
   var verbose bool
   flag.BoolVar(&verbose, "v", false, "verbose")
   flag.Parse()
   if verbose {
      setVerbose()
   }
   if d.mediaID != "" {
      err := d.setBase()
      if err != nil {
         panic(err)
      }
      if isDASH {
         err := d.getDASH(videoRate, audioRate)
         if err != nil {
            panic(err)
         }
      } else {
         err := d.getHLS(videoRate)
         if err != nil {
            panic(err)
         }
      }
   } else {
      flag.Usage()
   }
}

func hls_down(addr, base string) error {
   seg, err := newSegment(addr)
   if err != nil {
      return err
   }
   fmt.Println("GET", seg.RawKey)
   res, err := http.Get(seg.RawKey)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   block, err := hls.NewCipher(res.Body)
   if err != nil {
      return err
   }
   file, err := os.Create(base + hls.TS)
   if err != nil {
      return err
   }
   defer file.Close()
   pro := format.ProgressChunks(file, len(seg.Protected))
   for _, addr := range seg.Protected {
      res, err := http.Get(addr)
      if err != nil {
         return err
      }
      pro.AddChunk(res.ContentLength)
      if _, err := block.Copy(pro, res.Body, nil); err != nil {
         return err
      }
      if err := res.Body.Close(); err != nil {
         return err
      }
   }
   return nil
}

func (d download) hls_info(bandwidth int64) error {
   addr, err := d.hls_url()
   if err != nil {
      return err
   }
   fmt.Println("GET", addr)
   res, err := http.Get(addr.String())
   if err != nil {
      return err
   }
   defer res.Body.Close()
   master, err := hls.NewScanner(res.Body).Master()
   if err != nil {
      return err
   }
   sort.Slice(master.Streams, func(a, b int) bool {
      return master.Streams[a].Bandwidth < master.Streams[b].Bandwidth
   })
   stream := master.Streams.GetBandwidth(bandwidth)
   if d.info {
      for _, each := range master.Streams {
         if each.Bandwidth == stream.Bandwidth {
            fmt.Print("!")
         }
         fmt.Println(each)
      }
   } else {
      return download(stream.RawURI, d.base)
   }
   return nil
}

type download struct {
   *dash.Period
   *url.URL
   base string
   info bool
   key []byte
   mediaID string
   path struct {
      client string
      key string
   }
}

func (d *download) dash_down(band int64, fn dash.PeriodFunc) error {
   if band == 0 {
      return nil
   }
   reps := d.Period.Represents(fn)
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
         privateKey, err := os.ReadFile(d.path.key)
         if err != nil {
            return err
         }
         clientID, err := os.ReadFile(d.path.client)
         if err != nil {
            return err
         }
         keyID, err := d.Period.Protection().KeyID()
         if err != nil {
            return err
         }
         if err := d.setKey(); err != nil {
            return err
         }
      }
      ext, err := mech.ExtensionByType(rep.MimeType)
      if err != nil {
         return err
      }
      file, err := os.Create(d.base + ext)
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
      if _, err := file.ReadFrom(res.Body); err != nil {
         return err
      }
      media, err := rep.Media(d.URL)
      if err != nil {
         return err
      }
      pro := format.ProgressChunks(file, len(media))
      for _, addr := range media {
         res, err := http.Get(addr.String())
         if err != nil {
            return err
         }
         pro.AddChunk(res.ContentLength)
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

/////////////////////////////////////////////////////////

func (d download) dash_info(videoRate, audioRate int64) error {
   addr, err := paramount.NewMedia(d.mediaID).DASH()
   if err != nil {
      return err
   }
   fmt.Println("GET", addr)
   res, err := http.Get(addr.String())
   if err != nil {
      return err
   }
   defer res.Body.Close()
   d.URL = addr
   d.Period, err = dash.NewPeriod(res.Body)
   if err != nil {
      return err
   }
   if err := d.download(audioRate, dash.Audio); err != nil {
      return err
   }
   return d.download(videoRate, dash.Video)
}
