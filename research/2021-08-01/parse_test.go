package parse
import "testing"

func BenchmarkDateErr(b *testing.B) {
   for n := 0; n < b.N; n++ {
      for _, s := range dates {
         if _, err := dateErr(s); err != nil {
            b.Fatal(err)
         }
      }
   }
}

func BenchmarkDatePad(b *testing.B) {
   for n := 0; n < b.N; n++ {
      for _, s := range dates {
         if _, err := datePad(s); err != nil {
            b.Fatal(err)
         }
      }
   }
}

func BenchmarkTimeCrop(b *testing.B) {
   for n := 0; n < b.N; n++ {
      for _, s := range times {
         if _, err := timeCrop(s); err != nil {
            b.Fatal(err)
         }
      }
   }
}

func BenchmarkTimeErr(b *testing.B) {
   for n := 0; n < b.N; n++ {
      for _, s := range times {
         if _, err := timeErr(s); err != nil {
            b.Fatal(err)
         }
      }
   }
}

func BenchmarkTimePad(b *testing.B) {
   for n := 0; n < b.N; n++ {
      for _, s := range times {
         if _, err := timePad(s); err != nil {
            b.Fatal(err)
         }
      }
   }
}
