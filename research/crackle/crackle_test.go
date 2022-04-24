package crackle

import (
   "os"
   "testing"
)

// crackle.com/watch/2992/2499348
const id = 2499348

func TestMedia(t *testing.T) {
   res, err := media(id)
   if err != nil {
      t.Fatal(err)
   }
   defer res.Body.Close()
   os.Stdout.ReadFrom(res.Body)
}
