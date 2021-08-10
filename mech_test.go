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
   f, err := os.Open("nyt.html")
   if err != nil {
      panic(err)
   }
   defer f.Close()
   e := newEncoder(os.Stdout)
   e.setIndent(" ")
   if err := e.encode(f); err != nil {
      panic(err)
   }
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
