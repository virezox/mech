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
   if err := post(name, version); err != nil {
      t.Fatal(err)
   }
}

func TestMwebTier(t *testing.T) {
   const name = "MWEB_TIER_2"
   err := post(name, names[name])
   if err != nil {
      t.Fatal(err)
   }
}
