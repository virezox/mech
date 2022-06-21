package crackle

import (
   "fmt"
   "testing"
)

// crackle.com/watch/2992/2499348
const id = 2499348

func Test_Media(t *testing.T) {
   media, err := New_Media(id)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", media)
}
