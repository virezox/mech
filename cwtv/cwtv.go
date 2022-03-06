package cwtv

import (
   "github.com/89z/format"
   "net/http"
   "net/url"
   "strings"
)

const mpxURL = "http://link.theplatform.com/s/cwtv/media/guid/2703454149"

var LogLevel format.LogLevel

func GetPlay(addr string) (string, error) {
   par, err := url.Parse(addr)
   if err != nil {
      return "", err
   }
   return par.Query().Get("play"), nil
}

func Media(play string) (*url.URL, error) {
   var buf strings.Builder
   buf.WriteString(mpxURL)
   buf.WriteByte('/')
   buf.WriteString(play)
   req, err := http.NewRequest("GET", buf.String(), nil)
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
