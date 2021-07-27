package main

import (
   "github.com/89z/mech/youtube"
   "github.com/Nr90/imgsim"
   "image/jpeg"
   "net/http"
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
   return imgsim.DifferenceHash(i), nil
}
