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

package bandcamp

import (
   "fmt"
   "testing"
   "time"
)

const bandID = "2853020814"

func TestBand(t *testing.T) {
   Verbose(true)
   b := new(Band)
   if err := b.Get(bandID); err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", b)
   time.Sleep(100 * time.Millisecond)
   if err := b.Post(bandID); err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", b)
   time.Sleep(100 * time.Millisecond)
   if err := b.PostForm(bandID); err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", b)
   d := new(Discography)
   time.Sleep(100 * time.Millisecond)
   if err := d.Get(bandID); err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", d)
   time.Sleep(100 * time.Millisecond)
   if err := d.Post(bandID); err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", d)
   time.Sleep(100 * time.Millisecond)
   if err := d.PostForm(bandID); err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", d)
}
