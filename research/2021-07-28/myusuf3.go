package main

import (
   "fmt"
   "github.com/89z/mech/youtube"
   "github.com/myusuf3/imghash"
   "image/jpeg"
   "net/http"
   "time"
)

func myusuf3(addr string, img *youtube.Image) (uint64, error) {
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
   return imghash.Average(i), nil
}

func myusuf3_main(img youtube.Image) error {
   a, err := myusuf3(mb, nil)
   if err != nil {
      return err
   }
   for _, id := range ids {
      b, err := myusuf3(img.Address(id), &img)
      if err != nil {
         return err
      }
      fmt.Println(imghash.Distance(a, b), id)
      time.Sleep(100 * time.Millisecond)
   }
   return nil
}
