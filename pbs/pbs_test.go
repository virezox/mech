package pbs

import (
   "fmt"
   "testing"
)

const tAsset = "nova-universe-revealed-milky-way-4io957"

const tEpisode =
   "https://www.pbs.org/wgbh/nova/video/nova-universe-revealed-milky-way/"

func TestAsset(t *testing.T) {
   asset, err := NewAsset(tAsset)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", asset)
}

func TestEpisode(t *testing.T) {
   ep, err := NewEpisode(tEpisode)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", ep)
}
