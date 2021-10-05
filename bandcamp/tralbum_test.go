package bandcamp

import (
   "fmt"
   "testing"
   "time"
)

const trackID = 2809477874

func TestTrack(t *testing.T) {
   Verbose(true)
   tra := new(Track)
   if err := tra.Get(trackID); err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", tra)
   time.Sleep(100 * time.Millisecond)
   if err := tra.Post(trackID); err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", tra)
   time.Sleep(100 * time.Millisecond)
   if err := tra.PostForm(trackID); err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", tra)
}

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
