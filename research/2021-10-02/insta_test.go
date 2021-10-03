package insta

import (
   "fmt"
   "os"
   "testing"
)

func TestLogin(t *testing.T) {  
   pass, ok := os.LookupEnv("PASS")
   if ! ok {
      t.Fatal("PASS")
   }
   log, err := newLogin("srpen6", pass)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Println(log)
}
