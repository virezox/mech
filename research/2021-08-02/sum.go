package main

import (
   "fmt"
   "math"
   "time"
)

type video struct {
   duration time.Duration
   size int64
}

func relativeDifference(x, y float64) float64 {
   if x < y {
      return (y - x) / y
   }
   return (x - y) / x
}

func (p video) euclideanDistance(q video) float64 {
   a := relativeDifference(
      float64(p.duration), float64(q.duration),
   )
   b := relativeDifference(
      float64(p.size), float64(q.size),
   )
   return math.Pow(a, 2) + math.Pow(b, 2)
}

func main() {
   p := video{time.Minute, 10500}
   q := video{time.Hour, 20000}
   f := p.euclideanDistance(q)
   fmt.Println(f)
}
