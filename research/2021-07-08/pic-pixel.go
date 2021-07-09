package main

import (
   "fmt"
   "image"
   "image/jpeg"
   "math"
   "os"
)

func getPixel(i *image.YCbCr, re image.Rectangle, x, y int) (float64, float64, float64) {
   iy := (y- re.Min.Y)* i.YStride + (x - re.Min.X)
   ic := (y/2- re.Min.Y/2)* i.CStride + (x/2 - re.Min.X/2)
   y1 := float64(i.Y[iy])
   cb1 := float64(i.Cb[ic]) - 128
   cr1 := float64(i.Cr[ic]) - 128
   r1 := y1 + 1.402*cr1
   g1 := y1 - 0.344136*cb1 - 0.714136*cr1
   b1 := y1 + 1.772*cb1
   r := math.Min(math.Max(0, math.Round(r1)), 255)
   g := math.Min(math.Max(0, math.Round(g1)), 255)
   b := math.Min(math.Max(0, math.Round(b1)), 255)
   return r,g,b
}

func (a picture) difference(b *picture) float64 {
   r := a.Bounds()
   var part float64
   var total uint64
   for y := r.Min.Y; y < r.Max.Y; y++ {
      for x := r.Min.X; x < r.Max.X; x++ {
         r1, g1, b1 := getPixel(a.YCbCr, r, x, y)
         r2, g2, b2 := getPixel(b.YCbCr, r, x, y)
         part += absDiff(r1, r2)
         part += absDiff(g1, g2)
         part += absDiff(b1, b2)
         total += 0xFF + 0xFF + 0xFF
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
