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
   a := Ac2dm{q}
   o, err := a.OAuth2()
   if err != nil {
      t.Fatal(err)
   }
   b, err := o.Details("38B5418D8683ADBB", "com.google.android.youtube")
   if err != nil {
      t.Fatal(err)
   }
   os.Stdout.Write(b)
}

func TestWrite(t *testing.T) {
   a, err := NewAc2dm("oauth2_4/0AX4XfWg_wCzC2clsA3NJxGiVetvJcSA7kxw3k6ucbt-1j0Zex0WrVkzWFx2CM858fvhlEQ")
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
