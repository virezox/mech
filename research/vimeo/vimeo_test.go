package vimeo

import (
   "fmt"
   "testing"
   "time"
)

var tests = []string{
   "http://embed.vhx.tv/subscriptions/17901?vimeo=1",
   "http://embed.vhx.tv/videos/17901?vimeo=1",
}

func Test_Embed(t *testing.T) {
   for _, test := range tests {
      emb, err := New_Embed(test)
      if err != nil {
         t.Fatal(err)
      }
      con, err := emb.Config()
      if err != nil {
         t.Fatal(err)
      }
      fmt.Println(con)
      time.Sleep(time.Second)
   }
}
