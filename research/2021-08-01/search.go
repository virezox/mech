package main

import (
   "fmt"
   "github.com/89z/mech/youtube"
   "net/http"
   "time"
)

func durationDiff(x, y time.Duration) time.Duration {
   if x < y {
      return y - x
   }
   return x - y
}

func imageDiff(x, y int64) int64 {
   if x < y {
      return y - x
   }
   return x - y
}

func main() {
   youtube.Verbose = true
   s, err := youtube.NewSearch("oneohtrix point never describing bodies")
   if err != nil {
      panic(err)
   }
   for _, i := range s.Items() {
      // duration 1
      d1 := 4*time.Minute + 19*time.Second
      // duration 2
      d2, err := i.Duration()
      if err != nil {
         panic(err)
      }
      // duration difference
      dDiff := durationDiff(d1, d2)
      // image 1
      p := youtube.Picture{Base: "hq1", Format: youtube.JPG}
      time.Sleep(100 * time.Millisecond)
      r, err := http.Head(p.Address(i.VideoID()))
      if err != nil {
         panic(err)
      }
      i1 := r.ContentLength
      // image 2
      p = youtube.Picture{Base: "hq2", Format: youtube.JPG}
      time.Sleep(100 * time.Millisecond)
      r, err = http.Head(p.Address(i.VideoID()))
      if err != nil {
         panic(err)
      }
      i2 := r.ContentLength
      // image difference
      iDiff := imageDiff(i1, i2)
      // print
      fmt.Println("dDiff:", dDiff, "iDiff:", iDiff, i.VideoID(), i.Title())
   }
}
