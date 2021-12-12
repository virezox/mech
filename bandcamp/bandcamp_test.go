package bandcamp

import (
   "fmt"
   "testing"
   "time"
)

const artID = 3809045440

var tralbums = []tralbum{
   {1670971920, 'a', "https://schnaussandmunk.bandcamp.com/album/passage-2"},
   {2809477874, 't', "https://schnaussandmunk.bandcamp.com/track/amaris-2"},
}

type tralbum struct {
   id int
   typ byte
   addr string
}

func TestDataTralbum(t *testing.T) {
   for _, tralbum := range tralbums {
      data, err := NewDataTralbum(tralbum.addr)
      if err != nil {
         t.Fatal(err)
      }
      fmt.Printf("%+v\n", data)
      time.Sleep(time.Second)
   }
}

func TestImage(t *testing.T) {
   for _, img := range Images {
      addr := img.Format(artID)
      fmt.Println(addr)
   }
}

func TestTralbum(t *testing.T) {
   for _, ta := range tralbums {
      data, err := NewTralbum(ta.typ, ta.id)
      if err != nil {
         t.Fatal(err)
      }
      fmt.Printf("%+v\n", data)
      time.Sleep(time.Second)
   }
}
