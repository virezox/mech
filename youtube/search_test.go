package youtube_test

import (
   "fmt"
   "github.com/89z/mech/youtube"
   "image/jpeg"
   "net/http"
   "testing"
)

const mb =
   "http://ia800709.us.archive.org/34/items" +
   "/mbid-10cc746f-786c-4307-b8de-92a687489cb4" +
   "/mbid-10cc746f-786c-4307-b8de-92a687489cb4-4958564206.jpg"

func TestSearch(t *testing.T) {
   s, err := youtube.NewSearch("oneohtrix point never along")
   if err != nil {
      t.Fatal(err)
   }
   r, err := http.Get(mb)
   if err != nil {
      t.Fatal(err)
   }
   defer r.Body.Close()
   img, err := jpeg.Decode(r.Body)
   if err != nil {
      t.Fatal(err)
   }
   items := s.Items()
   items.Sort(img)
   for i, item := range items {
      fmt.Println(i, item)
   }
}
