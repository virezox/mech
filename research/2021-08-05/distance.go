package main

import (
   "fmt"
   "math"
)

func main() {
   names := []string{"description", "Adam", "Bob", "Chris"}
   heights := []float64{70, 60, 65, 70}
   weights := []float64{170, 160, 180, 200}
   vHeight := variance(heights)
   vWeight := variance(weights)
   for i := range names {
      m := mahalanobis{
         {heights[0], heights[i], vHeight},
         {weights[0], weights[i], vWeight},
      }
      fmt.Printf("%v %.3f\n", names[i], m.distance())
   }
}

func variance(xs []float64) float64 {
   var M2, mean float64
   for n, x := range xs {
      delta := x - mean
      mean += delta / float64(n + 1)
      M2 += delta * (x - mean)
   }
   return M2 / float64(len(xs) - 1)
}

type mahalanobis []struct {
   x, y, variance float64
}

func (m mahalanobis) distance() float64 {
   var sum float64
   for _, i := range m {
      sum += math.Pow(i.x-i.y, 2) / i.variance
   }
   return sum
}

