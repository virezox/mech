package vimeo

import (
   "fmt"
   "testing"
   "time"
)

var tests = []string{
   // 1080p
   // Villa Touma - clip #1
   "http://embed.vhx.tv/videos/17901?vimeo=1",
   // 1080p
   // Now That You're Married Trailer
   "http://embed.vhx.tv/subscriptions/17901?vimeo=1",
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
