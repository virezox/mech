package pandora

import (
   "fmt"
   "os"
   "testing"
   "time"
)

const addr =
   "https://pandora.com/artist/the-black-dog/radio-scarecrow" +
   "/train-by-the-autobahn-part-1/TRddpp5JJ2hqnVV"

func TestID(t *testing.T) {
   id, err := PandoraID(addr)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Println(id)
}
