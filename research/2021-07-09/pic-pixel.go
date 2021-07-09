package main

import (
   "fmt"
   "image"
   "image/jpeg"
   "math"
   "os"
)

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
