package main

import (
   "fmt"
   "math"
)

func absoluteDifference(x, y float64) float64 {
   if x < y {
      return y - x
   }
   return x - y
}

func stdDev(xs []float64) float64 {
   return math.Sqrt(variance(xs))
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

type dimensions []struct {
   x float64
   y float64
   s float64
}

func (d dimensions) mahalanobisDistance() float64 {
   var f float64
   for _, i := range d {
      f += absoluteDifference(i.x, i.y) / i.s
   }
   return f
}

func main() {
   names := []string{"description", "Adam", "Bob", "Chris"}
   heights := []float64{70, 60, 65, 70}
   weights := []float64{170, 160, 180, 200}
   for i := range names {
      d := dimensions{
         {heights[0], heights[i], stdDev(heights)},
         {weights[0], weights[i], stdDev(weights)},
      }
      fmt.Printf("%v %.3f\n", names[i], d.mahalanobisDistance())
   }
}
