package youtube_test

import (
   "fmt"
   "github.com/89z/mech/youtube"
   "testing"
)

func TestVideo(t *testing.T) {
   vid, err := youtube.NewVideo("UpNXI3_ctAc")
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", vid)
}
