package youtube

import (
   "fmt"
   "testing"
)

func TestPlayer(t *testing.T) {
   play, err := Mweb.Player("sjNRpkQd68s")
   if err != nil {
      t.Fatal(err)
   }
   if play.Microformat.PlayerMicroformatRenderer.PublishDate == "" {
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
