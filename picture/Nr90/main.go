package main

import (
   "fmt"
   "github.com/89z/mech/youtube"
   "github.com/Nr90/imgsim"
   "image/jpeg"
   "net/http"
   "picture"
   "time"
)

func hash(addr string, img *youtube.Image) (imgsim.Hash, error) {
   r, err := http.Get(addr)
   if err != nil {
      return 0, err
   }
   defer r.Body.Close()
   i, err := jpeg.Decode(r.Body)
   if err != nil {
      return 0, err
   }
   if img != nil {
      i = img.SubImage(i)
   }
   return imgsim.AverageHash(i), nil
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
      fmt.Println(imgsim.Distance(a, b), id)
      time.Sleep(100 * time.Millisecond)
   }
}
