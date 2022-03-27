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
   for _, client := range Clients {
      if client.Name == "MWEB" {
         s, err := client.Search("oneohtrix point never along")
         if err != nil {
            t.Fatal(err)
         }
         for _, i := range s.Items() {
            fmt.Println(i.CompactVideoRenderer)
         }
      }
   }
}
