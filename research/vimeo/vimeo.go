package vimeo

import (
   "encoding/json"
   "github.com/89z/format"
   "net/http"
   "strings"
)

type Oembed struct {
   URI string
}

func NewOembed(addr string) (*Oembed, error) {
   var buf strings.Builder
   buf.WriteString("http://vimeo.com/api/oembed.json?url=")
   buf.WriteString(addr)
   req, err := http.NewRequest("GET", buf.String(), nil)
   if err != nil {
      return nil, err
   }
   logLevel.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   embed := new(Oembed)
   if err := json.NewDecoder(res.Body).Decode(embed); err != nil {
      return nil, err
   }
   return embed, nil
}

var logLevel format.LogLevel

type Download struct {
   Public_Name string
   Size int64
   Link string
}

func (d Download) String() string {
   buf := []byte("Name:")
   buf = append(buf, d.Public_Name...)
   buf = append(buf, " Size:"...)
   buf = append(buf, format.Size.GetInt64(d.Size)...)
   buf = append(buf, " Link:"...)
   buf = append(buf, d.Link...)
   return string(buf)
}

type JsonWeb struct {
   Token string
}

func NewJsonWeb() (*JsonWeb, error) {
   req, err := http.NewRequest("GET", "https://vimeo.com/_rv/jwt", nil)
   if err != nil {
      return nil, err
   }
   req.Header.Set("X-Requested-With", "XMLHttpRequest")
   logLevel.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   web := new(JsonWeb)
   if err := json.NewDecoder(res.Body).Decode(web); err != nil {
      return nil, err
   }
   return web, nil
}

func (w JsonWeb) Video(path string) (*Video, error) {
   return newVideo(path, "JWT " + w.Token)
}

type Video struct {
   Download []Download
}

func newVideo(path, auth string) (*Video, error) {
   req, err := http.NewRequest("GET", "https://api.vimeo.com" + path, nil)
   if err != nil {
      return nil, err
   }
   req.Header.Set("Authorization", auth)
   req.URL.RawQuery = "fields=download"
   logLevel.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      return nil, errorString(res.Status)
   }
   vid := new(Video)
   if err := json.NewDecoder(res.Body).Decode(vid); err != nil {
      return nil, err
   }
   return vid, nil
}

type errorString string

func (e errorString) Error() string {
   return string(e)
}
