package main

import (
   "flag"
   "fmt"
   "github.com/89z/mech"
   "github.com/89z/mech/bandcamp"
   "net/http"
   "os"
   "strings"
   "time"
)

func main() {
   var (
      info, verbose bool
      sleep time.Duration
   )
   flag.BoolVar(&info, "i", false, "info only")
   flag.DurationVar(&sleep, "s", time.Second, "sleep")
   flag.BoolVar(&verbose, "v", false, "verbose")
   flag.Parse()
   if flag.NArg() == 0 {
      fmt.Println("bandcamp [flags] [track or album]")
      flag.PrintDefaults()
      return
   }
   bandcamp.Verbose = verbose
   addr := flag.Arg(0)
   item, err := bandcamp.NewItem(addr)
   if err != nil {
      panic(err)
   }
   ta, err := item.Tralbum()
   if err != nil {
      panic(err)
   }
   for _, track := range ta.Tracks {
      if info {
         fmt.Printf("%+v\n", track)
      } else {
         fmt.Println("GET", track.Streaming_URL.MP3_128)
         res, err := http.Get(track.Streaming_URL.MP3_128)
         if err != nil {
            panic(err)
         }
         defer res.Body.Close()
         name := ta.Tralbum_Artist + "-" + track.Title + ".mp3"
         file, err := os.Create(strings.Map(mech.Clean, name))
         if err != nil {
            panic(err)
         }
         defer file.Close()
         if _, err := file.ReadFrom(res.Body); err != nil {
            panic(err)
         }
         time.Sleep(sleep)
      }
   }
}
