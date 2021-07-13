package picture

import (
   "fmt"
   "github.com/89z/mech/musicbrainz"
   "github.com/89z/mech/youtube"
   "time"
)

func search(query string) (string, error) {
   s, err := youtube.ISearch(query)
   if err != nil {
      return "", err
   }
   for _, vid := range s.Videos() {
      if vid.VideoID != "" {
         return vid.VideoID, nil
      }
   }
   return "", fmt.Errorf("%q fail", query)
}

func tracks() {
   r, err := musicbrainz.NewRelease("a40cb6e9-c766-37c4-8677-7eb51393d5a1")
   if err != nil {
      panic(err)
   }
   artist := r.ArtistCredit[0].Name
   for _, m := range r.Media {
      for i, track := range m.Tracks {
         id, err := search(artist + " " + track.Title)
         if err != nil {
            panic(err)
         }
         fmt.Println(i, track.Title, id)
         time.Sleep(100 * time.Millisecond)
      }
   }
}
