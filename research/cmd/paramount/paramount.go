package main

import (
   "fmt"
   "github.com/89z/format"
   "github.com/89z/format/dash"
   "github.com/89z/format/hls"
   "github.com/89z/mech"
   "github.com/89z/mech/paramount"
   "net/http"
   "net/url"
   "os"
   "sort"
)

func setVerbose() {
   paramount.LogLevel = 1
}

type downloader struct {
   *dash.Period
   *paramount.Preview
   *url.URL
   key []byte
}

func (d *downloader) setKey() error {
   sess, err := paramount.NewSession(d.Preview.GUID)
   if err != nil {
      return err
   }
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
   d.key, err = sess.Key(privateKey, clientID, kID)
   if err != nil {
      return err
   }
   return nil
}

func (d *downloader) download(band int64, fn dash.PeriodFunc) error {
   if band == 0 {
      return nil
   }
   reps := d.Represents(fn)
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
         err := d.setKey()
         if err != nil {
            return err
         }
      }
      ext, err := mech.ExtensionByType(rep.MimeType)
      if err != nil {
         return err
      }
      file, err := os.Create(d.Base()+ext)
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

func (d downloader) DASH(video, audio int64) error {
   addr, err := paramount.NewMedia(d.GUID).DASH()
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
   if err := d.download(audio, dash.Audio); err != nil {
      return err
   }
   return d.download(video, dash.Video)
}


func (d downloader) HLS(bandwidth int64) error {
   addr, err := paramount.NewMedia(d.GUID).HLS()
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
      fmt.Println(d.Title)
      for _, each := range master.Streams {
         if each.Bandwidth == stream.Bandwidth {
            fmt.Print("!")
         }
         fmt.Println(each)
      }
   } else {
      return download(stream.RawURI, d.Base())
   }
   return nil
}

func download(addr, base string) error {
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

func newSegment(addr string) (*hls.Segment, error) {
   fmt.Println("GET", addr)
   res, err := http.Get(addr)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   return hls.NewScanner(res.Body).Segment()
}


