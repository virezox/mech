package paramount

import (
   "encoding/hex"
   "testing"
)

const (
   contentID = "eyT_RYkqNuH_6ZYrepLtxkiPO1HA7dIU"
   bearer = "eyJhbGciOiJIUzI1NiIsImtpZCI6IjNkNjg4NGJmLWViMDktNDA1Zi1hOWZjLWU0NGE1NmY3NjZiNiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJhbm9ueW1vdXNfVVMiLCJlbnQiOlt7ImJpZCI6IkFsbEFjY2Vzc01haW4iLCJlcGlkIjo3fV0sImlhdCI6MTY1MjgxODg4NCwiZXhwIjoxNjUyODI2MDg0LCJpc3MiOiJjYnMiLCJhaWQiOiJjYnNpIiwiaXNlIjp0cnVlLCJqdGkiOiI0M2EzMzU3YS03YjQ4LTRlMDUtOTVhMS0xNzdlNWY4Yzk4ODAifQ.6ghZKJXxG-nozchWLfGGKQc0G2xe-hoMkmdLbQVk8i0"
)

func TestParamount(t *testing.T) {
   keys, err := KeyContainers(contentID, bearer)
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
