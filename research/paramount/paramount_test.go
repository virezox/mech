package paramount

import (
   "encoding/hex"
   "testing"
)

const (
   contentID = "eyT_RYkqNuH_6ZYrepLtxkiPO1HA7dIU"
   bearer = "eyJhbGciOiJIUzI1NiIsImtpZCI6IjNkNjg4NGJmLWViMDktNDA1Zi1hOWZjLWU0NGE1NmY3NjZiNiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJhbm9ueW1vdXNfVVMiLCJlbnQiOlt7ImJpZCI6IkFsbEFjY2Vzc01haW4iLCJlcGlkIjo3fV0sImlhdCI6MTY1MjgxMTA2MCwiZXhwIjoxNjUyODE4MjYwLCJpc3MiOiJjYnMiLCJhaWQiOiJjYnNpIiwiaXNlIjp0cnVlLCJqdGkiOiIyN2NjOGQ5YS00Yzc1LTQxNjEtYjFhYy0zNDljMGE3M2M0YTUifQ.fDyVZgmazielPWFjHvFkURdCYm40LIXFldVh_vJm05A"
)

func TestParamount(t *testing.T) {
   keys, err := newKeys(contentID, bearer)
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
