package bleep

import (
   "encoding/json"
   "fmt"
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

// https://bleep.com/release/8728-four-tet-pause
func ReleaseID(addr string) (int, error) {
   _, id, ok := cut(addr, '/')
   if ok {
      id, _, ok := cut(id, '-')
      if ok {
         return strconv.Atoi(id)
      }
   }
   return 0, mech.Invalid{addr}
}

func cut(s string, sep byte) (string, string, bool) {
   i := strings.IndexByte(s, sep)
   if i == -1 {
      return s, "", false
   }
   return s[:i], s[i+1:], true
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

func (t Track) Error() string {
   return fmt.Sprintf("%+v not found", t)
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
