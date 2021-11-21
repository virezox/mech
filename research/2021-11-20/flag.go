package main

import (
   "flag"
   "fmt"
)

func main() {
   var ss []string
   flag.Func("s", "string (repeated)", func(s string) error {
      ss = append(ss, s)
      return nil
   })
   flag.Parse()
   fmt.Printf("%q\n", ss)
}
