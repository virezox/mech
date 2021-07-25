package main

import (
   "github.com/Nr90/imgsim"
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
   fb, err := os.Open("crop.jpg")
   if err != nil {
      panic(err)
   }
   defer fb.Close()
   ib, err := jpeg.Decode(fb)
   if err != nil {
      panic(err)
   }
   d := imgsim.Distance(imgsim.AverageHash(ia), imgsim.AverageHash(ib))
   println(d)
}
