package youtube

import (
   "fmt"
   "os"
   "testing"
   gp "github.com/89z/googleplay"
)

func TestAndroid(t *testing.T) {
   const name = "ANDROID"
   version, err := appVersion("com.google.android.youtube")
   if err != nil {
      t.Fatal(err)
   }
   res, err := post(name, version)
   if err != nil {
      t.Fatal(err)
   }
   defer res.Body.Close()
   fmt.Println(res.Status, name,  version)
}

func TestAndroidCreator(t *testing.T) {
   const name = "ANDROID_CREATOR"
   version, err := appVersion("com.google.android.apps.youtube.creator")
   if err != nil {
      t.Fatal(err)
   }
   res, err := post(name, version)
   if err != nil {
      t.Fatal(err)
   }
   defer res.Body.Close()
   fmt.Println(res.Status, name,  version)
}

func ANDROID_EMBEDDED_PLAYER() {
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
   det, err := head.Details("com.google.android.youtube")
   if err != nil {
      panic(err)
   }
   fmt.Println(det)
}

func ANDROID_KIDS() {
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
   det, err := head.Details("com.google.android.apps.youtube.kids")
   if err != nil {
      panic(err)
   }
   fmt.Println(det)
}

func ANDROID_MUSIC() {
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
