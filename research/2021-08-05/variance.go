package main

import (
   "fmt"
   "gonum.org/v1/gonum/stat"
)

func main() {
   heights := []float64{70, 60, 65, 70}
   v := stat.Variance(heights, nil)
   fmt.Println(v)
}
