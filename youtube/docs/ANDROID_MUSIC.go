package main

import (
   "fmt"
   "os"
   gp "github.com/89z/googleplay"
)

func main() {
   cache, err := os.UserCacheDir()
   if err != nil {
      panic(err)
   }
   tok, err := gp.OpenToken(cache, "googleplay/token.json")
   if err != nil {
      panic(err)
   }
   dev, err := gp.OpenDevice(cache, "googleplay/device.json")
   if err != nil {
      panic(err)
   }
   head, err := tok.Header(dev)
   if err != nil {
      panic(err)
   }
   det, err := head.Details("com.google.android.apps.youtube.music")
   if err != nil {
      panic(err)
   }
   fmt.Println(det)
}
