package main

import (
   "fmt"
   "github.com/89z/mech/musicbrainz"
   "github.com/89z/mech/youtube"
   "net/http"
   "sort"
   "time"
)

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


var cache = make(map[string]int64)

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

func main() {
   musicbrainz.Verbose = true
   r, err := musicbrainz.NewRelease("a40cb6e9-c766-37c4-8677-7eb51393d5a1")
   if err != nil {
      panic(err)
   }
   artist := r.ArtistCredit[0].Name
   for _, med := range r.Media {
      for _, t := range med.Tracks {
         s, err := youtube.NewSearch(artist + " " + t.Title)
         if err != nil {
            panic(err)
         }
         var results []*result
         for _, i := range s.Items() {
            r, err := newResult(t, i)
            if err != nil {
               panic(err)
            }
            results = append(results, r)
            time.Sleep(100 * time.Millisecond)
         }
         sort.Slice(results, func(a, b int) bool {
            ra, rb := results[a], results[b]
            if ra.duration < rb.duration {
               return true
            }
            if rb.duration < ra.duration {
               return false
            }
            return ra.contentLength < rb.contentLength
         })
         r := results[0]
         fmt.Println(
            "time:", r.duration, "image:", r.contentLength,
            r.VideoID(), r.Title(),
         )
      }
   }
}

type result struct {
   duration time.Duration
   contentLength int64
   youtube.Item
}
