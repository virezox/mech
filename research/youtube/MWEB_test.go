package youtube

import (
   "fmt"
   "testing"
)

func TestMweb(t *testing.T) {
   const name = "MWEB"
   version, err := newVersion("https://m.youtube.com", "iPad")
   if err != nil {
      t.Fatal(err)
   }
   res, err := post(name, version)
   if err != nil {
      t.Fatal(err)
   }
   defer res.Body.Close()
   fmt.Println(res.Status, name, version)
}

func TestMwebTier(t *testing.T) {
   const (
      name = "MWEB_TIER_2"
      version = "9.20220325"
   )
   res, err := post(name, version)
   if err != nil {
      t.Fatal(err)
   }
   defer res.Body.Close()
   fmt.Println(res.Status, name, version)
}
