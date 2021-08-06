package main

import (
   "fmt"
   "gonum.org/v1/gonum/stat"
   "math"
)

type table [][]float64

func (t table) variance() []float64 {
   var f []float64
   for _, r := range t {
      f = append(f, stat.Variance(r, nil))
   }
   return f
}

func (t table) distance(x, y int, variance []float64) float64 {
   var f float64
   for i, r := range t {
      f += math.Pow(r[x]-r[y], 2) / variance[i]
   }
   return f
}

func main() {
   t := table{
      { 70,  60,  65,  70}, // height
      {170, 160, 180, 200}, // weight
   }
   v := t.variance()
   for i := range t[0] {
      d := t.distance(0, i, v)
      fmt.Printf("%.3f\n", d)
   }
}
