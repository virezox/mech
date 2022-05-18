package paramount

import (
   "encoding/hex"
   "os"
   "testing"
)

const (
   contentID = "eyT_RYkqNuH_6ZYrepLtxkiPO1HA7dIU"
   bearer = "eyJhbGciOiJIUzI1NiIsImtpZCI6IjNkNjg4NGJmLWViMDktNDA1Zi1hOWZjLWU0NGE1NmY3NjZiNiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJhbm9ueW1vdXNfVVMiLCJlbnQiOlt7ImJpZCI6IkFsbEFjY2Vzc01haW4iLCJlcGlkIjo3fV0sImlhdCI6MTY1MjgyODQyMSwiZXhwIjoxNjUyODM1NjIxLCJpc3MiOiJjYnMiLCJhaWQiOiJjYnNpIiwiaXNlIjp0cnVlLCJqdGkiOiI1ODIwMGJmOC0yM2QyLTQwM2MtOThiYy1hNWM2YWRmOTU1NGEifQ.LAs9MvgyZDX3M2KxuhWuliVu_lb5TZA0hLGRGvrf4n8"
)

func TestParamount(t *testing.T) {
   media, err := NewMedia("ignore/stream.mpd")
   if err != nil {
      t.Fatal(err)
   }
   keys, err := media.Keys(contentID, bearer)
   if err != nil {
      t.Fatal(err)
   }
   var pass bool
   for _, key := range keys {
      if hex.EncodeToString(key.Key) == "44f12639c9c4a5a432338aca92e38920" {
         pass = true
      }
   }
   if !pass {
      t.Fatal(keys)
   }
}

func TestBearer(t *testing.T) {
   res, err := newBearer()
   if err != nil {
      t.Fatal(err)
   }
   defer res.Body.Close()
   os.Stdout.ReadFrom(res.Body)
}
