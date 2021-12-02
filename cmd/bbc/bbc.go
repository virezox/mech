package main

import (
   "flag"
   "fmt"
   "github.com/89z/mech/bbc"
   "strconv"
)

func main() {
   var info bool
   flag.BoolVar(&info, "i", false, "info only")
   flag.Parse()
   if flag.NArg() == 0 {
      fmt.Println("bbc [flags] [ID]")
      flag.PrintDefaults()
      return
   }
   sID := flag.Arg(0)
   if !bbc.Valid(sID) {
      panic("invalid ID")
   }
   id, err := strconv.Atoi(sID)
   if err != nil {
      panic(err)
   }
   video, err := bbc.NewNewsVideo(id)
   if err != nil {
      panic(err)
   }
   con, err := video.Connection()
   if err != nil {
      panic(err)
   }
   fmt.Printf("%+v\n", con)
}
