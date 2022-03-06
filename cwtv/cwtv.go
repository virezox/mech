package cwtv

import (
   "encoding/json"
   "github.com/89z/format"
   "net/http"
   "net/url"
   "strings"
)

var LogLevel format.LogLevel

func GetPlay(addr string) (string, error) {
   par, err := url.Parse(addr)
   if err != nil {
      return "", err
   }
   return par.Query().Get("play"), nil
}

type Video struct {
   Video struct {
      MPX_URL string
      Series_Name string
      Title string
   }
}

func NewVideo(play string) (*Video, error) {
   var buf strings.Builder
   buf.WriteString("http://images.cwtv.com/feed/mobileapp/video-meta")
   buf.WriteString("/apiversion_8/guid_")
   buf.WriteString(play)
   req, err := http.NewRequest("GET", buf.String(), nil)
   if err != nil {
      return nil, err
   }
   LogLevel.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   vid := new(Video)
   if err := json.NewDecoder(res.Body).Decode(vid); err != nil {
      return nil, err
   }
   return vid, nil
}

func (v Video) Base() string {
   var buf strings.Builder
   buf.WriteString(v.Video.Series_Name)
   buf.WriteByte('-')
   buf.WriteString(v.Video.Title)
   return buf.String()
}

func (v Video) Media() (*url.URL, error) {
   req, err := http.NewRequest("GET", v.Video.MPX_URL, nil)
   if err != nil {
      return nil, err
   }
   req.URL.RawQuery = "formats=M3U"
   LogLevel.Dump(req)
   // This redirects, but we only care about the URL, so we dont need to follow:
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   return res.Location()
}
