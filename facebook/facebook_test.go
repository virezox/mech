package facebook

import (
   "fmt"
   "testing"
)

const anon = 309868367063220

func TestVideo(t *testing.T) {
   vid, err := NewVideo(anon)
   if err != nil {
      t.Fatal(err)
   }
   if vid.DateCreated == "" {
      t.Fatal(vid)
   }
   if vid.Playable_URL_Quality_HD == "" {
      t.Fatal(vid)
   }
   if vid.Preferred_Thumbnail.Image.URI == "" {
      t.Fatal(vid)
   }
   if vid.Text == "" {
      t.Fatal(vid)
   }
   if err := vid.Preferred_Thumbnail.Parse(); err != nil {
      t.Fatal(err)
   }
   fmt.Println(vid)
}
