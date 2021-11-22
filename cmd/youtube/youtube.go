package main

import (
   "flag"
   "fmt"
   "github.com/89z/mech"
   "github.com/89z/mech/youtube"
   "os"
)

func infoPath(id string) error {
   p, err := youtube.NewPlayer(id, youtube.Key, youtube.Mweb)
   if err != nil {
      return err
   }
   fmt.Println("author:", p.Author())
   fmt.Println("title:", p.Title())
   fmt.Println("countries:", p.Countries())
   fmt.Println()
   for _, f := range p.StreamingData.AdaptiveFormats {
      fmt.Printf(
         "itag %v, height %v, %v, %v, %v\n",
         f.Itag, f.Height, f.Bitrate, f.ContentLength, f.MimeType,
      )
   }
   return nil
}

func main() {
   // OAuth
   var exchange, refresh bool
   flag.BoolVar(&exchange, "x", false, "OAuth token exchange")
   flag.BoolVar(&refresh, "r", false, "OAuth token refresh")
   // info
   var info bool
   flag.BoolVar(&info, "i", false, "info")
   // download
   down := choice{
      formats: make(map[string]bool),
   }
   flag.BoolVar(&down.embed, "e", false, "use embedded player")
   flag.BoolVar(&down.construct, "c", false, "OAuth construct request")
   flag.Func("f", "formats", func(format string) error {
      down.formats[format] = true
      return nil
   })
   // parse
   flag.Parse()
   if len(os.Args) == 1 {
      fmt.Println("youtube [flags] [video ID]")
      flag.PrintDefaults()
      return
   }
   mech.Verbose = true
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
      if ! youtube.Valid(id) {
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
