package main

import (
   "github.com/89z/mech/youtube"
   "image/jpeg"
   "os"
)

func main() {
   i := youtube.Image{Width: 1280, Height: 720, SubHeight: 720}
   r, err := os.Open("yt.jpg")
   if err != nil {
      panic(err)
   }
   defer r.Close()
   m, err := jpeg.Decode(r)
   if err != nil {
      panic(err)
   }
   w, err := os.Create("crop.jpg")
   if err != nil {
      panic(err)
   }
   defer w.Close()
   jpeg.Encode(w, i.SubImage(m), &jpeg.Options{100})
}
