package instagram

import (
   "fmt"
   "os"
   "testing"
)

func TestRead(t *testing.T) {
   f, err := os.Open("instagram.json")
   if err != nil {
      t.Fatal(err)
   }
   defer f.Close()
   var auth Login
   if err := auth.Decode(f); err != nil {
      t.Fatal(err)
   }
   c, err := GraphQL("CT-cnxGhvvO", &auth)
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
   if err := l.Encode(os.Stdout); err != nil {
      t.Fatal(err)
   }
}
