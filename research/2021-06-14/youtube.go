package main

import (
   "fmt"
   "strconv"
)

func main() {
   {
      s := `"\x7b\x22responseContext\x22:\x7b\x22serviceTrackingPara"`
      t, e := strconv.Unquote(s)
      if e != nil {
         panic(e)
      }
      println(t)
   }
   {
      s := `\x7b\x22responseContext\x22:\x7b\x22serviceTrackingPara`
      s = strconv.Quote(s)
      println(s)
   }
   {
      s := `\x7b\x22responseContext\x22:\x7b\x22serviceTrackingPara`
      fmt.Printf(`"%v"`, s)
   }
}
