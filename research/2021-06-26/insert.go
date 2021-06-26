package main

import (
   "fmt"
   "math/big"
   "sort"
)

func mod(a, b int64) int64 {
   c, d := big.NewInt(a), big.NewInt(b)
   return new(big.Int).Mod(c, d).Int64()
}

func score(i Image) int64 {
   return mod(480 - i.Height, 720) + i.Frame + i.Format
}

func main() {
   sort.Slice(Images, func(d, e int) bool {
      return score(Images[d]) < score(Images[e])
   })
   for _, each := range Images {
      fmt.Println(each)
   }
}
