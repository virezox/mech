package mech

import (
   "strconv"
)

type NumberFormat []string

func FormatNumber() NumberFormat {
   return NumberFormat{"", " K", " M", " B", " T"}
}

func FormatSize() NumberFormat {
   return NumberFormat{" B", " kB", " MB", " GB", " TB"}
}

func FormatRate() NumberFormat {
   return NumberFormat{" B/s", " kB/s", " MB/s", " GB/s", " TB/s"}
}

func (n NumberFormat) FormatFloat(f float64) string {
   var symbol string
   for _, symbol = range n {
      if f < 1000 {
         break
      }
      f /= 1000
   }
   return strconv.FormatFloat(f, 'f', 3, 64) + symbol
}

func (n NumberFormat) FormatInt(i int64) string {
   f := float64(i)
   return n.FormatFloat(f)
}

func (n NumberFormat) FormatUint(i uint64) string {
   f := float64(i)
   return n.FormatFloat(f)
}
