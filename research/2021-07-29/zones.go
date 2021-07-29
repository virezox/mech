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
   "https://ia801601.us.archive.org/8/items" +
   "/mbid-82a5fee2-69eb-4bd6-9dce-d4cc62bfad9a" +
   "/mbid-82a5fee2-69eb-4bd6-9dce-d4cc62bfad9a-2909521338.jpg"

var ids = []string{
   "10kiSF1lBdU",
   "3soucm9emO4",
   "7BAEZ59ZC64",
   "7T1_KxrX4vk",
   "9YE26ZSKzH0",
   "C4uBwEfftDg",
   "E0RAmNU5Es8",
   "HNd3gcrHd1w",
   "HcgD3Y79Uqc",
   "I4wPEgrIHyc",
   "Wk_AOIwGeOs",
   "dAedl0vK8Uo",
   "hW7PxGmZgTw",
   "m5Vwv42kxKo",
   "ue2o_EokaIw",
   "yergWdn968o",
   "yv_8zAB1OUM",
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
