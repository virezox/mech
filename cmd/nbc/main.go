package main

import (
   "flag"
   "fmt"
   "github.com/89z/mech/nbc"
)

func main() {
   // f
   var form int64
   flag.Int64Var(&form, "f", 0, "format")
   // i
   var info bool
   flag.BoolVar(&info, "i", false, "info")
   // n
   var guid int64
   flag.Int64Var(&guid, "n", 0, "GUID")
   // v
   var verbose bool
   flag.BoolVar(&verbose, "v", false, "verbose")
   flag.Parse()
   if verbose {
      nbc.LogLevel = 1
   }
   if guid >= 1 {
      vid, err := video(guid, info)
      if err != nil {
         panic(err)
      }
      vod, err := nbc.NewAccessVOD(guid)
      if err != nil {
         panic(err)
      }
      streams, err := vod.Streams()
      if err != nil {
         panic(err)
      }
      for _, stream := range streams {
         if info {
            fmt.Println(stream)
         } else if stream.ID == form {
            err := download(vid, stream)
            if err != nil {
               panic(err)
            }
         }
      }
   } else {
      flag.Usage()
   }
}
