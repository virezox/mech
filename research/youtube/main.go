package main

import (
   "flag"
   "fmt"
   "github.com/89z/mech/youtube"
)

func main() {
   // a
   var address string
   flag.StringVar(&address, "a", "", "address")
   // b
   var videoID string
   flag.StringVar(&videoID, "b", "", "video ID")
   // c
   var construct bool
   flag.BoolVar(&construct, "c", false, "OAuth construct request")
   // e
   var embed bool
   flag.BoolVar(&embed, "e", false, "use embedded player")
   // f
   // g
   // i
   var info bool
   flag.BoolVar(&info, "i", false, "information")
   // r
   var refresh bool
   flag.BoolVar(&refresh, "r", false, "OAuth token refresh")
   // v
   var verbose bool
   flag.BoolVar(&verbose, "v", false, "verbose")
   // x
   var exchange bool
   flag.BoolVar(&exchange, "x", false, "OAuth token exchange")
   flag.Parse()
   if verbose {
      youtube.LogLevel = 1
   }
   if exchange {
      err := doExchange()
      if err != nil {
         panic(err)
      }
   } else if refresh {
      err := doRefresh()
      if err != nil {
         panic(err)
      }
   } else if videoID != "" || address != "" {
      id, err := getID(videoID, address)
      if err != nil {
         panic(err)
      }
      fmt.Println(id)
   } else {
      flag.Usage()
   }
}
