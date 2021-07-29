package main

import (
   "fmt"
   "github.com/89z/mech/youtube"
   "github.com/corona10/goimagehash"
   "image/jpeg"
   "net/http"
   "time"
)

const mb =
   "https://ia802904.us.archive.org/26/items" +
   "/mbid-8697d816-8cfc-4dcf-ace3-8a553796c742" +
   "/mbid-8697d816-8cfc-4dcf-ace3-8a553796c742-15610950895.jpg"

var ids = []string{
   "2hqqyncPrd0",
   "6ymbt3iTaII",
   "7-xPyDECnhc",
   "B3szYRzZqp4",
   "GVj4v-UCIQo",
   "HQrDHsO5CPY",
   "I4wPEgrIHyc",
   "JNkohQ9uckE",
   "OAtdmV7dMTM",
   "TjCkhONqzhA",
   "UNeE5Xs5IjA",
   "V-YZLKWShiY",
   "V9lKsr6gvFE",
   "XbUOX4lr9Bw",
   "bFje8as5iPU",
   "jCMi9_6vnxk",
   "m5Vwv42kxKo",
   "nGj5N9Ll9pI",
   "oDMGk6n-tns",
   "uHrWHXL065g",
   "ue2o_EokaIw",
   "yergWdn968o",
}

var hqDef = youtube.Image{480, 360, 270, "hqdefault", youtube.JPG}

func hash(addr string, img *youtube.Image) (*goimagehash.ImageHash, error) {
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
   return goimagehash.DifferenceHash(i)
}

func main() {
   a, err := hash(mb, nil)
   if err != nil {
      panic(err)
   }
   for _, id := range ids {
      b, err := hash(hqDef.Address(id), &hqDef)
      if err != nil {
         panic(err)
      }
      d, err := a.Distance(b)
      if err != nil {
         panic(err)
      }
      fmt.Println(d, id)
      time.Sleep(100 * time.Millisecond)
   }
}
