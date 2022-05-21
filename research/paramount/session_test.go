package paramount

import (
   "fmt"
   "testing"
)

const contentID = "eyT_RYkqNuH_6ZYrepLtxkiPO1HA7dIU"

func TestSession(t *testing.T) {
   sess, err := NewSession(contentID)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", sess)
}
