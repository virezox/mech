package main

import (
   "flag"
   "fmt"
   "github.com/89z/mech/bandcamp"
   "net/http"
   "os"
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
         addr, ok := track.MP3_128()
         if ok {
            fmt.Println("GET", addr)
            res, err := http.Get(addr)
            if err != nil {
               return err
            }
            defer res.Body.Close()
            file, err := os.Create(track.Name(data))
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
   var choice flags
   // a
   var address string
   flag.StringVar(&address, "a", "", "address")
   // i
   flag.BoolVar(&choice.info, "i", false, "info only")
   // s
   flag.DurationVar(&choice.sleep, "s", time.Second, "sleep")
   // v
   var verbose bool
   flag.BoolVar(&verbose, "v", false, "verbose")
   flag.Parse()
   if verbose {
      bandcamp.LogLevel = 1
   }
   if address != "" {
      data, err := bandcamp.NewDataTralbum(address)
      if err != nil {
         panic(err)
      }
      if err := choice.process(data); err != nil {
         panic(err)
      }
   } else {
      flag.PrintDefaults()
   }
}
