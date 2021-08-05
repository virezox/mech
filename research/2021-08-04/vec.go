package main

import (
   "fmt"
   "gonum.org/v1/gonum/spatial/r2"
)

func main() {
   description := r2.Vec{70, 170}
   suspects := map[string]r2.Vec{
      "Adam": {60, 160},
      "Bob": {65, 180},
      "Chris": {70, 200},
   }
   for name, suspect := range suspects {
      // sub
      suspect = r2.Sub(description, suspect)
      // square
      f := r2.Norm2(suspect)
      // print
      fmt.Println(name, f)
   }
}
