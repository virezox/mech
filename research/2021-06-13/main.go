package main

import (
   "fmt"
   "time"
)

func main() {
   for range [9]struct{}{} {
      data, err := newYTInitialData("radiohead")
      if err != nil {
         panic(err)
      }
      for _, vid := range data.primaryContents().videoRenderers() {
         fmt.Printf("%+v\n", vid)
      }
      time.Sleep(time.Second)
   }
}
