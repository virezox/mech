package main

import (
   "encoding/json"
   "github.com/89z/mech/musicbrainz"
   "github.com/89z/mech/youtube"
   "net/http"
   "os"
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

type track struct {
   musicbrainz.Track
   Items []item
}

type item struct {
   youtube.Item
   HQ1 int64
   HQ2 int64
}

func write() {
   musicbrainz.Verbose = true
   youtube.Verbose = true
   r, err := musicbrainz.NewRelease("a40cb6e9-c766-37c4-8677-7eb51393d5a1")
   if err != nil {
      panic(err)
   }
   artist := r.ArtistCredit[0].Name
   var tracks []track
   for _, med := range r.Media {
      for _, t := range med.Tracks {
         s, err := youtube.NewSearch(artist + " " + r.Title + " " + t.Title)
         if err != nil {
            panic(err)
         }
         var items []item
         for _, i := range s.Items() {
            // image HQ1
            p := youtube.Picture{Base: "hq1", Format: youtube.JPG}
            hq1, err := size(p, i)
            if err != nil {
               panic(err)
            }
            // image HQ2
            p = youtube.Picture{Base: "hq2", Format: youtube.JPG}
            hq2, err := size(p, i)
            if err != nil {
               panic(err)
            }
            items = append(items, item{
               i, hq1, hq2,
            })
         }
         tracks = append(tracks, track{
            t, items,
         })
      }
   }
   f, err := os.Create("returnal.json")
   if err != nil {
      panic(err)
   }
   defer f.Close()
   enc := json.NewEncoder(f)
   enc.SetIndent("", " ")
   enc.Encode(tracks)
}
