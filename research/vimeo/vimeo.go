package vimeo

import (
   "encoding/json"
   "net/http"
   "net/url"
)

type jsonWeb struct {
   Token string
}

func newJsonWeb() (*jsonWeb, error) {
   var req http.Request
   req.Header = make(http.Header)
   req.URL = new(url.URL)
   req.URL.Host = "vimeo.com"
   req.URL.Path = "/_rv/jwt"
   req.URL.Scheme = "https"
   req.Header["X-Requested-With"] = []string{"XMLHttpRequest"}
   res, err := new(http.Transport).RoundTrip(&req)
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

func (w jsonWeb) video(path string) (*http.Response, error) {
   var req http.Request
   req.Header = make(http.Header)
   req.URL = new(url.URL)
   req.URL.Host = "api.vimeo.com"
   req.URL.Scheme = "http"
   req.Header["Authorization"] = []string{"jwt " + w.Token}
   req.URL.Path = path
   val := make(url.Values)
   val["fields"] = []string{"download"}
   req.URL.RawQuery = val.Encode()
   return new(http.Transport).RoundTrip(&req)
}
