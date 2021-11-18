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

const Origin = "https://bleep.com"

func ReleaseID(addr string) (int64, error) {
   low := strings.LastIndexByte(addr, '/')
   if low == -1 {
      return 0, mech.ByteNotFound('/')
   }
   addr = addr[low+1:]
   high := strings.IndexByte(addr, '-')
   if high == -1 {
      return 0, mech.ByteNotFound('-')
   }
   return strconv.ParseInt(addr[:high], 10, 64)
}

type Track struct {
   Title string
   ReleaseID int64
   Disc int64
   Number int64
}

func Release(id int64) ([]Track, error) {
   val := url.Values{
      "id": {strconv.FormatInt(id, 10)},
      "type": {"ReleaseProduct"},
   }
   res, err := http.PostForm(Origin + "/player/addToPlaylist", val)
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
   src := []byte(Origin)
   src = append(src, "/player/resolve/"...)
   src = strconv.AppendInt(src, t.ReleaseID, 10)
   src = append(src, '-')
   src = strconv.AppendInt(src, t.Disc, 10)
   src = append(src, '-')
   src = strconv.AppendInt(src, t.Number, 10)
   res, err := http.Get(string(src))
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
