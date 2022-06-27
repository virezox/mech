package paramount

import (
   "fmt"
   "github.com/89z/mech/widevine"
   "os"
   "testing"
   "time"
)

var tests = map[test_type]string{
   {episode, dash_cenc}: "eyT_RYkqNuH_6ZYrepLtxkiPO1HA7dIU",
   {episode, stream_pack}: "622520382",
   {movie, stream_pack}: "wQH9yE_y_Dt4ekDYm3yelhhY2KXvOra_",
}

const (
   dash_cenc = iota
   episode
   movie
   stream_pack
)

type test_type struct {
   content_type int
   asset int
}

func Test_Preview(t *testing.T) {
   for _, test := range tests {
      preview, err := New_Media(test).Preview()
      if err != nil {
         t.Fatal(err)
      }
      fmt.Printf("%+v\n", preview)
      time.Sleep(time.Second)
   }
}
