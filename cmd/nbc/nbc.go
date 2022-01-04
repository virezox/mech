package main

import (
   "flag"
   "fmt"
   "github.com/89z/format"
   "github.com/89z/mech/nbc"
   "net/http"
   "os"
   "strconv"
   "strings"
)

type choice struct {
   info bool
   formats map[string]bool
}

func main() {
   cHLS := choice{
      formats: make(map[string]bool),
   }
   flag.BoolVar(&cHLS.info, "i", false, "info")
   flag.Func("f", "formats", func(id string) error {
      cHLS.formats[id] = true
      return nil
   })
   var verbose bool
   flag.BoolVar(&verbose, "v", false, "verbose")
   flag.Parse()
   if flag.NArg() != 1 {
      fmt.Println("nbc [flags] [GUID]")
      flag.PrintDefaults()
      return
   }
   guid := flag.Arg(0)
   nGUID, err := nbc.Parse(guid)
   if err != nil {
      panic(err)
   }
   if verbose {
      nbc.Log.Level = 1
   }
   if err := cHLS.HLS(nGUID); err != nil {
      panic(err)
   }
}

func video(guid uint64, info bool) (*nbc.Video, error) {
   if info {
      return nil, nil
   }
   return nbc.NewVideo(guid)
}

func (c choice) HLS(guid uint64) error {
   vod, err := nbc.NewAccessVOD(guid)
   if err != nil {
      return err
   }
   forms, err := vod.Manifest()
   if err != nil {
      return err
   }
   vid, err := video(guid, c.info)
   if err != nil {
      return err
   }
   for id, form := range forms {
      switch {
      case c.info:
         fmt.Print("ID:", id)
         fmt.Print(" BANDWIDTH:", form["BANDWIDTH"])
         fmt.Print(" CODECS:", form["CODECS"])
         fmt.Print(" RESOLUTION:", form["RESOLUTION"])
         fmt.Println()
      case c.formats[strconv.Itoa(id)]:
         loc := form["URI"]
         fmt.Println("GET", loc)
         res, err := http.Get(loc)
         if err != nil {
            return err
         }
         defer res.Body.Close()
         srcs, err := nbc.Decode(res.Body, "")
         if err != nil {
            return err
         }
         name := vid.Name() + "-" + form["RESOLUTION"] + ".mp4"
         dst, err := os.Create(strings.Map(format.Clean, name))
         if err != nil {
            return err
         }
         defer dst.Close()
         total := len(srcs)
         for value, src := range srcs {
            loc := src["URI"]
            format.PercentInt(os.Stdout, value, total)
            os.Stdout.WriteString(" ")
            format.Trim(os.Stdout, loc)
            os.Stdout.WriteString("\n")
            res, err := http.Get(loc)
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
