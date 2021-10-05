package bandcamp

import (
   "fmt"
   "testing"
   "time"
)

const albumID = 79940049

func TestAlbum(t *testing.T) {
   Verbose(true)
   alb := new(Album)
   if err := alb.Get(albumID); err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", alb)
   time.Sleep(100 * time.Millisecond)
   if err := alb.Post(albumID); err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", alb)
   time.Sleep(100 * time.Millisecond)
   if err := alb.PostForm(albumID); err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", alb)
}
