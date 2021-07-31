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

func main() {
   for _, s := range times {
      t, err := time.Parse("15:04:05", s)
      if err != nil {
         panic(err)
      }
      fmt.Println(t)
   }
}
