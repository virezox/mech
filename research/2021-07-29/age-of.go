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
   "https://ia803106.us.archive.org/33/items" +
   "/mbid-46f66a2f-b782-47b4-9095-2ef933e1e7c7" +
   "/mbid-46f66a2f-b782-47b4-9095-2ef933e1e7c7-20015326395.jpg"

var ids = []string{
   "10kiSF1lBdU",
   "6iKPkxfljBY",
   "6jXbnydhNjU",
   "7vMHlm0lOks",
   "9rhadTURsrw",
   "E0RAmNU5Es8",
   "EGrv5FND4GY",
   "H-AWxL4DMTY",
   "HGcD8Xv0Ffw",
   "R7XcAaVumgc",
   "Wk_AOIwGeOs",
   "XbUOX4lr9Bw",
   "buYKf8iaUqg",
   "hFU8mVOfeEA",
   "iTjn1d-d2II",
   "nGj5N9Ll9pI",
   "qMQJF-7Y2h0",
   "rq74PDvCtLQ",
   "tgKAeggRHw8",
   "yergWdn968o",
}

var hqDef = youtube.Image{480, 360, 270, "hqdefault", youtube.JPG}

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
   return goimagehash.DifferenceHash(i)
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
