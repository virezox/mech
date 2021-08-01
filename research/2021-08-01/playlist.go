package main

import (
   "fmt"
   "github.com/89z/mech/musicbrainz"
   "github.com/89z/mech/youtube"
   "net/http"
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

func main() {
   musicbrainz.Verbose = true
   youtube.Verbose = true
   r, err := musicbrainz.NewRelease("a40cb6e9-c766-37c4-8677-7eb51393d5a1")
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
            // duration 1
            d1 := 4*time.Minute + 19*time.Second
            // duration 2
            d2, err := i.Duration()
            if err != nil {
               panic(err)
            }
            // duration difference
            dDiff := diffDuration(d1, d2)
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
            iDiff := diffImage(i1, i2)
            // print
            fmt.Println("dDiff:", dDiff, "iDiff:", iDiff, i.VideoID(), i.Title())
         }
      }
   }
}
