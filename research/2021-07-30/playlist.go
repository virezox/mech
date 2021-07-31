package main

import (
   "fmt"
   "github.com/89z/mech/musicbrainz"
   "github.com/89z/mech/youtube"
   "time"
)

func main() {
   releaseID := "a40cb6e9-c766-37c4-8677-7eb51393d5a1"
   musicbrainz.Verbose = true
   youtube.Verbose = true
   // musicbrainz hash
   c, err := musicbrainz.NewCover(releaseID)
   if err != nil {
      panic(err)
   }
   other, err := musicbrainz.Hash(c.Images[0].Image)
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
         for _, i := range s.Items() {
            _, err := i.Distance(other)
            if err != nil {
               panic(err)
            }
            time.Sleep(100 * time.Millisecond)
         }
      }
   }
   for key, val := range youtube.Distance {
      fmt.Println(key, val)
   }
}
