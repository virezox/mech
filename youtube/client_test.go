package youtube

import (
   "fmt"
   "os"
   "testing"
   "time"
)

func TestSearch(t *testing.T) {
   search, err := Mweb.Search("oneohtrix point never along")
   if err != nil {
      t.Fatal(err)
   }
   for _, item := range search.Items() {
      fmt.Println(item.CompactVideoRenderer)
   }
}

const android = "zv9NimPx3Es"

func TestAndroid(t *testing.T) {
   play, err := Android.Player(android)
   if err != nil {
      t.Fatal(err)
   }
   if play.PlayabilityStatus.Status != "OK" {
      t.Fatal(play)
   }
}

var androidEmbeds = []string{
   "HtVdAasjOgU",
   "WaOKSUlf4TM",
}

func TestAndroidEmbed(t *testing.T) {
   for _, embed := range androidEmbeds {
      play, err := AndroidEmbed.Player(embed)
      if err != nil {
         t.Fatal(err)
      }
      if play.PlayabilityStatus.Status != "OK" {
         t.Fatal(play)
      }
      time.Sleep(time.Second)
   }
}

var androidRacys = []string{
   "Cr381pDsSsA",
   "HsUATh_Nc2U", // signatureCipher
   "SZJvDhaSDnc", // url
   "Tq92D6wQ1mg", // url
   "dqRZDebPIGs", // signatureCipher
}

func TestAndroidRacy(t *testing.T) {
   cache, err := os.UserCacheDir()
   if err != nil {
      t.Fatal(err)
   }
   change, err := OpenExchange(cache, "mech/youtube.json")
   if err != nil {
      t.Fatal(err)
   }
   for _, racy := range androidRacys {
      play, err := AndroidRacy.Exchange(racy, change)
      if err != nil {
         t.Fatal(err)
      }
      if play.PlayabilityStatus.Status != "OK" {
         t.Fatal(play)
      }
      time.Sleep(time.Second)
   }
}

const androidContent = "nGC3D_FkCmg"

func TestAndroidContent(t *testing.T) {
   cache, err := os.UserCacheDir()
   if err != nil {
      t.Fatal(err)
   }
   change, err := OpenExchange(cache, "mech/youtube.json")
   if err != nil {
      t.Fatal(err)
   }
   play, err := AndroidContent.Exchange(androidContent, change)
   if err != nil {
      t.Fatal(err)
   }
   if play.PlayabilityStatus.Status != "OK" {
      t.Fatal(play)
   }
}
