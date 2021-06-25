// MusicBrainz
package musicbrainz

import (
   "encoding/json"
   "net/http"
)

const API = "http://musicbrainz.org/ws/2/release"

var Status = map[string]int{"Official": 0, "Bootleg": 1}

type Release struct {
   ArtistCredit []struct {
      Name string
      Artist struct {
         ID string
      }
   } `json:"artist-credit"`
   Date string
   Media []struct {
      TrackCount int `json:"track-count"`
      Tracks []struct {
         Length int
         Title string
      }
   }
   ReleaseGroup struct {
      FirstReleaseDate string `json:"first-release-date"`
      ID string
      SecondaryTypes []string `json:"secondary-types"`
      Title string
   } `json:"release-group"`
   Status string
   Title string
}

func NewRelease(releaseID string) (*Release, error) {
   req, err := http.NewRequest("GET", API + "/" + releaseID, nil)
   if err != nil {
      return nil, err
   }
   val := req.URL.Query()
   val.Set("fmt", "json")
   val.Set("inc", "artist-credits recordings")
   req.URL.RawQuery = val.Encode()
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   r := new(Release)
   json.NewDecoder(res.Body).Decode(r)
   return r, nil
}

func (r Release) date(width int) string {
   start := len(r.Date)
   right := "9999-12-31"[start:]
   return (r.Date + right)[:width]
}

func (r Release) trackLen() int {
   var count int
   for _, media := range r.Media {
      count += media.TrackCount
   }
   return count
}
