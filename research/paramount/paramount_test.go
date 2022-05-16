package paramount

import (
   "fmt"
   "testing"
)

const (
   contentID = "eyT_RYkqNuH_6ZYrepLtxkiPO1HA7dIU"
   bearer = "eyJhbGciOiJIUzI1NiIsImtpZCI6IjNkNjg4NGJmLWViMDktNDA1Zi1hOWZjLWU0NGE1NmY3NjZiNiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJhbm9ueW1vdXNfVVMiLCJlbnQiOlt7ImJpZCI6IkFsbEFjY2Vzc01haW4iLCJlcGlkIjo3fV0sImlhdCI6MTY1MjY3NTI5NSwiZXhwIjoxNjUyNjgyNDk1LCJpc3MiOiJjYnMiLCJhaWQiOiJjYnNpIiwiaXNlIjp0cnVlLCJqdGkiOiI1ZDAwMzRjNy1mZGY1LTQ5MmUtOTQzNS02NzQ4NzU0ZjEyMDMifQ.8TJfoE-JTMSjL0Nq7nevN_QJR0GEaKmF5FXhJNM6ksc"
)

func TestParamount(t *testing.T) {
   keys, err := newKeys(contentID, bearer)
   if err != nil {
      t.Fatal(err)
   }
   for _, key := range keys {
      fmt.Printf("%x\n", key.Value)
   }
}
