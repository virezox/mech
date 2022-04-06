package youtube

import (
   "testing"
)

func TestTvAndroid(t *testing.T) {
   const name = "TVANDROID"
   err := post(name, names[name])
   if err != nil {
      t.Fatal(err)
   }
}

func TestTvApple(t *testing.T) {
   const name = "TVAPPLE"
   err := post(name, names[name])
   if err != nil {
      t.Fatal(err)
   }
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
   if version != names[name] {
      t.Fatal(name, version)
   }
   if err := post(name, version); err != nil {
      t.Fatal(err)
   }
}

func TestAndroidCast(t *testing.T) {
   const name = "TVHTML5_CAST"
   version, err := appVersion(
      "com.google.android.apps.youtube.music.pwa", tablet,
   )
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

func TestTvKids(t *testing.T) {
   const name = "TVHTML5_KIDS"
   err := post(name, names[name])
   if err != nil {
      t.Fatal(err)
   }
}

func TestTvForKids(t *testing.T) {
   const name = "TVHTML5_FOR_KIDS"
   err := post(name, names[name])
   if err != nil {
      t.Fatal(err)
   }
}

func TestTvSimply(t *testing.T) {
   const name = "TVHTML5_SIMPLY"
   err := post(name, names[name])
   if err != nil {
      t.Fatal(err)
   }
}

func TestTvSimplyEmbed(t *testing.T) {
   const name = "TVHTML5_SIMPLY_EMBEDDED_PLAYER"
   err := post(name, names[name])
   if err != nil {
      t.Fatal(err)
   }
}

func TestTvVr(t *testing.T) {
   const name = "TVHTML5_VR"
   err := post(name, names[name])
   if err != nil {
      t.Fatal(err)
   }
}

func TestTvYongle(t *testing.T) {
   const name = "TVHTML5_YONGLE"
   err := post(name, names[name])
   if err != nil {
      t.Fatal(err)
   }
}

func TestTvUnplug(t *testing.T) {
   const name = "TVHTML5_UNPLUGGED"
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

func TestAndroidTvunplug(t *testing.T) {
   const name = "TV_UNPLUGGED_ANDROID"
   version, err := appVersion("com.google.android.youtube.tvunplugged", tv)
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
