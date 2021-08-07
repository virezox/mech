package main
import "fmt"

func variance(rows [][]float64) []float64 {
   M2 := make(map[int]float64)
   mean := make(map[int]float64)
   for i, row := range rows {
      for j, col := range row {
         delta := col - mean[j]
         mean[j] += delta / float64(i+1)
         M2[j] += delta * (col - mean[j])
      }
   }
   v := make([]float64, len(M2))
   for i := range v {
      v[i] = M2[i] / float64(len(rows)-1)
   }
   return v
}

func main() {
   f := [][]float64{
      {70,170},
      {60,160},
      {65,180},
      {70,200},
   }
   v := variance(f)
   fmt.Printf("%.3f\n", v) // [22.917 291.667]
}
