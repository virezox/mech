package main

import (
   "fmt"
   "math"
)

func variance(xs []float64) float64 {
   mean, M2 := 0.0, 0.0
   for n, x := range xs {
      delta := x - mean
      mean += delta / float64(n+1)
      M2 += delta * (x - mean)
   }
   return M2 / float64(len(xs)-1)
}

func stdDev(xs []float64) float64 {
   return math.Sqrt(variance(xs))
}

func main() {
   names := []string{"description", "Adam", "Bob", "Chris"}
   heights := []float64{70, 60, 65, 70}
   weights := []float64{170, 160, 180, 200}
   for i := range names {
      h := (heights[i]-heights[0]) / stdDev(heights)
      w := (weights[i]-weights[0]) / stdDev(weights)
      fmt.Printf("%v %.3f\n", names[i], h*h+w*w)
   }
}
