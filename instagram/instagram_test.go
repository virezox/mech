package instagram

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
   Verbose = true
   log, err := NewLogin("srpen6", pass)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Println(log)
}
