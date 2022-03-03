package main

import (
   "fmt"
   "github.com/89z/format/hls"
   "github.com/89z/mech/mtv"
   "net/http"
   "os"
   "sort"
)

func doManifest(addr string, bandwidth int64, info bool) error {
   prop, err := mtv.NewItem(addr).Property()
   if err != nil {
      return err
   }
   top, err := prop.Topaz()
   if err != nil {
      return err
   }
   fmt.Println("GET", top.StitchedStream.Source)
   res, err := http.Get(top.StitchedStream.Source)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   mas, err := hls.NewMaster(res.Request.URL, res.Body)
   if err != nil {
      return err
   }
   sort.Slice(mas.Stream, func(a, b int) bool {
      return mas.Stream[a].Bandwidth < mas.Stream[b].Bandwidth
   })
   if info {
      for _, str := range mas.Stream {
         str.URI = ""
         fmt.Println(str)
      }
   } else {
      uris := mas.URIs(func(str hls.Stream) bool {
         return str.Bandwidth >= bandwidth
      })
      for _, uri := range uris {
         fmt.Println("GET", uri)
         res, err := http.Get(uri)
         if err != nil {
            return err
         }
         defer res.Body.Close()
         seg, err := hls.NewSegment(res.Request.URL, res.Body)
         if err != nil {
            return err
         }
         ext, err := seg.Ext()
         if err != nil {
            return err
         }
         if err := download(seg, "ignore/" + prop.Base() + ext); err != nil {
            return err
         }
      }
   }
   return nil
}

func download(seg *hls.Segment, name string) error {
   fmt.Println("GET", seg.Key.URI)
   res, err := http.Get(seg.Key.URI)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   dec, err := hls.NewDecrypter(res.Body)
   if err != nil {
      return err
   }
   file, err := os.Create(name)
   if err != nil {
      return err
   }
   defer file.Close()
   for i, info := range seg.Info {
      fmt.Println(i, len(seg.Info)-1)
      res, err := http.Get(info.URI)
      if err != nil {
         return err
      }
      if _, err := dec.Copy(file, res.Body); err != nil {
         return err
      }
      if err := res.Body.Close(); err != nil {
         return err
      }
   }
   return nil
}
