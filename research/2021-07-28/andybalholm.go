package main

import (
   "fmt"
   "github.com/89z/mech/youtube"
   "github.com/andybalholm/dhash"
   "image/jpeg"
   "net/http"
   "time"
)

func andybalholm(addr string, img *youtube.Image) (dhash.Hash, error) {
   r, err := http.Get(addr)
   if err != nil {
      return dhash.Hash{}, err
   }
   defer r.Body.Close()
   i, err := jpeg.Decode(r.Body)
   if err != nil {
      return dhash.Hash{}, err
   }
   if img != nil {
      i = img.SubImage(i)
   }
   return dhash.New(i), nil
}

func andybalholm_main(form youtube.Image) error {
   a, err := andybalholm(mb, nil)
   if err != nil {
      return err
   }
   for _, id := range ids {
      b, err := andybalholm(form.Address(id), &form)
      if err != nil {
         return err
      }
      fmt.Println(dhash.Distance(a, b), id)
      time.Sleep(100 * time.Millisecond)
   }
   return nil
}
