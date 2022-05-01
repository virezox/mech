package main

import (
   "flag"
   "fmt"
   "github.com/89z/format"
   "github.com/89z/format/hls"
   "github.com/89z/mech/cbc"
   "io"
   "net/http"
   "os"
   "sort"
)

func main() {
   // b
   var guid int64
   flag.Int64Var(&guid, "b", 0, "GUID")
   // e
   var email string
   flag.StringVar(&email, "e", "", "email")
   // f
   var bandwidth int
   flag.IntVar(&bandwidth, "f", 3_000_000, "target bandwidth")
   // i
   var info bool
   flag.BoolVar(&info, "i", false, "information")
   // p
   var password string
   flag.StringVar(&password, "p", "", "password")
   // v
   var verbose bool
   flag.BoolVar(&verbose, "v", false, "verbose")
   flag.Parse()
   if verbose {
      cbc.LogLevel = 1
   }
   if email != "" {
      err := doLogin(email, password)
      if err != nil {
         panic(err)
      }
   } else if guid >= 1 {
      err := doManifest(guid, bandwidth, info)
      if err != nil {
         panic(err)
      }
   } else {
      flag.Usage()
   }
}
