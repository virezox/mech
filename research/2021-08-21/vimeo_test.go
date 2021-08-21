package vimeo

import (
   "testing"
)

func TestConfig(t *testing.T) {
   c, err := newConfig("66531465")
   if err != nil {
      t.Fatal(err)
   }
   if err := c.videos(); err != nil {
      t.Fatal(err)
   }
}
