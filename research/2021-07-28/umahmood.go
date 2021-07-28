package main

import (
   "fmt"
   "github.com/89z/mech/youtube"
   "github.com/umahmood/perceptive"
   "image/jpeg"
   "net/http"
   "time"
)

func umahmood(addr string, img *youtube.Image) (uint64, error) {
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
   return perceptive.Dhash(i)
}

func umahmood_main(img youtube.Image) error {
   a, err := umahmood(mb, nil)
   if err != nil {
      return err
   }
   for _, id := range ids {
      b, err := umahmood(img.Address(id), &img)
      if err != nil {
         return err
      }
      fmt.Println(perceptive.HammingDistance(a, b), id)
      time.Sleep(100 * time.Millisecond)
   }
   return nil
}
