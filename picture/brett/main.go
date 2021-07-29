package main

import (
   "fmt"
   "github.com/89z/mech/youtube"
   "github.com/brett-lempereur/ish"
   "image/jpeg"
   "net/http"
   "picture"
   "time"
)

const (
   width = 8
   height = 8
)

func hash(addr string, img *youtube.Image) ([]byte, error) {
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
   return ish.NewAverageHash(width, height).Hash(i)
}

func main() {
   a, err := hash(picture.MB, nil)
   if err != nil {
      panic(err)
   }
   for _, id := range picture.Ids {
      b, err := hash(picture.HqDef.Address(id), &picture.HqDef)
      if err != nil {
         panic(err)
      }
      fmt.Println(ish.NewAverageHash(width, height).Distance(a, b), id)
      time.Sleep(100 * time.Millisecond)
   }
}
