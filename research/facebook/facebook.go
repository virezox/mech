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

func video(v int64) (*http.Response, error) {
   vars, err := json.Marshal(map[string]int64{"videoID": v})
   if err != nil {
      return nil, err
   }
   body := url.Values{
      "doc_id": {"7246858708718613"},
      "variables": {string(vars)},
   }.Encode()
   req, err := http.NewRequest(
      "POST", "https://www.facebook.com/api/graphql/", strings.NewReader(body),
   )
   if err != nil {
      return nil, err
   }
   req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
   LogLevel.Dump(req)
   return new(http.Transport).RoundTrip(req)
}

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
