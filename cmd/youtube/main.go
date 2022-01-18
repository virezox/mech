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
      play, err := player(construct, embed, flag.Arg(0))
      if err != nil {
         panic(err)
      }
      if err := choose.adaptiveFormats(play); err != nil {
         panic(err)
      }
   default:
      fmt.Println("youtube [flags] [video ID]")
      flag.PrintDefaults()
   }
}
