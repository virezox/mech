package main

import (
   "github.com/89z/mech/musicbrainz"
   "github.com/89z/mech/youtube"
   "math"
   "net/http"
   "time"
)

var cache = make(map[string]int64)

func relativeDifference(x, y float64) float64 {
   if x < y {
      return (y - x) / y
   }
   return (x - y) / x
}

func size(p youtube.Picture, i youtube.Item) (int64, error) {
   addr := p.Address(i.VideoID())
   if l, ok := cache[addr]; ok {
      return l, nil
   }
   r, err := http.Head(addr)
   if err != nil {
      return 0, err
   }
   cache[addr] = r.ContentLength
   return r.ContentLength, nil
}

type result struct {
   duration time.Duration
   size int64
   youtube.Item
}

func newResult(t musicbrainz.Track, i youtube.Item) (*result, error) {
   /*
   // duration MB
   dMB := t.Duration()
   // duration YT
   dYT, err := i.Duration()
   if err != nil {
      return nil, err
   }
   */
   // duration difference
   var dDiff time.Duration = 1
   /*
   // image HQ1
   p := youtube.Picture{Base: "hq1", Format: youtube.JPG}
   hq1, err := size(p, i)
   if err != nil {
      return nil, err
   }
   // image HQ2
   p = youtube.Picture{Base: "hq2", Format: youtube.JPG}
   hq2, err := size(p, i)
   if err != nil {
      return nil, err
   }
   */
   // image difference
   var iDiff int64 = 1
   // return
   return &result{dDiff, iDiff, i}, nil
}

func (p result) euclideanDistance(q result) float64 {
   a := relativeDifference(
      float64(p.duration), float64(q.duration),
   )
   b := relativeDifference(
      float64(p.size), float64(q.size),
   )
   return math.Pow(a, 2) + math.Pow(b, 2)
}
