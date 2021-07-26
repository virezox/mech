package main

import (
   "github.com/Nr90/imgsim"
   "image/jpeg"
   "os"
)

func hash(name string) (imgsim.Hash, error) {
   f, err := os.Open(name)
   if err != nil {
      return 0, err
   }
   defer f.Close()
   i, err := jpeg.Decode(f)
   if err != nil {
      return 0, err
   }
   return imgsim.AverageHash(i), nil
}

func main() {
   a, err := hash("mb.jpg")
   if err != nil {
      panic(err)
   }
   b, err := hash("crop.jpg")
   if err != nil {
      panic(err)
   }
   d := imgsim.Distance(a, b)
   println(d)
}
