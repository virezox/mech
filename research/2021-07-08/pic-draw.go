package main

import (
   "fmt"
   "image"
   "image/draw"
   "image/jpeg"
   "os"
)

type picture struct {
   *image.RGBA
}

func newPicture(name string) (*picture, error) {
   f, err := os.Open(name)
   if err != nil {
      return nil, err
   }
   defer f.Close()
   src, err := jpeg.Decode(f)
   if err != nil {
      return nil, err
   }
   b := src.Bounds()
   dst := image.NewRGBA(image.Rect(0, 0, b.Dx(), b.Dy()))
   draw.Draw(dst, dst.Bounds(), src, b.Min, draw.Src)
   return &picture{dst}, nil
}

func absDiff(a, b uint8) uint8 {
   if a < b {
      return b - a
   }
   return a - b
}

func (a picture) difference(b *picture) float64 {
   var part, total uint64
   for i := range a.Pix {
      part += uint64(absDiff(a.Pix[i], b.Pix[i]))
      total += 0xFF
   }
   return float64(part) / float64(total)
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
