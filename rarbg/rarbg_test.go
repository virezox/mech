package rarbg
import "testing"

func TestDefence(t *testing.T) {
   r, err := NewDefence()
   if err != nil {
      t.Fatal(err)
   }
   php, id, err := r.ThreatCaptcha()
   if err != nil {
      t.Fatal(err)
   }
   solve, err := Solve(php)
   if err != nil {
      t.Fatal(err)
   }
   if err := r.IamHuman(id, solve); err != nil {
      t.Fatal(err)
   }
}

func TestResult(t *testing.T) {
   _, err := NewResults("2020", "")
   if err != nil {
      t.Fatal(err)
   }
}
