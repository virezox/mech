package youtube_test

import (
   "fmt"
   "github.com/89z/mech/youtube"
   "testing"
)

/*
http://ia800709.us.archive.org/34/items/
mbid-10cc746f-786c-4307-b8de-92a687489cb4/
mbid-10cc746f-786c-4307-b8de-92a687489cb4-4958564206.jpg
*/

func TestSearch(t *testing.T) {
   s, err := youtube.NewSearch("nelly furtado say it right")
   if err != nil {
      t.Fatal(err)
   }
   for _, item := range s.Items() {
      fmt.Printf("%+v\n", item)
   }
}
