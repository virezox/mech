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
   // A
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
   // B
   for _, id := range ids {
      // crop
      img := youtube.Image{120, 90, 68, "default", youtube.JPG}
      res, err := http.Get(img.Address(id))
      if err != nil {
         panic(err)
      }
      defer res.Body.Close()
      i, err := jpeg.Decode(res.Body)
      if err != nil {
         panic(err)
      }
      b, err := goimagehash.AverageHash(img.SubImage(i))
      if err != nil {
         panic(err)
      }
      // diff
      d, err := a.Distance(b)
      if err != nil {
         panic(err)
      }
      fmt.Println(d, id)
      time.Sleep(100 * time.Millisecond)
   }
}
