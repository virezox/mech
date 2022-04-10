package youtube

import (
   "encoding/json"
   "fmt"
   "github.com/89z/format"
   "github.com/89z/mech/youtube"
   "io"
   "net/http"
   "net/url"
   "strings"
   "testing"
)

var body = strings.NewReader(`{
  "context": {
    "client": {
      "clientName": "TVHTML5_SIMPLY_EMBEDDED_PLAYER",
      "clientVersion": "2.0"
    }
  },
  "videoId": "Tq92D6wQ1mg"
}
`)

func TestYouTube(t *testing.T) {
   req := new(http.Request)
   req.Body = io.NopCloser(body)
   req.Header = make(http.Header)
   req.Header["X-Goog-Api-Key"] = []string{"AIzaSyAO_FJ2SlqU8Q4STEHLGCilw_Y9_11qcW8"}
   req.Method = "POST"
   req.URL = new(url.URL)
   req.URL.Host = "www.youtube.com"
   req.URL.Path = "/youtubei/v1/player"
   req.URL.Scheme = "https"
   format.LogLevel(0).Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      t.Fatal(err)
   }
   defer res.Body.Close()
   var play youtube.Player
   if err := json.NewDecoder(res.Body).Decode(&play); err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%a\n", play.StreamingData.AdaptiveFormats[0])
}
