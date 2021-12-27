package mech

import (
   "fmt"
   "testing"
)

const number = 1_234_567_890

func TestNotation(t *testing.T) {
   f := compact().format(number)
   fmt.Println(f)
   f = compactSize().format(number)
   fmt.Println(f)
   f = compactRate().format(number)
   fmt.Println(f)
}
