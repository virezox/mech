package youtube

import (
   "fmt"
   "testing"
   "time"
)

func TestPlayer(t *testing.T) {
   for _, client := range Clients {
      play, err := client.Player("sjNRpkQd68s")
      if err != nil {
         t.Fatal(err)
      }
      date := play.Microformat.PlayerMicroformatRenderer.PublishDate
      fmt.Printf("%+v %q\n", client, date)
      time.Sleep(time.Second)
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
