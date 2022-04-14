package facebook

import (
   "os"
   "testing"
)

const videoID = 309868367063220

func TestFacebook(t *testing.T) {
   res, err := video(videoID)
   if err != nil {
      t.Fatal(err)
   }
   defer res.Body.Close()
   file, err := os.Create("ignore.json")
   if err != nil {
      t.Fatal(err)
   }
   defer file.Close()
   if _, err := file.ReadFrom(res.Body); err != nil {
      t.Fatal(err)
   }
}
