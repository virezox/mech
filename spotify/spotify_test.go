package spotify

import (
   "fmt"
   "os"
   "testing"
)

func TestAlbum(t *testing.T) {
   f, err := os.Open("spotify.json")
   if err != nil {
      t.Fatal(err)
   }
   defer f.Close()
   var c Config
   if err := c.Decode(f); err != nil {
      t.Fatal(err)
   }
   Verbose(true)
   a, err := c.Album("4Aumawi2PZuCxo10dQc3vn")
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", a)
}

func TestPlaylist(t *testing.T) {
   f, err := os.Open("spotify.json")
   if err != nil {
      t.Fatal(err)
   }
   defer f.Close()
   var c Config
   if err := c.Decode(f); err != nil {
      t.Fatal(err)
   }
   p, err := c.Playlist("6rZ28nCpmG5Wo1Ik64EoDm")
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", p)
}

func TestWrite(t *testing.T) {
   c, err := NewConfig()
   if err != nil {
      t.Fatal(err)
   }
   f, err := os.Create("spotify.json")
   if err != nil {
      t.Fatal(err)
   }
   defer f.Close()
   if err := c.Encode(f); err != nil {
      t.Fatal(err)
   }
}
