package main

import (
   "flag"
   "fmt"
   "github.com/89z/mech"
   "github.com/89z/mech/youtube"
   "os"
)

func infoPath(id string) error {
   play, err := youtube.NewPlayer(id, youtube.Key, youtube.Mweb)
   if err != nil {
      return err
   }
   if len(play.StreamingData.AdaptiveFormats) == 0 {
      return play.PlayabilityStatus
   }
   fmt.Println("author:", play.Author())
   fmt.Println("title:", play.Title())
   fmt.Println("countries:", play.Countries())
   fmt.Println()
   for _, f := range play.StreamingData.AdaptiveFormats {
      fmt.Printf(
         "itag %v, height %v, %v, %v, %v\n",
         f.Itag, f.Height, f.Bitrate, f.ContentLength, f.MimeType,
      )
   }
   return nil
}

func main() {
   var exchange, info, refresh, verbose bool
   down := choice{
      formats: make(map[string]bool),
   }
   flag.BoolVar(&down.construct, "c", false, "OAuth construct request")
   flag.BoolVar(&down.embed, "e", false, "use embedded player")
   flag.Func("f", "formats", func(format string) error {
      down.formats[format] = true
      return nil
   })
   flag.BoolVar(&info, "i", false, "info")
   flag.BoolVar(&refresh, "r", false, "OAuth token refresh")
   flag.BoolVar(&verbose, "v", false, "verbose")
   flag.BoolVar(&exchange, "x", false, "OAuth token exchange")
   flag.Parse()
   if len(os.Args) == 1 {
      fmt.Println("youtube [flags] [video ID]")
      flag.PrintDefaults()
      return
   }
   mech.Verbose = verbose
   switch {
   case exchange:
      err := authExchange()
      if err != nil {
         panic(err)
      }
   case refresh:
      err := authRefresh()
      if err != nil {
         panic(err)
      }
   default:
      id := flag.Arg(0)
      if !youtube.Valid(id) {
         panic("invalid ID")
      }
      switch {
      case info:
         err := infoPath(id)
         if err != nil {
            panic(err)
         }
      default:
         err := down.download(id)
         if err != nil {
            panic(err)
         }
      }
   }
}
