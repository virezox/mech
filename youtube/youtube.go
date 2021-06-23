// YouTube
package youtube

import (
   "fmt"
   "net/http"
   "strings"
)

const (
   PlayerAPI = "https://www.youtube.com/youtubei/v1/player"
   chunk = 10_000_000
   invert = "\x1b[7m"
   reset = "\x1b[m"
)

func post(id, name, version string) (*http.Response, error) {
   body := fmt.Sprintf(`
   {
      "videoId": %q, "context": {
         "client": {"clientName": %q, "clientVersion": %q}
      }
   }
   `, id, name, version)
   req, err := http.NewRequest(
      "POST", PlayerAPI, strings.NewReader(body),
   )
   if err != nil {
      return nil, err
   }
   val := req.URL.Query()
   val.Set("key", "AIzaSyAO_FJ2SlqU8Q4STEHLGCilw_Y9_11qcW8")
   req.URL.RawQuery = val.Encode()
   fmt.Println(invert, "POST", reset, req.URL)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   if res.StatusCode != http.StatusOK {
      return nil, fmt.Errorf("status %v", res.Status)
   }
   return res, nil
}

type VideoDetails struct {
   Author string
   ShortDescription string
   Title string
   ViewCount int `json:"viewCount,string"`
}
