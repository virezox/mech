package youtube

import (
   "testing"
)

func TestAndroid(t *testing.T) {
   const name = "ANDROID"
   version, err := appVersion("com.google.android.youtube", phone)
   if err != nil {
      t.Fatal(err)
   }
   if version != names[name] {
      t.Fatal(name, version)
   }
   if err := post(name, version); err != nil {
      t.Fatal(err)
   }
}

func TestAndroidCreator(t *testing.T) {
   const name = "ANDROID_CREATOR"
   version, err := appVersion("com.google.android.apps.youtube.creator", phone)
   if err != nil {
      t.Fatal(err)
   }
   if version != names[name] {
      t.Fatal(name, version)
   }
   if err := post(name, version); err != nil {
      t.Fatal(err)
   }
}

func TestAndroidEmbeddedPlayer(t *testing.T) {
   const name = "ANDROID_EMBEDDED_PLAYER"
   version, err := appVersion("com.google.android.youtube", phone)
   if err != nil {
      t.Fatal(err)
   }
   if version != names[name] {
      t.Fatal(name, version)
   }
   if err := post(name, version); err != nil {
      t.Fatal(err)
   }
}

func TestAndroidKids(t *testing.T) {
   const name = "ANDROID_KIDS"
   version, err := appVersion("com.google.android.apps.youtube.kids", phone)
   if err != nil {
      t.Fatal(err)
   }
   if version != names[name] {
      t.Fatal(name, version)
   }
   if err := post(name, version); err != nil {
      t.Fatal(err)
   }
}

func TestAndroidLite(t *testing.T) {
   const name = "ANDROID_LITE"
   version, err := appVersion("com.google.android.apps.youtube.mango", phone)
   if err != nil {
      t.Fatal(err)
   }
   if version != names[name] {
      t.Fatal(name, version)
   }
   if err := post(name, version); err != nil {
      t.Fatal(err)
   }
}

func TestAndroidMusic(t *testing.T) {
   const name = "ANDROID_MUSIC"
   version, err := appVersion("com.google.android.apps.youtube.music", phone)
   if err != nil {
      t.Fatal(err)
   }
   if version != names[name] {
      t.Fatal(name, version)
   }
   if err := post(name, version); err != nil {
      t.Fatal(err)
   }
}

func TestAndroidTestsuite(t *testing.T) {
   const name = "ANDROID_TESTSUITE"
   err := post(name, names[name])
   if err != nil {
      t.Fatal(err)
   }
}

func TestAndroidTv(t *testing.T) {
   const name = "ANDROID_TV"
   version, err := appVersion("com.google.android.youtube.tv", tv)
   if err != nil {
      t.Fatal(err)
   }
   if version != names[name] {
      t.Fatal(name, version)
   }
   if err := post(name, version); err != nil {
      t.Fatal(err)
   }
}

func TestAndroidTvKids(t *testing.T) {
   const name = "ANDROID_TV_KIDS"
   version, err := appVersion("com.google.android.youtube.tvkids", tv)
   if err != nil {
      t.Fatal(err)
   }
   if version != names[name] {
      t.Fatal(name, version)
   }
   if err := post(name, version); err != nil {
      t.Fatal(err)
   }
}

func TestAndroidUnplugged(t *testing.T) {
   const name = "ANDROID_UNPLUGGED"
   version, err := appVersion("com.google.android.apps.youtube.unplugged", phone)
   if err != nil {
      t.Fatal(err)
   }
   if version != names[name] {
      t.Fatal(name, version)
   }
   if err := post(name, version); err != nil {
      t.Fatal(err)
   }
}

func TestAndroidVr(t *testing.T) {
   const name = "ANDROID_VR"
   version, err := appVersion("com.google.android.apps.youtube.vr", phone)
   if err != nil {
      t.Fatal(err)
   }
   if version != names[name] {
      t.Fatal(name, version)
   }
   if err := post(name, version); err != nil {
      t.Fatal(err)
   }
}
