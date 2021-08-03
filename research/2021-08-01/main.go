package main

import (
   "fmt"
   "github.com/89z/mech/musicbrainz"
   "github.com/89z/mech/youtube"
   "sort"
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
         s, err := youtube.NewSearch(artist + " " + r.Title + " " + t.Title)
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
         }
         sort.Slice(results, func(a, b int) bool {
            return results[a].distance < results[b].distance
         })
         r := results[0]
         fmt.Printf(
            "%.3f %v %v\n", r.distance, r.VideoID(), r.Title(),
         )
      }
   }
}
