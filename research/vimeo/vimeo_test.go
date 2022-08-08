package vimeo

import (
   "fmt"
   "strconv"
   "testing"
   "time"
)

const (
   id = 18432
   origin = "http://embed.vhx.tv"
)

func tests() [2]string {
   var refs [2]string
   refs[0] = origin + "/subscriptions/" + strconv.Itoa(id) + "?vimeo=1"
   refs[1] = origin + "/videos/" + strconv.Itoa(id) + "?vimeo=1"
   return refs
}

func Test_Embed(t *testing.T) {
   for i := 0; i < 9; i++ {
      for _, test := range tests() {
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
}
