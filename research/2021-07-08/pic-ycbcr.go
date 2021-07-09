package main

import (
   "fmt"
   "image"
   "image/jpeg"
   "os"
)

func ycbcrToRgb(y, cb, cr uint8) (uint8, uint8, uint8) {
   fy, fcb, fcr := float64(y), float64(cb), float64(cr)
   r := fy + 1.402 * (fcr - 128)
   g := fy - 0.34414*(fcb-128) - 0.71414*(fcr-128)
   b := fy + 1.77200*(fcb-128)
   return uint8(r), uint8(g), uint8(b)
}

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

func absDiff(a, b uint8) uint8 {
   if a < b {
      return b - a
   }
   return a - b
}

func main() {
   a, err := newPicture("Lenna100.jpg")
   if err != nil {
      panic(err)
   }
   b, err := newPicture("Lenna50.jpg")
   if err != nil {
      panic(err)
   }
   fmt.Println(a.difference(b))
}
