package main

import (
   "github.com/corona10/goimagehash"
   "image/jpeg"
   "os"
)

func hash(name string) (*goimagehash.ImageHash, error) {
   f, err := os.Open(name)
   if err != nil {
      return nil, err
   }
   defer f.Close()
   i, err := jpeg.Decode(f)
   if err != nil {
      return nil, err
   }
   return goimagehash.AverageHash(i)
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
   d, err := a.Distance(b)
   if err != nil {
      panic(err)
   }
   println(d)
}
