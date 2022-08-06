package vimeo

import (
   "fmt"
   "testing"
   "time"
)

var tests = []string{
   "https://embed.vhx.tv/subscriptions/28599?vimeo=1",
   "https://embed.vhx.tv/videos/1264265?vimeo=1",
   "https://embed.vhx.tv/videos/28599?vimeo=1",
}

func Test_Embed(t *testing.T) {
   for _, test := range tests {
      emb, err := New_Embed(test)
      if err != nil {
         t.Fatal(err)
      }
      fmt.Printf("%+v\n", emb)
      time.Sleep(time.Second)
   }
}
