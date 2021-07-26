package main

import (
   "fmt"
   "github.com/89z/mech/youtube"
   "github.com/corona10/goimagehash"
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

var forms = []youtube.Image{
   {120, 90, 68, "default", youtube.JPG},
   {320, 180, 180, "mqdefault", youtube.JPG},
   {480, 360, 270, "hqdefault", youtube.JPG},
   {640, 480, 480, "sd1", youtube.JPG},
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
   a, err := goimagehash.PerceptionHash(i)
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
         b, err := goimagehash.PerceptionHash(form.SubImage(i))
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
