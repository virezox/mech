package facebook

import (
   "encoding/json"
   "github.com/89z/format"
   "net/http"
   "net/url"
   "strconv"
   "strings"
)

var LogLevel format.LogLevel

func dash(v int64) (*http.Response, error) {
   req, err := http.NewRequest(
      "GET", "https://www.facebook.com/video/playback/dash_mpd_debug.mpd", nil,
   )
   if err != nil {
      return nil, err
   }
   req.URL.RawQuery = "v=" + strconv.FormatInt(v, 10)
   LogLevel.Dump(req)
   return new(http.Transport).RoundTrip(req)
}
