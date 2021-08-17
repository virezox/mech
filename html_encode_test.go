package mech

import (
   "os"
   "testing"
)

func TestEncode(t *testing.T) {
   f, err := os.Open("index.html")
   if err != nil {
      t.Fatal(err)
   }
   defer f.Close()
   enc := NewEncoder(os.Stdout)
   enc.SetIndent(" ")
   if err := enc.Encode(f); err != nil {
      t.Fatal(err)
   }
}
