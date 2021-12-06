package main

import (
   "flag"
   "fmt"
   "github.com/89z/mech/nbc"
   "github.com/89z/parse/m3u"
   "net/http"
   "os"
   "strconv"
)

type choice struct {
   format bool
   ids map[string]bool
}

func main() {
   cHLS := choice{
      ids: make(map[string]bool),
   }
   flag.BoolVar(&cHLS.format, "hf", false, "HLS formats")
   flag.Func("h", "HLS IDs", func(id string) error {
      cHLS.ids[id] = true
      return nil
   })
   flag.Parse()
   if flag.NArg() != 1 {
      fmt.Println("nbc [flags] [GUID]")
      flag.PrintDefaults()
      return
   }
   guid := flag.Arg(0)
   if !nbc.Valid(guid) {
      panic("invalid GUID")
   }
   nGUID, err := strconv.Atoi(guid)
   if err != nil {
      panic(err)
   }
   if err := cHLS.HLS(nGUID); err != nil {
      panic(err)
   }
}

func (c choice) HLS(guid int) error {
   vod, err := nbc.NewAccessVOD(guid)
   if err != nil {
      return err
   }
   forms, err := vod.Manifest()
   if err != nil {
      return err
   }
   vid, err := nbc.NewVideo(guid)
   if err != nil {
      return err
   }
   for id, form := range forms {
      switch {
      case c.format:
         fmt.Print("ID:", id)
         fmt.Print(" BANDWIDTH:", form["BANDWIDTH"])
         fmt.Print(" CODECS:", form["CODECS"])
         fmt.Print(" RESOLUTION:", form["RESOLUTION"])
         fmt.Println()
      case c.ids[strconv.Itoa(id)]:
         addr := form["URI"]
         fmt.Println("GET", addr)
         res, err := http.Get(addr)
         if err != nil {
            return err
         }
         defer res.Body.Close()
         srcs, err := m3u.Decode(res.Body, "")
         if err != nil {
            return err
         }
         name := vid.Name() + "-" + form["RESOLUTION"] + ".mp4"
         dst, err := os.Create(name)
         if err != nil {
            return err
         }
         defer dst.Close()
         for key, src := range srcs {
            addr := src["URI"]
            fmt.Println(len(srcs)-key, "GET", addr)
            res, err := http.Get(addr)
            if err != nil {
               return err
            }
            defer res.Body.Close()
            if _, err := dst.ReadFrom(res.Body); err != nil {
               return err
            }
         }
      }
   }
   return nil
}
