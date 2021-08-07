package main

import "fmt"

func distance(x, y int, v []float64, d [][]float64) float64 {
   var f float64
   for i, r := range d {
      f += (r[x] - r[y]) * (r[x] - r[y]) / v[i]
   }
   return f
}

func main() {
   head := []string{"description", "Adam", "Bob", "Chris"}
   body := [][]float64{
      {70, 60, 65, 70},
      {170, 160, 180, 200},
   }
   var v []float64
   for _, r := range body {
      v = append(v, variance(r))
   }
   for i := range head {
      fmt.Printf("%v %.3f\n", head[i], distance(0, i, v, body))
   }
}
