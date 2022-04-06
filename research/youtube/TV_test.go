package youtube

import (
   "fmt"
   "testing"
)

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

func TestTvApple(t *testing.T) {
   const (
      name = "TVAPPLE"
      version = "1.0"
   )
   res, err := post(name, version)
   if err != nil {
      t.Fatal(err)
   }
   defer res.Body.Close()
   fmt.Println(res.Status, name, version)
}

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

func TestTvKids(t *testing.T) {
   const (
      name = "TVHTML5_KIDS"
      version = "3.20220325"
   )
   res, err := post(name, version)
   if err != nil {
      t.Fatal(err)
   }
   defer res.Body.Close()
   fmt.Println(res.Status, name, version)
}

func TestTvForKids(t *testing.T) {
   const (
      name = "TVHTML5_FOR_KIDS"
      version = "7.20220325"
   )
   res, err := post(name, version)
   if err != nil {
      t.Fatal(err)
   }
   defer res.Body.Close()
   fmt.Println(res.Status, name, version)
}

func TestTvSimply(t *testing.T) {
   const (
      name = "TVHTML5_SIMPLY"
      version = "1.0"
   )
   res, err := post(name, version)
   if err != nil {
      t.Fatal(err)
   }
   defer res.Body.Close()
   fmt.Println(res.Status, name, version)
}

func TestTvSimplyEmbed(t *testing.T) {
   const (
      name = "TVHTML5_SIMPLY_EMBEDDED_PLAYER"
      version = "2.0"
   )
   res, err := post(name, version)
   if err != nil {
      t.Fatal(err)
   }
   defer res.Body.Close()
   fmt.Println(res.Status, name, version)
}

func TestTvVr(t *testing.T) {
   const (
      name = "TVHTML5_VR"
      version = "0.1"
   )
   res, err := post(name, version)
   if err != nil {
      t.Fatal(err)
   }
   defer res.Body.Close()
   fmt.Println(res.Status, name, version)
}

func TestTvYongle(t *testing.T) {
   const (
      name = "TVHTML5_YONGLE"
      version = "0.1"
   )
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
