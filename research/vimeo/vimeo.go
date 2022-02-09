package vimeo

import (
   "encoding/json"
   "github.com/89z/format"
   "net/http"
)

var logLevel format.LogLevel

type jsonWeb struct {
   Token string
}

func newJsonWeb() (*jsonWeb, error) {
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
   web := new(jsonWeb)
   if err := json.NewDecoder(res.Body).Decode(web); err != nil {
      return nil, err
   }
   return web, nil
}

type download struct {
   Public_Name string
   Size int64
   Link string
}

func (d download) String() string {
   buf := []byte("Name:")
   buf = append(buf, d.Public_Name...)
   buf = append(buf, " Size:"...)
   buf = append(buf, format.Size.GetInt64(d.Size)...)
   buf = append(buf, " Link:"...)
   buf = append(buf, d.Link...)
   return string(buf)
}

type video struct {
   Download []download
}

func (w jsonWeb) video(path string) (*video, error) {
   req, err := http.NewRequest("GET", "http://api.vimeo.com" + path, nil)
   if err != nil {
      return nil, err
   }
   req.Header.Set("Authorization", "jwt " + w.Token)
   req.URL.RawQuery = "fields=download"
   logLevel.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   vid := new(video)
   if err := json.NewDecoder(res.Body).Decode(vid); err != nil {
      return nil, err
   }
   return vid, nil
}
