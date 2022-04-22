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
   if vid.Date.DateCreated == "" {
      t.Fatal(vid)
   }
   if vid.Media.Playable_URL_Quality_HD == "" {
      t.Fatal(vid)
   }
   if vid.Media.Preferred_Thumbnail.Image.URI == "" {
      t.Fatal(vid)
   }
   if vid.Title.Text == "" {
      t.Fatal(vid)
   }
   fmt.Println(vid)
   if err := vid.Media.Preferred_Thumbnail.Parse(); err != nil {
      t.Fatal(err)
   }
   fmt.Println(vid)
}
