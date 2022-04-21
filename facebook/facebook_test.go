package facebook

import (
   "fmt"
   "testing"
)

// const auth = 444624393796648

func TestRegular(t *testing.T) {
   login, err := NewLogin()
   if err != nil {
      t.Fatal(err)
   }
   reg, err := login.Regular(email, password)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", reg)
}

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
}
