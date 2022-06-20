package youtube

import (
   "fmt"
   "os"
   "testing"
   "time"
)

func Test_Search(t *testing.T) {
   search, err := Mweb.Search("oneohtrix point never along")
   if err != nil {
      t.Fatal(err)
   }
   for _, item := range search.Items() {
      fmt.Println(item.CompactVideoRenderer)
   }
}

const android = "zv9NimPx3Es"

func Test_Android(t *testing.T) {
   play, err := Android.Player(android)
   if err != nil {
      t.Fatal(err)
   }
   if play.PlayabilityStatus.Status != "OK" {
      t.Fatal(play)
   }
}

var android_embeds = []string{
   "HtVdAasjOgU",
   "WaOKSUlf4TM",
}

func Test_Android_Embed(t *testing.T) {
   for _, embed := range android_embeds {
      play, err := Android_Embed.Player(embed)
      if err != nil {
         t.Fatal(err)
      }
      if play.PlayabilityStatus.Status != "OK" {
         t.Fatal(play)
      }
      time.Sleep(time.Second)
   }
}

var android_racys = []string{
   "Cr381pDsSsA",
   "HsUATh_Nc2U", // signatureCipher
   "SZJvDhaSDnc", // url
   "Tq92D6wQ1mg", // url
   "dqRZDebPIGs", // signatureCipher
}

func Test_Android_Racy(t *testing.T) {
   home, err := os.UserHomeDir()
   if err != nil {
      t.Fatal(err)
   }
   change, err := Open_Exchange(home + "/mech/youtube.json")
   if err != nil {
      t.Fatal(err)
   }
   for _, racy := range android_racys {
      play, err := Android_Racy.Exchange(racy, change)
      if err != nil {
         t.Fatal(err)
      }
      if play.PlayabilityStatus.Status != "OK" {
         t.Fatal(play)
      }
      time.Sleep(time.Second)
   }
}

const android_content = "nGC3D_FkCmg"

func Test_Android_Content(t *testing.T) {
   home, err := os.UserHomeDir()
   if err != nil {
      t.Fatal(err)
   }
   change, err := Open_Exchange(home + "/mech/youtube.json")
   if err != nil {
      t.Fatal(err)
   }
   play, err := Android_Content.Exchange(android_content, change)
   if err != nil {
      t.Fatal(err)
   }
   if play.PlayabilityStatus.Status != "OK" {
      t.Fatal(play)
   }
}
