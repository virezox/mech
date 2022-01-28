package main

import (
   "flag"
   "fmt"
   "github.com/89z/mech/nbc"
   "strconv"
)

func main() {
   formats := make(map[string]bool)
   var info, verbose bool
   flag.Func("f", "formats", func(id string) error {
      formats[id] = true
      return nil
   })
   flag.BoolVar(&info, "i", false, "info")
   flag.BoolVar(&verbose, "v", false, "verbose")
   flag.Parse()
   if verbose {
      nbc.LogLevel = 1
   }
   if flag.NArg() == 1 {
      guid := flag.Arg(0)
      nGUID, err := nbc.Parse(guid)
      if err != nil {
         panic(err)
      }
      vid, err := video(nGUID, info)
      if err != nil {
         panic(err)
      }
      vod, err := nbc.NewAccessVOD(nGUID)
      if err != nil {
         panic(err)
      }
      streams, err := vod.Streams()
      if err != nil {
         panic(err)
      }
      for _, stream := range streams {
         switch {
         case info:
            fmt.Println(stream)
         case formats[strconv.FormatInt(stream.ID, 10)]:
            err := download(vid, stream)
            if err != nil {
               panic(err)
            }
         }
      }
   } else {
      fmt.Println("nbc [flags] [GUID]")
      flag.PrintDefaults()
   }
}
