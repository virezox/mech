package main

import (
   "fmt"
   "image"
   "image/jpeg"
   "os"
)

func clampi32(val, min, max int64) int64 {
   if val > max {
      return max
   }
   if val > min {
      return val
   }
   return 0
}

func getPixel(i *image.YCbCr, re image.Rectangle, x, y int) (float64, float64, float64) {
   iy := (y- re.Min.Y)* i.YStride + (x - re.Min.X)
   ic := (y/2- re.Min.Y/2)* i.CStride + (x/2 - re.Min.X/2)
   const (
      max = 255 * 1e5
      inv = 1.0 / max
   )
   y1 := int64(i.Y[iy]) * 1e5
   cb1 := int64(i.Cb[ic]) - 128
   cr1 := int64(i.Cr[ic]) - 128
   r1 := y1 + 140200*cr1
   g1 := y1 - 34414*cb1 - 71414*cr1
   b1 := y1 + 177200*cb1
   r := float64(clampi32(r1, 0, max)) * inv
   g := float64(clampi32(g1, 0, max)) * inv
   b := float64(clampi32(b1, 0, max)) * inv
   return r,g,b
}

func (a picture) difference(b *picture) float64 {
   r := a.Bounds()
   var part float64
   var total int64
   for y := r.Min.Y; y < r.Max.Y; y++ {
      for x := r.Min.X; x < r.Max.X; x++ {
         r1, g1, b1 := getPixel(a.YCbCr, r, x, y)
         r2, g2, b2 := getPixel(b.YCbCr, r, x, y)
         part += absDiff(r1, r2)
         part += absDiff(g1, g2)
         part += absDiff(b1, b2)
         total += 1 + 1 + 1
      }
   }
   return part / float64(total)
}

func absDiff(a, b float64) float64 {
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
