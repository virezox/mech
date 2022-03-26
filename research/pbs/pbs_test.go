package pbs

import (
   "testing"
)

const good = "https://urs.pbs.org/redirect/20969856f40f4935b53b86555d573b1f/"

func TestPBS(t *testing.T) {
   asset, err := newAsset("nova-universe-revealed-milky-way-4io957")
   if err != nil {
      t.Fatal(err)
   }
   if asset.Resource.MP4_Videos[0].URL != good {
      t.Fatal(asset)
   }
}
