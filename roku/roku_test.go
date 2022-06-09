package roku

import (
   "fmt"
   "testing"
)

// therokuchannel.roku.com/watch/105c41ea75775968b670fbb26978ed76
const id = "105c41ea75775968b670fbb26978ed76"

func TestVideo(t *testing.T) {
   con, err := NewContent(id)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%a\n", con)
   video, err := con.HLS()
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", video)
}
