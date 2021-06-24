package mech_test

import (
   "fmt"
   "github.com/89z/mech"
   "os"
   "testing"
)

func TestTransport(t *testing.T) {
   for _, gz := range []bool{false, true} {
      req, err := mech.NewRequest("GET", "https://github.com/manifest.json", nil)
      if err != nil {
         t.Fatal(err)
      }
      var t mech.Transport
      t.DisableCompression = gz
      res, err := t.RoundTrip(req)
      if err != nil {
         panic(err)
      }
      defer res.Body.Close()
      fmt.Println(res.ContentLength)
      os.Stdout.ReadFrom(res.Body)
      fmt.Println()
   }
}
