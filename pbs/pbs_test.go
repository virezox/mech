package pbs

import (
   "fmt"
   "testing"
)

const tAsset = "nova-universe-revealed-milky-way-4io957"

const tSlug =
   "https://www.pbs.org/wgbh/nova/video/nova-universe-revealed-milky-way/"

func TestAsset(t *testing.T) {
   asset, err := NewAsset(tAsset)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", asset)
}

func TestSlug(t *testing.T) {
   slug, err := Slug(tSlug)
   if err != nil {
      t.Fatal(err)
   }
   if slug != tAsset {
      t.Fatal(slug)
   }
}
