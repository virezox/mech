package pbs

import (
   "fmt"
   "github.com/89z/mech"
   "testing"
)

const tAsset = "nova-universe-revealed-milky-way-4io957"

const tEpisode =
   "https://www.pbs.org/wgbh/nova/video/nova-universe-revealed-milky-way/"

func TestAsset(t *testing.T) {
   mech.Verbose(true)
   asset, err := NewAsset(tAsset)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", asset)
}

func TestEpisode(t *testing.T) {
   mech.Verbose(true)
   ep, err := NewEpisode(tEpisode)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", ep)
}
