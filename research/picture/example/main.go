package main

import (
   "github.com/89z/mech/picture"
   "os"
)

func main() {
   f, err := os.Open("Lenna50.jpg")
   if err != nil {
      panic(err)
   }
   defer f.Close()
   picture.Decode(f)
}
