package main

import (
   "fmt"
   // "github.com/89z/mech/youtube"
   // "image/jpeg"
   // "os"
)

var ids = []string{
   "B3szYRzZqp4",
   "11Bvzknjo2Q",
   "4FnsdJkUBhk",
   "m3TqulO8vXA",
   "nGj5N9Ll9pI",
   "Osh3waD3pVU",
   "fivLqoP0WhU",
   "ZXNscpJIzQs",
   "_vhnMkcK5yo",
   "AvEm3a20Yc4",
   "8jCbvqFqftg",
   "jCMi9_6vnxk",
   "2bDfLtRqKFs",
   "uIeoAzVUEJw",
   "EGrv5FND4GY",
   "Nw6k8JdZmo8",
   "2hqqyncPrd0",
   "uHrWHXL065g",
   "jt5tRaV3iY0",
   "qX1uuYWtc7A",
   "XbUOX4lr9Bw",
   "vJMjpX4Ck2o",
}

func main() {
   fmt.Println(ids)
   /*
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
   */
}
