package crackle

import (
   "fmt"
   "testing"
)

// crackle.com/watch/2992/2499348
const id = 2499348

func TestMedia(t *testing.T) {
   media, err := NewMedia(id)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", media)
}
