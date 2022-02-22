package cwtv

import (
   "github.com/89z/format"
   "net/http"
)

var LogLevel format.LogLevel

func media(play string) (*http.Response, error) {
   req, err := http.NewRequest(
      "GET", "http://link.theplatform.com/s/cwtv/media/guid/2703454149/" + play,
      nil,
   )
   if err != nil {
      return nil, err
   }
   req.URL.RawQuery = "formats=M3U"
   LogLevel.Dump(req)
   // this redirects:
   return new(http.Client).Do(req)
}
