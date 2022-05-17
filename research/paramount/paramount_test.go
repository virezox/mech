package paramount

import (
   "encoding/hex"
   "testing"
)

const (
   contentID = "eyT_RYkqNuH_6ZYrepLtxkiPO1HA7dIU"
   bearer = "eyJhbGciOiJIUzI1NiIsImtpZCI6IjNkNjg4NGJmLWViMDktNDA1Zi1hOWZjLWU0NGE1NmY3NjZiNiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJhbm9ueW1vdXNfVVMiLCJlbnQiOlt7ImJpZCI6IkFsbEFjY2Vzc01haW4iLCJlcGlkIjo3fV0sImlhdCI6MTY1Mjc1NzYwMiwiZXhwIjoxNjUyNzY0ODAyLCJpc3MiOiJjYnMiLCJhaWQiOiJjYnNpIiwiaXNlIjp0cnVlLCJqdGkiOiI5ODFiNWUzNy02N2JiLTQyMGQtOTAzNC05NWRhYTA0YzZmMTcifQ.JXVM05k3a9O0Jf98viyszX_8Szf2pqch1NpDBwC4yK4"
)

func TestParamount(t *testing.T) {
   keys, err := newKeys(contentID, bearer)
   if err != nil {
      t.Fatal(err)
   }
   var pass bool
   for _, key := range keys {
      if hex.EncodeToString(key.Value) == "44f12639c9c4a5a432338aca92e38920" {
         pass = true
      }
   }
   if !pass {
      t.Fatal(keys)
   }
}
