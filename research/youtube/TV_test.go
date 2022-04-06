package youtube

import (
   "fmt"
   "testing"
)

func TestTvhtml5(t *testing.T) {
   const name = "TVHTML5"
   version, err := newVersion(
      "https://www.youtube.com/tv",
      "Mozilla/5.0 (ChromiumStylePlatform) Cobalt/Version",
   )
   if err != nil {
      t.Fatal(err)
   }
   res, err := post(name, version)
   if err != nil {
      t.Fatal(err)
   }
   defer res.Body.Close()
   fmt.Println(res.Status, name, version)
}

func TestTvUnplug(t *testing.T) {
   const name = "TVHTML5_UNPLUGGED"
   version, err := appVersion("com.google.android.apps.youtube.unplugged", phone)
   if err != nil {
      t.Fatal(err)
   }
   res, err := post(name, version)
   if err != nil {
      t.Fatal(err)
   }
   defer res.Body.Close()
   fmt.Println(res.Status, name, version)
}

func TestTvAndroid(t *testing.T) {
   const (
      name = "TVANDROID"
      version = "1.0"
   )
   res, err := post(name, version)
   if err != nil {
      t.Fatal(err)
   }
   defer res.Body.Close()
   fmt.Println(res.Status, name, version)
}

func TestAndroidCast(t *testing.T) {
   const name = "TVHTML5_CAST"
   version, err := appVersion(
      "com.google.android.apps.youtube.music.pwa", tablet,
   )
   if err != nil {
      t.Fatal(err)
   }
   res, err := post(name, version)
   if err != nil {
      t.Fatal(err)
   }
   defer res.Body.Close()
   fmt.Println(res.Status, name, version)
}

func TestAndroidTvunplug(t *testing.T) {
   const name = "TV_UNPLUGGED_ANDROID"
   version, err := appVersion("com.google.android.youtube.tvunplugged", tv)
   if err != nil {
      t.Fatal(err)
   }
   res, err := post(name, version)
   if err != nil {
      t.Fatal(err)
   }
   defer res.Body.Close()
   fmt.Println(res.Status, name, version)
}

