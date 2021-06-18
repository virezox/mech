package youtube_test

import (
   "github.com/89z/mech/youtube"
   "testing"
)

func TestVideo(t *testing.T) {
   _, err := youtube.NewVideo("UpNXI3_ctAc")
   if err != nil {
      t.Fatal(err)
   }
}
