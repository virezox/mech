package main

import (
   "flag"
   "fmt"
   "github.com/89z/mech"
   "github.com/89z/mech/youtube"
   "os"
)

func main() {
   var construct, embed, exchange, info, refresh, verbose bool
   down := make(choice)
   flag.BoolVar(&construct, "c", false, "OAuth construct request")
   flag.BoolVar(&embed, "e", false, "use embedded player")
   flag.Func("f", "formats", func(format string) error {
      down[format] = true
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
   if verbose {
      mech.LogLevel = 2
   }
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
      auth := youtube.Key
      if construct {
         var exc youtube.Exchange
         err := authConstruct(&exc)
         if err != nil {
            panic(err)
         }
         auth = youtube.Auth{"Authorization", "Bearer " + exc.Access_Token}
      }
      client := youtube.Android
      if embed {
         client = youtube.Embed
      }
      play, err := youtube.NewPlayer(id, auth, client)
      if err != nil {
         panic(err)
      }
      if len(play.StreamingData.AdaptiveFormats) == 0 {
         panic(play.PlayabilityStatus)
      }
      if info {
         err := infoPath(play, id)
         if err != nil {
            panic(err)
         }
      } else {
         err := down.download(play, id)
         if err != nil {
            panic(err)
         }
      }
   }
}
