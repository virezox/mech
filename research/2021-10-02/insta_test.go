package insta

import (
   "os"
   "testing"
)

func TestInsta(t *testing.T) {  
   pass, ok := os.LookupEnv("PASS")
   if ! ok {
      t.Fatal("PASS")
   }
   err := login("srpen6", pass)
   if err != nil {
      t.Fatal(err)
   }
}
