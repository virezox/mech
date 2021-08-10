package mech

import (
   "fmt"
   "golang.org/x/net/html"
   "io"
   "os"
   "strings"
   "testing"
)

func TestEncode() {
   r, err := os.Open("in.html")
   if err != nil {
      panic(err)
   }
   defer r.Close()
   e := mech.NewEncoder(os.Stdout)
   e.SetIndent(" ")
   e.Encode(r)
}

func TestScan(t *testing.T) {
   f, err := os.Open("nyt.html")
   if err != nil {
      panic(err)
   }
   defer f.Close()
   s := mech.NewScanner(f)
   s.ScanAttr("type", "application/ld+json")
   fmt.Println(s.Attr("data-rh"))
   s.ScanText()
   fmt.Println(s.Data)
}
