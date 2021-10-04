// MusicBrainz
package musicbrainz

import (
   "encoding/json"
   "github.com/89z/mech"
   "net/http"
   "net/url"
   "strings"
   "time"
)

const API = "http://musicbrainz.org/ws/2/release"

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
      Tracks []Track
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
   v := url.Values{
      "fmt": {"json"},
      "inc": {"artist-credits recordings"},
   }
   req, err := http.NewRequest(
      "GET", API + "/" + releaseID, strings.NewReader(v.Encode()),
   )
   if err != nil {
      return nil, err
   }
   req.Header.Set("content-type", "application/x-www-form-urlencoded")
   res, err := mech.RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   rls := new(Release)
   if err := json.NewDecoder(res.Body).Decode(rls); err != nil {
      return nil, err
   }
   return rls, nil
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

type Track struct {
   Length float64
   Title string
}

func (t Track) Duration() time.Duration {
   return time.Duration(t.Length) * time.Millisecond
}
