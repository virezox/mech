package main

import (
   "bufio"
   "strings"
   "testing"
)

func BenchmarkBufio(b *testing.B) {
   for n := 0; n < b.N; n++ {
      bufio.NewScanner(strings.NewReader(ww)).Scan()
   }
}

func BenchmarkFieldFunc(b *testing.B) {
   for n := 0; n < b.N; n++ {
      a := strings.FieldsFunc(ww, func(r rune) bool {
         return r == '.'
      })
      _ = a[0]
   }
}

func BenchmarkSplit(b *testing.B) {
   for n := 0; n < b.N; n++ {
      a := strings.Split(ww, ".")
      _ = a[0]
   }
}

func BenchmarkField(b *testing.B) {
   for n := 0; n < b.N; n++ {
      a := strings.Fields(ww)
      _ = a[0]
   }
}
