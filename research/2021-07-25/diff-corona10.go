package main

import (
   "github.com/corona10/goimagehash"
   "image/jpeg"
   "os"
)

func main() {
   fa, err := os.Open("mb.jpg")
   if err != nil {
      panic(err)
   }
   defer fa.Close()
   ia, err := jpeg.Decode(fa)
   if err != nil {
      panic(err)
   }
   ha, err := goimagehash.AverageHash(ia)
   if err != nil {
      panic(err)
   }
   fb, err := os.Open("crop.jpg")
   if err != nil {
      panic(err)
   }
   defer fb.Close()
   ib, err := jpeg.Decode(fb)
   if err != nil {
      panic(err)
   }
   hb, err := goimagehash.AverageHash(ib)
   if err != nil {
      panic(err)
   }
   d, err := ha.Distance(hb)
   if err != nil {
      panic(err)
   }
   println(d)
}
