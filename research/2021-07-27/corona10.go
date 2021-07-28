package main

import (
   "github.com/89z/mech/youtube"
   "github.com/corona10/goimagehash"
)

func corona10(addr string, img *youtube.Image) (*goimagehash.ImageHash, error) {
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

func corona10_main(img youtube.Image) error {
   a, err := corona10(mb, nil)
   if err != nil {
      return err
   }
   for _, id := range ids {
      b, err := corona10(img.Address(id), &img)
      if err != nil {
         return err
      }
      d, err := a.Distance(b)
      if err != nil {
         return err
      }
      fmt.Println(d, id)
      time.Sleep(100 * time.Millisecond)
   }
   return nil
}
