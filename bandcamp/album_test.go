package bandcamp

import (
   "fmt"
   "testing"
   "time"
)

const albumID = "79940049"

func TestAlbum(t *testing.T) {
   Verbose(true)
   a := new(Album)
   if err := a.Get(albumID); err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", a)
   time.Sleep(100 * time.Millisecond)
   if err := a.Post(albumID); err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", a)
   time.Sleep(100 * time.Millisecond)
   if err := a.PostForm(albumID); err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", a)
}
