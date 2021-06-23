// MusicBrainz
package musicbrainz

import (
   "encoding/json"
   "fmt"
   "net/http"
   "sort"
   "strconv"
)

const (
   API = "http://musicbrainz.org/ws/2/release"
   invert = "\x1b[7m"
   reset = "\x1b[m"
)

var Status = map[string]int{"Official": 0, "Bootleg": 1}

type Group struct {
   ReleaseCount int `json:"release-count"`
   Releases []Release
}

func (g Group) Sort() {
   sort.Slice(g.Releases, func(d, e int) bool {
      one, two := g.Releases[d], g.Releases[e]
      // 1. STATUS
      if Status[one.Status] < Status[two.Status] {
         return true
      }
      if Status[one.Status] > Status[two.Status] {
         return false
      }
      // 2. YEAR
      if one.date(4) < two.date(4) {
         return true
      }
      if one.date(4) > two.date(4) {
         return false
      }
      // 3. TRACKS
      if one.trackLen() < two.trackLen() {
         return true
      }
      if one.trackLen() > two.trackLen() {
         return false
      }
      // 4. DATE
      return one.date(10) < two.date(10)
   })
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

func NewGroup(groupID string) (*Group, error) {
   req, err := http.NewRequest("GET", API, nil)
   if err != nil {
      return nil, err
   }
   val := req.URL.Query()
   val.Set("fmt", "json")
   val.Set("inc", "artist-credits recordings")
   val.Set("release-group", groupID)
   req.URL.RawQuery = val.Encode()
   fmt.Println(invert, "GET", reset, req.URL)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   g := new(Group)
   json.NewDecoder(res.Body).Decode(g)
   return g, nil
}

func GroupFromArtist(artistID string, offset int) (*Group, error) {
   req, err := http.NewRequest("GET", API, nil)
   if err != nil {
      return nil, err
   }
   val := req.URL.Query()
   val.Set("artist", artistID)
   val.Set("fmt", "json")
   val.Set("inc", "release-groups")
   val.Set("limit", "100")
   val.Set("status", "official")
   val.Set("type", "album")
   if offset > 0 {
      val.Set("offset", strconv.Itoa(offset))
   }
   req.URL.RawQuery = val.Encode()
   fmt.Println(invert, "GET", reset, req.URL)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   g := new(Group)
   json.NewDecoder(res.Body).Decode(g)
   return g, nil
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
   fmt.Println(invert, "GET", reset, req.URL)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   r := new(Release)
   json.NewDecoder(res.Body).Decode(r)
   return r, nil
}
