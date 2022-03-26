package pbs

import (
   "encoding/json"
   "net/http"
   "net/url"
   "github.com/89z/format"
)

type Asset struct {
   Resource struct {
      Duration int64
      MP4_Videos []struct {
         Profile string
         URL string
      }
      Title string
   }
}

type errorString string

func (e errorString) Error() string {
   return string(e)
}

var logLevel format.LogLevel

func newAsset(slug string) (*Asset, error) {
   req := new(http.Request)
   req.Header = make(http.Header)
   req.Header["Authorization"] = []string{"Basic YW5kcm9pZDpiYVhFN2h1bXVWYXQ="}
   req.URL = new(url.URL)
   req.URL.Host = "content.services.pbs.org"
   req.URL.Path = "/v3/android/screens/video-assets/" + slug + "/"
   req.URL.Scheme = "http"
   req.Header["X-Pbs-Platformversion"] = []string{"5.5.5"}
   val := make(url.Values)
   val["station_id"] = []string{"b3291387-78a4-41e1-beb0-da2f61a96a3e"}
   req.URL.RawQuery = val.Encode()
   logLevel.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      return nil, errorString(res.Status)
   }
   ass := new(Asset)
   if err := json.NewDecoder(res.Body).Decode(ass); err != nil {
      return nil, err
   }
   return ass, nil
}
