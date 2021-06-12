package musicbrainz
import "testing"

func TestArtist(t *testing.T) {
   _, err := GroupFromArtist("678d88b2-87b0-403b-b63d-5da7465aecc3", 0)
   if err != nil {
      t.Fatal(err)
   }
}

func TestGroup(t *testing.T) {
   g, err := NewGroup("9b5006e5-b276-3a05-bcdd-8d986842320b")
   if err != nil {
      t.Fatal(err)
   }
   g.Sort()
   d := g.Releases[0].Date
   if d != "1973-03-28" {
      t.Fatal(d)
   }
}

func TestRelease(t *testing.T) {
   _, err := NewRelease("94f940c4-b1f0-4ccc-8886-424319eb39c8")
   if err != nil {
      t.Fatal(err)
   }
}
