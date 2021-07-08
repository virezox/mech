package main

import (
   "fmt"
   "image"
   "image/jpeg"
   "os"
)

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

func (a picture) difference(b *picture) float32 {
   var part, total float32
   for i := range a.Y {
      part += float32(absDiff(a.Y[i], b.Y[i]))
      total += 0xFF
   }
   for i := range a.Cb {
      part += float32(absDiff(a.Cb[i], b.Cb[i]))
      total += 0xFF
   }
   for i := range a.Cr {
      part += float32(absDiff(a.Cr[i], b.Cr[i]))
      total += 0xFF
   }
   return part / total
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
