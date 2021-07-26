package main

import (
   "github.com/corona10/goimagehash"
   "image"
   "image/jpeg"
   "net/http"
)

const (
   urlA =
      "https://ia800309.us.archive.org/9/items" +
      "/mbid-a40cb6e9-c766-37c4-8677-7eb51393d5a1" +
      "/mbid-a40cb6e9-c766-37c4-8677-7eb51393d5a1-9261666555.jpg"
   urlB = "http://i.ytimg.com/vi/11Bvzknjo2Q/hqdefault.jpg"
   urlC = "http://i.ytimg.com/vi/jCMi9_6vnxk/hqdefault.jpg"
)

func hash(addr string, crop bool) (*goimagehash.ImageHash, error) {
   r, err := http.Get(addr)
   if err != nil {
      return nil, err
   }
   defer r.Body.Close()
   i, err := jpeg.Decode(r.Body)
   if err != nil {
      return nil, err
   }
   if crop {
      height := 270
      x0 := (480 - height) / 2
      y0 := (360 - height) / 2
      r := image.Rect(x0, y0, x0 + height, y0 + height)
      i = i.(*image.YCbCr).SubImage(r)
   }
   return goimagehash.PerceptionHash(i)
}

func main() {
   a, err := hash(urlA, false)
   if err != nil {
      panic(err)
   }
   b, err := hash(urlB, true)
   if err != nil {
      panic(err)
   }
   c, err := hash(urlC, true)
   if err != nil {
      panic(err)
   }
   ab, err := a.Distance(b)
   if err != nil {
      panic(err)
   }
   ac, err := a.Distance(c)
   if err != nil {
      panic(err)
   }
   println(ab, ac)
}
