package facebook

import (
   "fmt"
   "io"
   "os"
   "testing"
)

const videoID = 309868367063220

func TestVideo(t *testing.T) {
   res, err := video(videoID)
   if err != nil {
      t.Fatal(err)
   }
   defer res.Body.Close()
   buf, err := io.ReadAll(res.Body)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Println(len(buf))
   os.WriteFile("ignore.html", buf, os.ModePerm)
}
