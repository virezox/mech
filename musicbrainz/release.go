package musicbrainz

import (
   "encoding/json"
   "net/http"
   "net/http/httputil"
   "net/url"
   "os"
   "strings"
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
   d, err := httputil.DumpRequest(req, true)
   if err != nil {
      return nil, err
   }
   os.Stdout.Write(append(d, '\n'))
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
