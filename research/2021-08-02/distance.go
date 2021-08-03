package main

import (
   "github.com/89z/mech/musicbrainz"
   "github.com/89z/mech/youtube"
   "math"
   "net/http"
)

var cache = make(map[string]int64)

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

func relativeDifference(x, y float64) float64 {
   if x < y {
      return (y - x) / y
   }
   return (x - y) / x
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
   dDiff := 100 * relativeDifference(
      float64(dMB), float64(dYT),
   )
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
   // size difference
   sDiff := 100 * relativeDifference(
      float64(hq1), float64(hq2),
   )
   // return
   return &result{
      dDiff, sDiff, math.Pow(dDiff, 2) + math.Pow(sDiff, 2), i,
   }, nil
}

type result struct {
   duration float64
   size float64
   distance float64
   youtube.Item
}
