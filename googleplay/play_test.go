package googleplay

import (
   "net/url"
   "os"
   "testing"
)

func TestRead(t *testing.T) {
   txt, err := os.ReadFile("ac2dm.txt")
   if err != nil {
      t.Fatal(err)
   }
   q, err := url.ParseQuery(string(txt))
   if err != nil {
      t.Fatal(err)
   }
   a := ac2dm{q}
   o, err := a.oauth2()
   if err != nil {
      t.Fatal(err)
   }
   b, err := o.details("38B5418D8683ADBB")
   if err != nil {
      t.Fatal(err)
   }
   os.Stdout.Write(b)
}

func TestWrite(t *testing.T) {
   a, err := newAc2dm("oauth2_4/0AX4XfWg_wCzC2clsA3NJxGiVetvJcSA7kxw3k6ucbt-1j0Zex0WrVkzWFx2CM858fvhlEQ")
   if err != nil {
      t.Fatal(err)
   }
   f, err := os.Create("ac2dm.txt")
   if err != nil {
      t.Fatal(err)
   }
   defer f.Close()
   f.WriteString(a.Encode())
}
