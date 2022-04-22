package cbc

import (
   "os"
   "testing"
)

func TestCBC(t *testing.T) {
   res, err := login(email, password)
   if err != nil {
      t.Fatal(err)
   }
   defer res.Body.Close()
   os.Stdout.ReadFrom(res.Body)
}
