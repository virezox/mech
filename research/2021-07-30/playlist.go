package main

import (
   "fmt"
   "github.com/89z/mech/musicbrainz"
   "github.com/89z/mech/youtube"
   "time"
)

var distances = make(map[string]int)

func distance(i youtube.Item, other *goimagehash.ImageHash) (int, error) {
   if d, ok := distances[i.VideoID]; ok {
      return d, nil
   }
   return i.Distance(other)
}

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
   var (
      artist = r.ArtistCredit[0].Name
      distance = make(map[string]int)
   )
   for _, med := range r.Media {
      for _, track := range med.Tracks {
         s, err := youtube.NewSearch(artist + " " + track.Title)
         if err != nil {
            panic(err)
         }
         for _, i := range s.Items() {
            if _, ok := distance[i.VideoID]; ok {
               continue
            }
            d, err := i.Distance(other)
            if err != nil {
               panic(err)
            }
            distance[i.VideoID] = d
            time.Sleep(100 * time.Millisecond)
         }
      }
   }
   for key, val := range distance {
      fmt.Println(key, val)
   }
}
