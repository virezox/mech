package main

import (
   "fmt"
   "github.com/89z/youtube"
)

func main() {
   for _, each := range []string{"ipOogrq1m24", "BnEn7X3Pr7o"} {
      v, e := youtube.NewVideo(each)
      if e != nil {
         panic(e)
      }
      fmt.Printf("%+v\n\n", v)
   }
}
