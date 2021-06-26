package main

import (
   "fmt"
   "math/big"
   "sort"
)

const (
   WidthAutoHeightBlack = 0
   WidthAuto = 1
   WidthBlack = 2
   HeightCrop = 3
)

var Images = []Image{
   {720, WidthBlack, "maxres3.webp"},
   {720, WidthBlack, "maxres2.webp"},
   {720, WidthBlack, "maxres1.webp"},
   {720, WidthAuto, "maxresdefault.webp"},
   {720, WidthAuto, "hq720.webp"},
   {480, WidthBlack, "sd3.webp"},
   {480, WidthBlack, "sd2.webp"},
   {480, WidthBlack, "sd1.webp"},
   {480, WidthAutoHeightBlack, "sddefault.webp"},
   {360, WidthBlack, "hq3.webp"},
   {360, WidthBlack, "hq2.webp"},
   {360, WidthBlack, "hq1.webp"},
   {360, WidthAutoHeightBlack, "hqdefault.webp"},
   {360, WidthAutoHeightBlack, "0.webp"},
   {180, WidthAuto, "mqdefault.webp"},
   {180, HeightCrop, "mq3.webp"},
   {180, HeightCrop, "mq2.webp"},
   {180, HeightCrop, "mq1.webp"},
   {90, WidthBlack, "3.webp"},
   {90, WidthBlack, "2.webp"},
   {90, WidthBlack, "1.webp"},
   {90, WidthAutoHeightBlack, "default.webp"},
}

func mod(a, b int64) int64 {
   c, d := big.NewInt(a), big.NewInt(b)
   return new(big.Int).Mod(c, d).Int64()
}

func score(i Image) int64 {
   return mod(480 - i.Height, 720) + i.Frame
}

type Image struct {
   Height int64
   Frame int64
   File string
}

func main() {
   sort.Slice(Images, func(a, b int) bool {
      c, d := Images[a], Images[b]
      return score(c) < score(d)
   })
   for _, each := range Images {
      fmt.Println(each)
   }
}
