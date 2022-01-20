package main

import (
   "flag"
   "fmt"
   "github.com/89z/format"
   "github.com/89z/mech/bandcamp"
   "net/http"
   "os"
   "strings"
   "time"
)

type flags struct {
   info bool
   sleep time.Duration
}

func (f flags) process(data *bandcamp.DataTralbum) error {
   for _, track := range data.TrackInfo {
      if f.info {
         fmt.Printf("%+v\n", track)
      } else {
         addr, ok := track.File.MP3_128()
         if ok {
            fmt.Println("GET", addr)
            res, err := http.Get(addr)
            if err != nil {
               return err
            }
            defer res.Body.Close()
            name := data.Artist + "-" + track.Title + ".mp3"
            file, err := os.Create(strings.Map(format.Clean, name))
            if err != nil {
               return err
            }
            defer file.Close()
            if _, err := file.ReadFrom(res.Body); err != nil {
               return err
            }
            time.Sleep(f.sleep)
         }
      }
   }
   return nil
}

func main() {
   var (
      choice flags
      verbose bool
   )
   flag.BoolVar(&choice.info, "i", false, "info only")
   flag.DurationVar(&choice.sleep, "s", time.Second, "sleep")
   flag.BoolVar(&verbose, "v", false, "verbose")
   flag.Parse()
   if verbose {
      bandcamp.LogLevel = 1
   }
   if flag.NArg() == 1 {
      addr := flag.Arg(0)
      data, err := bandcamp.NewDataTralbum(addr)
      if err != nil {
         panic(err)
      }
      if err := choice.process(data); err != nil {
         panic(err)
      }
   } else {
      fmt.Println("bandcamp [flags] [track or album]")
      flag.PrintDefaults()
   }
}
