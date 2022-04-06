package youtube

import (
   "testing"
)

func TestMweb(t *testing.T) {
   const name = "MWEB"
   version, err := newVersion("https://m.youtube.com", "iPad")
   if err != nil {
      t.Fatal(err)
   }
   if version != names[name] {
      t.Fatal(name, version)
   }
}
