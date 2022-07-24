package main

import (
   "flag"
   "github.com/89z/mech/widevine"
   "github.com/89z/rosso/http"
)

var client = http.Default_Client

type flags struct {
   Client_ID string
   Info bool
   Name string
   Poster widevine.Poster
   Private_Key string
   address string
   bandwidth int64
   codec string
   key string
   base string
}

func main() {
   var f flags
   // a
   flag.StringVar(&f.address, "a", "", "address")
   // f
   flag.Int64Var(&f.bandwidth, "f", 1, "video bandwidth")
   // g
   flag.StringVar(&f.codec, "g", "mp4a", "audio codec")
   // i
   flag.BoolVar(&f.Info, "i", false, "information")
   // k
   flag.StringVar(&f.key, "k", "", "key")
   // o
   flag.StringVar(&f.Name, "o", "ballysports", "output")
   flag.Parse()
   if f.address != "" {
      if err := f.DASH(); err != nil {
         panic(err)
      }
   } else {
      flag.Usage()
   }
}
