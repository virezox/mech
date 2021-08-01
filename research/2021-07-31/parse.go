package main

import (
   "fmt"
   "time"
)

var times = []string{
   "4:19",
   "4:27",
   "4:20",
   "4:19",
   "42:11",
   "7:43",
   "3:30",
   "7:33",
   "4:03",
   "1:47:20",
   "4:36",
   "43:22",
   "3:33",
   "2:23",
}

// 1768 ns/op
func leftRight(r string) (time.Time, error) {
   l := "00:00:00"
   return time.Parse("15:04:05", l[:len(l)-len(r)] + r)
}

// 1137 ns/op
func minHour(s string) (time.Time, error) {
   t, err := time.Parse("4:05", s)
   if err == nil {
      return t, nil
   }
   return time.Parse("15:04:05", s)
}

func main() {
   for _, r := range times {
      t, err := leftRight(r)
      if err != nil {
         panic(err)
      }
      fmt.Println(t)
   }
}
