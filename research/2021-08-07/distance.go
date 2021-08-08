package main

func distance(x, y int, v []float64, data [][]float64) float64 {
   var sum float64
   for i, r := range data {
      delta := r[x] - r[y]
      sum += delta * delta / v[i]
   }
   return sum
}

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
