package decode

import (
   "os"
   "testing"
)

func TestDecode(t *testing.T) {
   f, err := os.Open("index.html")
   if err != nil {
      t.Fatal(err)
   }
   defer f.Close()
   d := NewDecoder(f)
   d.NextTag("title")
   d.Data()
}
