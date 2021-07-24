package main

import (
   "golang.org/x/image/draw"
   "image"
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
   fb, err := os.Create("scale.jpg")
   if err != nil {
      panic(err)
   }
   defer fb.Close()
   ib := image.NewRGBA(image.Rect(0, 0, 720, 720))
   draw.CatmullRom.Scale(ib, ib.Rect, ia, ia.Bounds(), draw.Over, nil)
   jpeg.Encode(fb, ib, &jpeg.Options{100})
}
