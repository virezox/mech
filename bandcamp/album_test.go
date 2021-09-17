package bandcamp

import (
   "fmt"
   "testing"
)

const albumID = "79940049"

func TestAlbum(t *testing.T) {
   a := new(Album)
   err := a.Get(albumID)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", a)
}
