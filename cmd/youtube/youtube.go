package main

import (
   "flag"
   "fmt"
   "github.com/89z/mech"
   "github.com/89z/mech/youtube"
   "os"
)

func main() {
   var construct, exchange, refresh bool
   var embed bool
   flag.BoolVar(&construct, "c", false, "OAuth construct request")
   flag.BoolVar(&embed, "e", false, "use embedded player")
   flag.BoolVar(&exchange, "x", false, "OAuth token exchange")
   flag.BoolVar(&refresh, "r", false, "OAuth token refresh")
   var (
      atag, vtag int
      info bool
   )
   flag.BoolVar(&info, "i", false, "info only")
   flag.IntVar(&atag, "a", 0, "audio itag (-1 to skip)")
   flag.IntVar(&vtag, "v", 0, "video itag (-1 to skip)")
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
         err := downloadPath(construct, embed, atag, vtag, id)
         if err != nil {
            panic(err)
         }
      }
   }
}
