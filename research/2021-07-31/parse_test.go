package main
import "testing"

func BenchmarkLeftRight(b *testing.B) {
   for n := 0; n < b.N; n++ {
      for _, s := range times {
         leftRight(s)
      }
   }
}

func BenchmarkMinHour(b *testing.B) {
   for n := 0; n < b.N; n++ {
      for _, s := range times {
         minHour(s)
      }
   }
}
