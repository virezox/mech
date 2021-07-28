package main

import (
   "fmt"
   "github.com/89z/mech/youtube"
   "github.com/devedge/imagehash"
   "image/jpeg"
   "net/http"
   "time"
)

func devedge_ahash(addr string, img *youtube.Image) ([]byte, error) {
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
   return imagehash.Ahash(i, 8)
}

func devedge_main(form youtube.Image) error {
   a, err := devedge_ahash(mb, nil)
   if err != nil {
      return err
   }
   fmt.Println("Ahash", form.Base)
   for _, id := range ids {
      b, err := devedge_ahash(form.Address(id), &form)
      if err != nil {
         return err
      }
      fmt.Println(imagehash.GetDistance(a, b), id)
      time.Sleep(100 * time.Millisecond)
   }
   return nil
}
