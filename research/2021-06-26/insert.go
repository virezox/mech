package main

import (
   "fmt"
   "github.com/89z/mech/youtube"
   "math/big"
   "sort"
)

func mod(a, b int64) int64 {
   c, d := big.NewInt(a), big.NewInt(b)
   return new(big.Int).Mod(c, d).Int64()
}

func score(i youtube.Image) int64 {
   return mod(480 - i.Height, 720) + i.Frame + i.Format
}

func main() {
   sort.Slice(youtube.Images, func(d, e int) bool {
      return score(youtube.Images[d]) < score(youtube.Images[e])
   })
   for _, each := range youtube.Images {
      fmt.Println(each)
   }
}
