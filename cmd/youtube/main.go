package main

import (
   "flag"
   "fmt"
   "github.com/89z/format"
   "github.com/89z/mech/youtube"
   "os"
)

func main() {
   var construct bool
   flag.BoolVar(&construct, "c", false, "OAuth construct request")
   var embed bool
   flag.BoolVar(&embed, "e", false, "use embedded player")
   var refresh bool
   flag.BoolVar(&refresh, "r", false, "OAuth token refresh")
   var verbose bool
   flag.BoolVar(&verbose, "v", false, "verbose")
   var exchange bool
   flag.BoolVar(&exchange, "x", false, "OAuth token exchange")
   // choice
   var choose choice
   flag.BoolVar(&choose.info, "i", false, "video information")
   flag.BoolVar(
      &choose.useFormats,
      "formats",
      false,
      "use formats instead of adaptiveFormats",
   )
   choose.itags = make(map[string]bool)
   flag.Func("f", "formats", func(itag string) error {
      choose.itags[itag] = true
      return nil
   })
   // Parse
   flag.Parse()
   if verbose {
      format.Log.Level = 1
   }
   switch {
   case exchange:
      oauth, err := youtube.NewOAuth()
      if err != nil {
         panic(err)
      }
      fmt.Println(oauth)
      fmt.Scanln()
      exc, err := oauth.Exchange()
      if err != nil {
         panic(err)
      }
      cache, err := os.UserCacheDir()
      if err != nil {
         panic(err)
      }
      if err := exc.Create(cache + "/mech/youtube.json"); err != nil {
         panic(err)
      }
   case refresh:
      cache, err := os.UserCacheDir()
      if err != nil {
         panic(err)
      }
      exc, err := youtube.OpenExchange(cache + "/mech/youtube.json")
      if err != nil {
         panic(err)
      }
      if err := exc.Refresh(); err != nil {
         panic(err)
      }
      if err := exc.Create(cache + "/mech/youtube.json"); err != nil {
         panic(err)
      }
   case flag.NArg() == 1:
      id := flag.Arg(0)
      auth := youtube.Key
      if construct {
         cache, err := os.UserCacheDir()
         if err != nil {
            panic(err)
         }
         exc, err := youtube.OpenExchange(cache + "/mech/youtube.json")
         if err != nil {
            panic(err)
         }
         auth = exc.Auth()
      }
      client := youtube.Android
      if embed {
         client = youtube.Embed
      }
      play, err := youtube.NewPlayer(id, auth, client)
      if err != nil {
         panic(err)
      }
      if choose.useFormats {
         err := choose.formats(play, id)
         if err != nil {
            panic(err)
         }
      } else {
         err := choose.adaptiveFormats(play, id)
         if err != nil {
            panic(err)
         }
      }
   default:
      fmt.Println("youtube [flags] [video ID]")
      flag.PrintDefaults()
   }
}
