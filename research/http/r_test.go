package http

import (
   "testing"
)

func TestTimeout(t *testing.T) {
   err := four("http://example.com")
   if err != nil {
      t.Fatal(err)
   }
}
