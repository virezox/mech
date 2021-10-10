package spotify

import (
   "fmt"
   "testing"
)

func TestConfig(t *testing.T) {
   cfg, err := newConfig()
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", cfg)
}
