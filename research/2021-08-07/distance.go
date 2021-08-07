package main
import "fmt"

func variance(data []float64) float64 {
   var count, mean, M2 float64
   for _, x := range data {
      count += 1
      delta := x - mean
      mean += delta / count
      M2 += delta * (x - mean)
   }
   return M2 / (count - 1)
}

func distance(x, y int, v []float64, d [][]float64) float64 {
   var dist float64
   for i, r := range d {
      delta := r[x] - r[y]
      dist += delta * delta / v[i]
   }
   return dist
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
