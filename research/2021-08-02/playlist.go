package main

import (
   "github.com/89z/mech/musicbrainz"
   "github.com/89z/mech/youtube"
   "math"
   "net/http"
   "time"
)

var cache = make(map[string]int64)

func diffDuration(x, y time.Duration) time.Duration {
   if x < y {
      return y - x
   }
   return x - y
}

func diffImage(x, y int64) int64 {
   if x < y {
      return y - x
   }
   return x - y
}

func length(p youtube.Picture, i youtube.Item) (int64, error) {
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

func newResult(t musicbrainz.Track, i youtube.Item) (*result, error) {
   // duration MB
   dMB := t.Duration()
   // duration YT
   dYT, err := i.Duration()
   if err != nil {
      return nil, err
   }
   // duration difference
   dDiff := diffDuration(dMB, dYT)
   // image HQ1
   p := youtube.Picture{Base: "hq1", Format: youtube.JPG}
   hq1, err := length(p, i)
   if err != nil {
      return nil, err
   }
   // image HQ2
   p = youtube.Picture{Base: "hq2", Format: youtube.JPG}
   hq2, err := length(p, i)
   if err != nil {
      return nil, err
   }
   // image difference
   iDiff := diffImage(hq1, hq2)
   // return
   return &result{dDiff, iDiff, i}, nil
}

type result struct {
   duration time.Duration
   contentLength int64
   youtube.Item
}

type video struct {
   duration time.Duration
   size int64
}

func relativeDifference(x, y float64) float64 {
   if x < y {
      return (y - x) / y
   }
   return (x - y) / x
}

func (p video) euclideanDistance(q video) float64 {
   a := relativeDifference(
      float64(p.duration), float64(q.duration),
   )
   b := relativeDifference(
      float64(p.size), float64(q.size),
   )
   return math.Pow(a, 2) + math.Pow(b, 2)
}
