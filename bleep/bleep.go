package bleep

import (
   "encoding/json"
   "github.com/89z/mech"
   "io"
   "net/http"
   "net/url"
   "strconv"
   "strings"
)

const (
   Origin = "https://bleep.com"
   cloudFront = "https://d1rgjmn2wmqeif.cloudfront.net"
)

func Image(releaseID int) string {
   return cloudFront + "/r/b/" + strconv.Itoa(releaseID) + "-1.jpg"
}

func ReleaseID(addr string) (int, error) {
   low := strings.LastIndexByte(addr, '/')
   if low == -1 {
      return 0, mech.ByteNotFound('/')
   }
   addr = addr[low+1:]
   high := strings.IndexByte(addr, '-')
   if high == -1 {
      return 0, mech.ByteNotFound('-')
   }
   return strconv.Atoi(addr[:high])
}

type Track struct {
   Disc int
   ID int
   Number int
   ReleaseID int
   Title string
}

func Release(id int) ([]Track, error) {
   val := url.Values{
      "id": {strconv.Itoa(id)},
      "type": {"ReleaseProduct"},
   }
   req, err := http.NewRequest(
      "POST", Origin + "/player/addToPlaylist", strings.NewReader(val.Encode()),
   )
   if err != nil {
      return nil, err
   }
   req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
   res, err := mech.RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   var rel []Track
   if err := json.NewDecoder(res.Body).Decode(&rel); err != nil {
      return nil, err
   }
   return rel, nil
}

func (t Track) Resolve() (string, error) {
   src := "/player/resolve/"
   src += strconv.Itoa(t.ReleaseID)
   src += "-"
   src += strconv.Itoa(t.Disc)
   src += "-"
   src += strconv.Itoa(t.Number)
   req, err := http.NewRequest("GET", Origin + src, nil)
   if err != nil {
      return "", err
   }
   res, err := mech.RoundTrip(req)
   if err != nil {
      return "", err
   }
   defer res.Body.Close()
   dst, err := io.ReadAll(res.Body)
   if err != nil {
      return "", err
   }
   return string(dst), nil
}
