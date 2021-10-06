package bandcamp

import (
   "fmt"
   "testing"
)

const (
   albumID = 79940049
   trackID = 2809477874
)

func TestAlbum(t *testing.T) {
   Verbose(true)
   alb := new(Album)
   if err := alb.Post(albumID); err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", alb)
}

func TestTrack(t *testing.T) {
   Verbose(true)
   tra := new(Track)
   if err := tra.Post(trackID); err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", tra)
}
