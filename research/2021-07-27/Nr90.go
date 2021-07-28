package main

import (
   "fmt"
   "github.com/89z/mech/youtube"
   "github.com/Nr90/imgsim"
   "image/jpeg"
   "net/http"
   "time"
)

func Nr90_hash(addr string, img *youtube.Image) (imgsim.Hash, error) {
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
   return imgsim.DifferenceHash(i), nil
}

func Nr90_main(form youtube.Image) error {
   a, err := Nr90_hash(mb, nil)
   if err != nil {
      return err
   }
   for _, id := range ids {
      b, err := Nr90_hash(form.Address(id), &form)
      if err != nil {
         return err
      }
      fmt.Println(imgsim.Distance(a, b), id, form.Base)
      time.Sleep(100 * time.Millisecond)
   }
   return nil
}
