package cbc

import (
   "fmt"
   "os"
   "testing"
)

const downton = "downton-abbey/s01e05"

func TestMedia(t *testing.T) {
   home, err := os.UserHomeDir()
   if err != nil {
      t.Fatal(err)
   }
   profile, err := OpenProfile(home, "mech/cbc.json")
   if err != nil {
      t.Fatal(err)
   }
   asset, err := NewAsset(downton)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%a\n", asset)
   media, err := profile.Media(asset)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", media)
}
