package main

import (
   "github.com/89z/youtube"
   "os"
)

func main() {
   v, err := youtube.NewVideo("967pHNZk3OM")
   if err != nil {
      panic(err)
   }
   format, err := v.NewFormat(248)
   if err != nil {
      panic(err)
   }
   file, err := os.Create("file.webm")
   if err != nil {
      panic(err)
   }
   defer file.Close()
   err = format.Write(file, false)
   if err != nil {
      panic(err)
   }
}
