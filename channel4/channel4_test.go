package channel4

import (
   "fmt"
   "testing"
)

// channel4.com/programmes/frasier/on-demand/18926-001
const frasier = "18926-001"

func TestStream(t *testing.T) {
   stream, err := NewStream(frasier)
   if err != nil {
      t.Fatal(err)
   }
   widevine := stream.Widevine()
   fmt.Printf("%+v\n", widevine)
}
