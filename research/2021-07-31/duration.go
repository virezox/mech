package main

import (
   "fmt"
   "github.com/89z/mech/youtube"
   "sort"
   "time"
)

func distance(d, e time.Duration) time.Duration {
   if d < e {
      return e - d
   }
   return d - e
}

func main() {
   youtube.Verbose = true
   s, err := youtube.NewSearch("oneohtrix point never describing bodies")
   if err != nil {
      panic(err)
   }
   var items []item
   for _, i := range s.Items() {
      d, err := i.Duration()
      if err != nil {
         panic(err)
      }
      items = append(items, item{
         distance(d, 4*time.Minute+19*time.Second), i,
      })
   }
   sort.Slice(items, func(a, b int) bool {
      return items[a].distance < items[b].distance
   })
   for _, i := range items {
      fmt.Println(i.distance, i.VideoID(), i.Title())
   }
}

type item struct {
   distance time.Duration
   youtube.Item
}
