package youtube

import (
   "fmt"
   "testing"
)

func TestAndroidTV(t *testing.T) {
   const name = "ANDROID_UNPLUGGED"
   version, err := appVersion("com.google.android.apps.youtube.unplugged")
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

func TestAndroidMusic(t *testing.T) {
   const name = "ANDROID_MUSIC"
   version, err := appVersion("com.google.android.apps.youtube.music")
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
   fmt.Println(res.Status, name, version)
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
   fmt.Println(res.Status, name, version)
}

func TestAndroidEmbeddedPlayer(t *testing.T) {
   const name = "ANDROID_EMBEDDED_PLAYER"
   version, err := appVersion("com.google.android.youtube")
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

func TestAndroidKids(t *testing.T) {
   const name = "ANDROID_KIDS"
   version, err := appVersion("com.google.android.apps.youtube.kids")
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
