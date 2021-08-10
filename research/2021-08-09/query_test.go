package main

import (
   "bytes"
   "golang.org/x/net/html"
   "os"
   "testing"
)

// BenchmarkNode-12             526           2286841 ns/op
func BenchmarkNode(b *testing.B) {
   d, err := os.ReadFile("nyt.html")
   if err != nil {
      b.Fatal(err)
   }
   for n := 0; n < b.N; n++ {
      html.Parse(bytes.NewReader(d))
   }
}

// BenchmarkToken-12            582           2064419 ns/op
func BenchmarkToken(b *testing.B) {
   d, err := os.ReadFile("nyt.html")
   if err != nil {
      b.Fatal(err)
   }
   for n := 0; n < b.N; n++ {
      token(bytes.NewReader(d))
   }
}
