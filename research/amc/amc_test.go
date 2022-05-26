package amc

import (
   "os"
   "testing"
)

func TestAMC(t *testing.T) {
   res, err := unauth()
   if err != nil {
      t.Fatal(err)
   }
   defer res.Body.Close()
   os.Stdout.ReadFrom(res.Body)
}
