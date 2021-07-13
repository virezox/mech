package youtube

import (
   "os"
   "path/filepath"
   "testing"
)

func TestDebug(t *testing.T) {
   i := newYouTubeI()
   i.VideoID = "RBHwO-2Amzs"
   res, err := i.post("/youtubei/v1/player")
   if err != nil {
      t.Fatal(err)
   }
   defer res.Body.Close()
   tmp := filepath.Join(os.TempDir(), "mech-youtube.json")
   f, err := os.Create(tmp)
   if err != nil {
      t.Fatal(err)
   }
   defer f.Close()
   f.ReadFrom(res.Body)
   println("file:" + tmp)
}
