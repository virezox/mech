package vimeo

import (
   "fmt"
   "testing"
   "time"
)

var embed_refs = []string{
   "https://embed.vhx.tv/subscriptions/18432?vimeo=1",
   "https://embed.vhx.tv/videos/18432?vimeo=1",
}

func Test_Embed(t *testing.T) {
   for i := 0; i < 9; i++ {
      for _, ref := range embed_refs {
         emb, err := New_Embed(ref)
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
}
