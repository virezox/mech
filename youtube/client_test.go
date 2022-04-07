package youtube

import (
   "fmt"
   "testing"
)

/*
bravo
HtVdAasjOgU
SZJvDhaSDnc
Tq92D6wQ1mg
WaOKSUlf4TM
i1Ko8UG-Tdo
nGC3D_FkCmg
yYr8q0y5Jfg
*/
var charlie = []string{
   "Cr381pDsSsA",
   "HsUATh_Nc2U",
}

const alfa = "zv9NimPx3Es"

func TestPlayer(t *testing.T) {
   play, err := Android.Player(alfa)
   if err != nil {
      t.Fatal(err)
   }
   if play.PlayabilityStatus.Status != "OK" {
      t.Fatal(play)
   }
}

func TestSearch(t *testing.T) {
   search, err := Mweb.Search("oneohtrix point never along")
   if err != nil {
      t.Fatal(err)
   }
   for _, item := range search.Items() {
      fmt.Println(item.CompactVideoRenderer)
   }
}
