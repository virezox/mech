package main

import (
   "fmt"
   "image"
   "image/jpeg"
   "os"
)


type picture struct {
   image.Image
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
   return &picture{i}, nil
}

func absDiff(a, b uint32) uint32 {
   if a < b {
      return b - a
   }
   return a - b
}

func (a picture) difference(b *picture) float32 {
   var (
      part uint32
      total float32
   )
   r := a.Bounds()
   for x := r.Min.X; x < r.Max.X; x++ {
      for y := r.Min.Y; y < r.Max.Y; y++ {
         r1, g1, b1, _ := a.At(x, y).RGBA()
         r2, g2, b2, _ := b.At(x, y).RGBA()
         part += absDiff(r1, r2) + absDiff(g1, g2) + absDiff(b1, b2)
         total += 0xFFFF + 0xFFFF + 0xFFFF
      }
   }
   return float32(part) / total
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
