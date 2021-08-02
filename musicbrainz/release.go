package musicbrainz

import (
   "encoding/json"
   "fmt"
   "net/http"
   "net/http/httputil"
   "net/url"
   "os"
   "strings"
   "time"
)

const API = "http://musicbrainz.org/ws/2/release"

var Verbose = false

type Cover struct {
   Images []struct {
      Image string
   }
}

func NewCover(releaseID string) (*Cover, error) {
   addr := fmt.Sprintf(
      "http://archive.org/download/mbid-%v/index.json", releaseID,
   )
   if Verbose {
      fmt.Println("GET", addr)
   }
   res, err := http.Get(addr)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   cov := new(Cover)
   if err := json.NewDecoder(res.Body).Decode(cov); err != nil {
      return nil, err
   }
   return cov, nil
}

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
   if Verbose {
      d, err := httputil.DumpRequest(req, true)
      if err != nil {
         return nil, err
      }
      os.Stdout.Write(append(d, '\n'))
   }
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

type Track struct {
   Length int
   Title string
}

func (t Track) Duration() time.Duration {
   return time.Duration(t.Length) * time.Millisecond
}
