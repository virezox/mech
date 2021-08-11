package mech

import (
   "fmt"
   "os"
   "testing"
)

func TestEncode(t *testing.T) {
   r, err := os.Open("index.html")
   if err != nil {
      panic(err)
   }
   defer r.Close()
   e := NewEncoder(os.Stdout)
   e.SetIndent(" ")
   e.Encode(r)
}

func TestScan(t *testing.T) {
   f, err := os.Open("nyt.html")
   if err != nil {
      panic(err)
   }
   defer f.Close()
   s := NewScanner(f)
   s.ScanAttr("type", "application/ld+json")
   fmt.Println(s.Attr("data-rh"))
   s.ScanText()
   fmt.Println(s.Data)
}
