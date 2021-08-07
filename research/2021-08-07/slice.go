package main
import "fmt"

func column(data [][]float64, i int) []float64 {
   var c []float64
   for _, r := range data {
      c = append(c, r[i])
   }
   return c
}

func distance(x, y, variances []float64) float64 {
   var f float64
   for i := range x {
      f += (x[i] - y[i]) * (x[i] - y[i]) / variances[i]
   }
   return f
}

func variance(xs []float64) float64 {
   mean, M2 := 0.0, 0.0
   for n, x := range xs {
      delta := x - mean
      mean += delta / float64(n+1)
      M2 += delta * (x - mean)
   }
   return M2 / float64(len(xs)-1)
}

func main() {
   head := []string{"description", "Adam", "Bob", "Chris"}
   body := [][]float64{
      { 70,  60,  65,  70},
      {170, 160, 180, 200},
   }
   var v []float64
   for _, r := range body {
      v = append(v, variance(r))
   }
   for i := range head {
      x := column(body, 0)
      y := column(body, i)
      fmt.Printf("%v %.3f\n", head[i], distance(x, y, v))
   }
}
