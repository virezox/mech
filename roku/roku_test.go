package roku

import (
   "fmt"
   "testing"
)

// therokuchannel.roku.com/watch/105c41ea75775968b670fbb26978ed76
const id = "105c41ea75775968b670fbb26978ed76"

func TestRoku(t *testing.T) {
   con, err := NewContent(id)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%a\n", con)
   fmt.Printf("%+v\n", con.Video())
}
