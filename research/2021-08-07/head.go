package main

import (
   "fmt"
   "gonum.org/v1/gonum/stat"
   "math"
)

func distance(x, y int, v []float64, n [][]float64) float64 {
   var f float64
   for i, r := range n {
      f += math.Pow(r[x]-r[y], 2) / v[i]
   }
   return f
}

func main() {
   head := []string{"description", "Adam", "Bob", "Chris"}
   body := [][]float64{
      { 70,  60,  65,  70},
      {170, 160, 180, 200},
   }
   var variances []float64
   for _, r := range body {
      variances = append(variances, stat.Variance(r, nil))
   }
   fmt.Printf("%.3f\n", variances) // [22.917 291.667]
   for i := 1; i < len(head); i++ {
      d := distance(0, i, variances, body)
      fmt.Printf("%v %.3f\n", head[i], d)
   }
}
