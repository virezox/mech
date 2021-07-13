package main

import (
   "fmt"
   "image"
   "image/jpeg"
   "math"
   "os"
)

func (a picture) difference(b *picture) float64 {
   var part, total uint64
   r := a.Bounds()
   for y := r.Min.Y; y < r.Max.Y; y++ {
      for x := r.Min.X; x < r.Max.X; x++ {
         aa, bb := a.YCbCrAt(x, y), b.YCbCrAt(x, y)
         r1, g1, b1 := ycbcrToRgb(aa.Y, aa.Cb, aa.Cr)
         r2, g2, b2 := ycbcrToRgb(bb.Y, bb.Cb, bb.Cr)
         part += uint64(absDiff(r1, r2))
         part += uint64(absDiff(g1, g2))
         part += uint64(absDiff(b1, b2))
         total += 0xFF + 0xFF + 0xFF
      }
   }
   return float64(part) / float64(total)
}

func absDiff(a, b uint8) uint8 {
   if a < b {
      return b - a
   }
   return a - b
}

type picture struct {
   *image.YCbCr
}

func newPicture(name string) (*picture, error) {
   f, err := os.Open(name)
   if err != nil {
      return nil, err
   }
   defer f.Close()
   i, err := jpeg.Decode(f)
   if err != nil {
      return nil, err
   }
   return &picture{
      i.(*image.YCbCr),
   }, nil
}

func main() {
   f, err := os.Open("Lenna50.jpg")
   if err != nil {
      panic(err)
   }
   defer f.Close()
   picture, err := jpeg.Decode(f)
   if err != nil {
      panic(err)
   }
   pic := picture.(*image.YCbCr).YCbCrAt(15, 0)
   r1 := float64(pic.Y) + 1.402 * float64(pic.Cr - 128)
   r := math.Min(math.Max(0, math.Round(r1)), 255)
   fmt.Println(r1, r)
}
