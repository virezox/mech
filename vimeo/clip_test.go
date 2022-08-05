package vimeo

import (
   "fmt"
   "testing"
   "time"
)

var clips = []Clip{
   // vimeo.com/581039021/9603038895
   {581039021, "9603038895"},
}

func Test_Clip(t *testing.T) {
   web, err := New_JSON_Web()
   if err != nil {
      t.Fatal(err)
   }
   for _, clip := range clips {
      video, err := web.Video(&clip)
      if err != nil {
         t.Fatal(err)
      }
      fmt.Println(video)
      time.Sleep(time.Second)
   }
}
