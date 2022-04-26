package channel4

import (
   "testing"
)

const (
   in = "00000000-0000-0000-0000-000004246624"
   out = "AAAAMnBzc2gAAAAA7e+LqXnWSs6jyCfc1R0h7QAAABISEAAAAAAAAAAAAAAAAAQkZiQ="
)

func TestPSSH(t *testing.T) {
   pssh, err := createPSSH(in)
   if err != nil {
      t.Fatal(err)
   }
   if pssh != out {
      t.Fatal(pssh)
   }
}
