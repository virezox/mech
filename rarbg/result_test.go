package torrent_test

import (
   "github.com/89z/torrent"
   "testing"
)

func TestResult(t *testing.T) {
   _, err := torrent.NewResults("2020", "")
   if err != nil {
      t.Fatal(err)
   }
}
