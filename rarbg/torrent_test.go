package torrent_test

import (
   "github.com/89z/torrent"
   "testing"
)

func TestTorrent(t *testing.T) {
   r, err := torrent.NewDefence()
   if err != nil {
      t.Fatal(err)
   }
   php, id, err := r.ThreatCaptcha()
   if err != nil {
      t.Fatal(err)
   }
   solve, err := torrent.Solve(php)
   if err != nil {
      t.Fatal(err)
   }
   if err := r.IamHuman(id, solve); err != nil {
      t.Fatal(err)
   }
}
