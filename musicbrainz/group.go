package musicbrainz

import (
   "encoding/json"
   "fmt"
   "net/http"
   "sort"
   "strconv"
)

type Group struct {
   ReleaseCount int `json:"release-count"`
   Releases []*Release
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

