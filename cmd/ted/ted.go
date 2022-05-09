package main

import (
   "flag"
   "fmt"
   "github.com/89z/format"
   "github.com/89z/mech/ted"
   "net/http"
   "os"
   "path"
)

func process(slug string, info bool, bitrate int64) error {
   talk, err := ted.NewTalkResponse(slug)
   if err != nil {
      return err
   }
   for _, vid := range talk.Downloads.Video {
      if info {
         fmt.Println(vid)
      } else if vid.Bitrate == bitrate {
         addr := vid.GetURL()
         fmt.Println("GET", addr)
         res, err := http.Get(addr)
         if err != nil {
            return err
         }
         defer res.Body.Close()
         file, err := os.Create(path.Base(addr))
         if err != nil {
            return err
         }
         defer file.Close()
         pro := format.NewProgress(res)
         if _, err := file.ReadFrom(pro); err != nil {
            return err
         }
      }
   }
   return nil
}

func main() {
   // b
   var bitrate int64
   flag.Int64Var(&bitrate, "b", 180, "bitrate")
   // i
   var info bool
   flag.BoolVar(&info, "i", false, "info only")
   // s
   var slug string
   flag.StringVar(&slug, "s", "", "slug")
   // v
   var verbose bool
   flag.BoolVar(&verbose, "v", false, "verbose")
   flag.Parse()
   if verbose {
      ted.LogLevel = 1
   }
   if slug != "" {
      err := process(slug, info, bitrate)
      if err != nil {
         panic(err)
      }
   } else {
      flag.Usage()
   }
}
