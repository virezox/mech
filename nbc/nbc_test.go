package nbc

import (
   "fmt"
   "testing"
   "time"
)

var guids = []int64{
   // nbc.com/botched/video/seeing-double/3049418
   3049418,
   // nbc.com/la-brea/video/pilot/9000194212
   9000194212,
   // nbc.com/saturday-night-live/video/april-2-jerrod-carmichael/9000199373
   9000199373,
   // nbc.com/pasion-de-gavilanes/video/la-valentia-de-norma/9000221348
   9000221348,
}

func Test_Video(t *testing.T) {
   for _, guid := range guids {
      page, err := New_Bonanza_Page(guid)
      if err != nil {
         t.Fatal(err)
      }
      vid, err := page.Video()
      if err != nil {
         t.Fatal(err)
      }
      fmt.Printf("%+v\n", vid)
      time.Sleep(time.Second)
   }
}
