package main

import (
   "flag"
   "fmt"
   "github.com/89z/mech/cbc"
   "github.com/89z/rosso/hls"
   "github.com/89z/rosso/os"
   "io"
   "net/url"
   "strings"
)

type flags struct {
   bandwidth int
   email string
   id string
   mech.Flag
   password string
}

func main() {
   var f flags
   // b
   flag.StringVar(&f.id, "b", "", "ID")
   // e
   flag.StringVar(&f.email, "e", "", "email")
   // f
   flag.IntVar(&f.bandwidth, "f", 2052370, "video bandwidth")
   // i
   flag.BoolVar(&f.Info, "i", false, "information")
   // p
   flag.StringVar(&f.password, "p", "", "password")
   flag.Parse()
   if f.email != "" {
      err := f.profile()
      if err != nil {
         panic(err)
      }
   } else if f.id != "" {
      err := f.download()
      if err != nil {
         panic(err)
      }
   } else {
      flag.Usage()
   }
}
