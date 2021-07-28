package main

import (
   "fmt"
   "github.com/89z/mech/youtube"
   "github.com/brett-lempereur/ish"
   "image/jpeg"
   "net/http"
   "time"
)

func brett(addr string, img *youtube.Image) ([]byte, error) {
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
   return ish.NewAverageHash(8, 8).Hash(i)
}

func brett_main(img youtube.Image) error {
   a, err := brett(mb, nil)
   if err != nil {
      return err
   }
   for _, id := range ids {
      b, err := brett(img.Address(id), &img)
      if err != nil {
         return err
      }
      fmt.Println(ish.NewAverageHash(8, 8).Distance(a, b), id)
      time.Sleep(100 * time.Millisecond)
   }
   return nil
}
