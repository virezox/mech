package main

import (
   "fmt"
   "github.com/89z/mech/musicbrainz"
   "github.com/89z/mech/youtube"
   "sort"
)


func main() {
   releaseID := "a40cb6e9-c766-37c4-8677-7eb51393d5a1"
   // musicbrainz hash
   c, err := musicbrainz.NewCover(releaseID)
   if err != nil {
      panic(err)
   }
   o, err := musicbrainz.Hash(c.Images[0].Image)
   if err != nil {
      panic(err)
   }
   // youtube hash
   r, err := musicbrainz.NewRelease(releaseID)
   if err != nil {
      panic(err)
   }
   artist := r.ArtistCredit[0].Name
   for _, med := range r.Media {
      for _, track := range med.Tracks {
         s, err := youtube.NewSearch(artist + " " + track.Title)
         if err != nil {
            panic(err)
         }
         r := s.Results()
         for i := range r {
            err := r[i].SetDistance(o)
            if err != nil {
               panic(err)
            }
         }
         sort.SliceStable(r, func(a, b int) bool {
            return r[a].Distance < r[b].Distance
         })
         fmt.Println(r[0].VideoID, track.Title)
      }
   }
}
