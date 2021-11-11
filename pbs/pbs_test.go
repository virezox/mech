package pbs

import (
   "fmt"
   "github.com/89z/mech"
   "testing"
)

const addr =
   "https://www.pbs.org/wgbh/nova/video/nova-universe-revealed-milky-way/"

const slug = "nova-universe-revealed-milky-way-4io957"

func TestAsset(t *testing.T) {
   mech.Verbose(true)
   ass, err := NewAsset(slug)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", ass)
}

func TestEpisode(t *testing.T) {
   ep, err := NewEpisode(addr)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", ep)
}
