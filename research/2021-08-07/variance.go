package main

// Adam 4.706
// Bob 1.434
// Chris 3.086

func variance(xs []float64) float64 {
   var mean, M2, n float64
   for _, x := range xs {
      delta := x - mean
      n += 1
      mean += delta / n
      M2 += delta * (x - mean)
   }
   return M2 / (n - 1)
}
