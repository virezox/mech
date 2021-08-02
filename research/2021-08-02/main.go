package main

import (
   "fmt"
   "github.com/89z/mech/musicbrainz"
   "github.com/89z/mech/youtube"
   "sort"
   "time"
)

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
