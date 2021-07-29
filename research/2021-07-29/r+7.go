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
   "https://ia600709.us.archive.org/34/items" +
   "/mbid-10cc746f-786c-4307-b8de-92a687489cb4" +
   "/mbid-10cc746f-786c-4307-b8de-92a687489cb4-4958564206.jpg"

var ids = []string{
   "1UztCDH2xuQ",
   "6iKPkxfljBY",
   "F1YdyaJeb1E",
   "GlhV-OKHecI",
   "MYr5MypHAhQ",
   "R7XcAaVumgc",
   "VKvn_YxuJQc",
   "WA8oNVFPppw",
   "Wk_AOIwGeOs",
   "XbUOX4lr9Bw",
   "eud9OOVM4to",
   "mjnAE5go9dI",
   "qMQJF-7Y2h0",
   "qmlJveN9IkI",
   "svTiG5vZ0_A",
   "uKna8o35UsU",
   "uhcnxH9zTEo",
   "unN7QvSWSTo",
   "w5azY0dH67U",
   "yGsCzZuK9GI",
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
   return goimagehash.AverageHash(i)
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
