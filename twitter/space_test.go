package twitter

import (
   "fmt"
   "testing"
)

// twitter.com/i/spaces/1jMJgLVmMlbxL
const space_ID = "1jMJgLVmMlbxL"

func Test_Space(t *testing.T) {
   g, err := New_Guest()
   if err != nil {
      t.Fatal(err)
   }
   Client.Log_Level = 2
   s, err := g.Audio_Space(space_ID)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Println(s)
}
