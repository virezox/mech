package instagram

import (
   "fmt"
   "os"
   "testing"
)

func TestRead(t *testing.T) {
   f, err := os.Open("ig.json")
   if err != nil {
      t.Fatal(err)
   }
   defer f.Close()
   auth, err := Decode(f)
   if err != nil {
      t.Fatal(err)
   }
   c, err := NewQuery("CT-cnxGhvvO").Sidecar(auth)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", c)
}

func TestWrite(t *testing.T) {
   pass, ok := os.LookupEnv("PASS")
   if ! ok {
      t.Fatal("PASS")
   }
   Verbose = true
   l, err := NewLogin("srpen6", pass)
   if err != nil {
      t.Fatal(err)
   }
   f, err := os.Create("ig.json")
   if err != nil {
      t.Fatal(err)
   }
   defer f.Close()
   if err := l.Encode(f); err != nil {
      t.Fatal(err)
   }
}
