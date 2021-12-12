package bandcamp

import (
   "fmt"
   "os"
   "testing"
)

func TestData(t *testing.T) {
   file, err := os.Open("index.html")
   if err != nil {
      t.Fatal(err)
   }
   defer file.Close()
   data, err := newDataTralbum(file)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", data)
}
