package bleep

import (
   "encoding/json"
   "github.com/89z/format"
   "github.com/89z/format/net"
   "io"
   "net/http"
   "net/url"
   "strconv"
   "strings"
   "time"
)

var LogLevel format.LogLevel

type Meta []net.Node

func NewMeta(releaseID int64) (Meta, error) {
   buf := []byte("https://bleep.com/release/")
   buf = strconv.AppendInt(buf, releaseID, 10)
   req, err := http.NewRequest("GET", string(buf), nil)
   if err != nil {
      return nil, err
   }
   LogLevel.Dump(req)
   // this redirects, so we cannot use RoundTrip
   res, err := new(http.Client).Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   return net.ReadHTML(res.Body, "meta"), nil
}

func (m Meta) Image() string {
   for _, node := range m {
      if node.Attr["property"] == "og:image" {
         return node.Attr["content"]
      }
   }
   return ""
}

// can be either one of these:
//  2001-05-01 00:00:00.0
//  Tue May 01 00:00:00 UTC 2001
func (m Meta) ReleaseDate() (time.Time, error) {
   for _, node := range m {
      if node.Attr["property"] == "music:release_date" {
         value := node.Attr["content"]
         date, err := time.Parse(time.UnixDate, value)
         if err != nil {
            return time.Parse("2006-01-02 15:04:05.9", value)
         }
         return date, nil
      }
   }
   return time.Time{}, notFound{"music:release_date"}
}

type notFound struct {
   value string
}

func (n notFound) Error() string {
   return strconv.Quote(n.value) + " not found"
}

func Tracks(id int64) ([]Track, error) {
   val := url.Values{
      "id": {strconv.FormatInt(id, 10)},
      "type": {"ReleaseProduct"},
   }.Encode()
   req, err := http.NewRequest(
      "POST", "https://bleep.com/player/addToPlaylist", strings.NewReader(val),
   )
   if err != nil {
      return nil, err
   }
   req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
   LogLevel.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   var tracks []Track
   if err := json.NewDecoder(res.Body).Decode(&tracks); err != nil {
      return nil, err
   }
   return tracks, nil
}

type Track struct {
   ReleaseID int64
   Disc int64
   Number int64
   Artist string
   Title string
}

func (t Track) Resolve() (string, error) {
   buf := strconv.AppendInt(nil, t.ReleaseID, 10)
   buf = append(buf, '-')
   buf = strconv.AppendInt(buf, t.Disc, 10)
   buf = append(buf, '-')
   buf = strconv.AppendInt(buf, t.Number, 10)
   return string(buf)
   req, err := http.NewRequest(
      "GET", "https://bleep.com/player/resolve/" + t.String(), nil,
   )
   if err != nil {
      return "", err
   }
   LogLevel.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
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
