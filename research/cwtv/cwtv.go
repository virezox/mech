package cwtv

import (
   "github.com/89z/format"
   "net/http"
   "net/url"
)

const endpoint = "http://link.theplatform.com/s/cwtv/media/guid/2703454149/"

var LogLevel format.LogLevel

func getPlay(addr string) (string, error) {
   par, err := url.Parse(addr)
   if err != nil {
      return "", err
   }
   return par.Query().Get("play"), nil
}

func media(play string) (*url.URL, error) {
   req, err := http.NewRequest("GET", endpoint + play, nil)
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
