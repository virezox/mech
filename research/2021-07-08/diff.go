package main

import (
   "fmt"
   "image"
   "image/jpeg"
   "os"
)

func absDiff(a, b uint32) uint32 {
   if a < b {
      return b - a
   }
   return a - b
}

func loadJpeg(filename string) (image.Image, error) {
   f, err := os.Open(filename)
   if err != nil {
      return nil, err
   }
   defer f.Close()
   return jpeg.Decode(f)
}

func main() {
   ia, err := loadJpeg("Lenna50.jpg")
   if err != nil {
      panic(err)
   }
   ib, err := loadJpeg("Lenna100.jpg")
   if err != nil {
      panic(err)
   }
   b := ia.Bounds()
   var sum uint32
   for y := b.Min.Y; y < b.Max.Y; y++ {
      for x := b.Min.X; x < b.Max.X; x++ {
         r1, g1, b1, _ := ia.At(x, y).RGBA()
         r2, g2, b2, _ := ib.At(x, y).RGBA()
         sum += absDiff(r1, r2)
         sum += absDiff(g1, g2)
         sum += absDiff(b1, b2)
      }
   }
   nPixels := (b.Max.X - b.Min.X) * (b.Max.Y - b.Min.Y)
   fmt.Println(float64(sum*100)/(float64(nPixels)*0xffff*3))
}
