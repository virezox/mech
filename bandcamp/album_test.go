package bandcamp

import (
   "fmt"
   "testing"
)

func TestAlbum(t *testing.T) {
   a, err := AlbumGet("79940049")
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", a)
}
