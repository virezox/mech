package main

import (
   "fmt"
   "time"
)

func main() {
   for range [9]struct{}{} {
      d, err := newData("radiohead")
      if err != nil {
         panic(err)
      }
      for _, item := range d.items() {
         fmt.Printf("%+v\n", item)
      }
      time.Sleep(time.Second)
   }
}
