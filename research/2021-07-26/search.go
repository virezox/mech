package main

import (
   "fmt"
   "github.com/89z/mech/youtube"
   "github.com/Nr90/imgsim"
   "image/jpeg"
   "net/http"
   "os"
   "time"
)

var ids = []string{
   "11Bvzknjo2Q", // good
   "2bDfLtRqKFs",
   "2hqqyncPrd0",
   "4FnsdJkUBhk",
   "8jCbvqFqftg",
   "AvEm3a20Yc4",
   "B3szYRzZqp4",
   "EGrv5FND4GY",
   "Nw6k8JdZmo8", // good
   "Osh3waD3pVU",
   "XbUOX4lr9Bw",
   "ZXNscpJIzQs",
   "_vhnMkcK5yo",
   "fivLqoP0WhU",
   "jCMi9_6vnxk",
   "jt5tRaV3iY0",
   "m3TqulO8vXA",
   "nGj5N9Ll9pI", // good
   "qX1uuYWtc7A",
   "uHrWHXL065g",
   "uIeoAzVUEJw",
   "vJMjpX4Ck2o", // good
}

const (
   urlA =
      "https://ia800309.us.archive.org/9/items" +
      "/mbid-a40cb6e9-c766-37c4-8677-7eb51393d5a1" +
      "/mbid-a40cb6e9-c766-37c4-8677-7eb51393d5a1-9261666555.jpg"
   urlB = "http://i.ytimg.com/vi/11Bvzknjo2Q/hqdefault.jpg"
   urlC = "http://i.ytimg.com/vi/jCMi9_6vnxk/hqdefault.jpg"
)

var forms = []youtube.Image{
   {120, 90, 68, "default", youtube.JPG},
   {320, 180, 180, "mqdefault", youtube.JPG},
   {480, 360, 270, "hqdefault", youtube.JPG},
}

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
   f, err := os.Open("mb.jpg")
   if err != nil {
      panic(err)
   }
   defer f.Close()
   i, err := jpeg.Decode(f)
   if err != nil {
      panic(err)
   }
   a, err := goimagehash.AverageHash(i)
   if err != nil {
      panic(err)
   }
   for _, id := range ids {
      for _, form := range forms {
         res, err := http.Get(form.Address(id))
         if err != nil {
            panic(err)
         }
         defer res.Body.Close()
         i, err := jpeg.Decode(res.Body)
         if err != nil {
            panic(err)
         }
         b, err := goimagehash.AverageHash(form.SubImage(i))
         if err != nil {
            panic(err)
         }
         d, err := a.Distance(b)
         if err != nil {
            panic(err)
         }
         fmt.Println(d, id, form.Base)
         time.Sleep(100 * time.Millisecond)
      }
   }
}
