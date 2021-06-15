package tomato

import (
   "fmt"
   "testing"
)

func TestScore(t *testing.T) {
   s, err := NewScore("https://www.rottentomatoes.com/m/one_night_in_miami")
   if err != nil {
      t.Fatal(err)
   }
   a, err := s.NewAudience()
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", a)
   r, err := s.NewReview()
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", r)
}
