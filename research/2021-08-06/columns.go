package main

import (
   "fmt"
   "gonum.org/v1/gonum/stat"
   "math"
)

type matrix [][]float64

func (m matrix) variance() []float64 {
   var f []float64
   for _, r := range m {
      f = append(f, stat.Variance(r, nil))
   }
   return f
}

func (m matrix) distance(x, y int, variance []float64) float64 {
   var f float64
   for i, r := range m {
      f += math.Pow(r[x]-r[y], 2) / variance[i]
   }
   return f
}

func main() {
   // suspects as columns
   m := matrix{
      { 70,  60,  65,  70}, // height
      {170, 160, 180, 200}, // weight
   }
   v := m.variance()
   for i := range m[0] {
      d := m.distance(0, i, v)
      fmt.Printf("%.3f\n", d)
   }
}
