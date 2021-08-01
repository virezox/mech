package main

import (
   "fmt"
   "github.com/89z/mech/youtube"
   "time"
)

func distance(d, e time.Duration) time.Duration {
   if d < e {
      return e - d
   }
   return d - e
}

type item struct {
   id string
   title string
   distance time.Duration
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
         i.VideoID(), i.Title(), distance(d, 4*time.Minute+19*time.Second),
      })
   }
   for _, i := range items {
      fmt.Println(i)
   }
}
