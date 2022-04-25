package channel4

import (
   "fmt"
   "testing"
)

const kid = "00000000-0000-0000-0000-000004246624"

func TestPSSH(t *testing.T) {
   pssh := createPSSH(kid)
   fmt.Println(pssh)
}
