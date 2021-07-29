package main

import (
   "fmt"
   "github.com/89z/mech/youtube"
   "github.com/corona10/goimagehash"
   "image/jpeg"
   "net/http"
   "time"
)

const mb =
   "https://ia800309.us.archive.org/9/items" +
   "/mbid-a40cb6e9-c766-37c4-8677-7eb51393d5a1" +
   "/mbid-a40cb6e9-c766-37c4-8677-7eb51393d5a1-9261666555.jpg"

var hqDef = youtube.Image{480, 360, 270, "hqdefault", youtube.JPG}

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

func hash(addr string, img *youtube.Image) (*goimagehash.ImageHash, error) {
   r, err := http.Get(addr)
   if err != nil {
      return nil, err
   }
   defer r.Body.Close()
   i, err := jpeg.Decode(r.Body)
   if err != nil {
      return nil, err
   }
   if img != nil {
      i = img.SubImage(i)
   }
   return goimagehash.PerceptionHash(i)
}

func main() {
   a, err := hash(mb, nil)
   if err != nil {
      panic(err)
   }
   for _, id := range ids {
      b, err := hash(hqDef.Address(id), &hqDef)
      if err != nil {
         panic(err)
      }
      d, err := a.Distance(b)
      if err != nil {
         panic(err)
      }
      fmt.Println(d, id)
      time.Sleep(100 * time.Millisecond)
   }
}
