package internal

import (
   "io"
   "net/http"
   "net/url"
   "fmt"
   "time"
   "encoding/json"
   "strings"
   "testing"
)

func TestBravo(t *testing.T) {
   var req http.Request
   req.Header = make(http.Header)
   req.Header["Accept"] = []string{"text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8"}
   req.Header["Accept-Language"] = []string{"en-us,en;q=0.5"}
   req.Header["Connection"] = []string{"close"}
   req.Header["Content-Type"] = []string{"application/json"}
   req.Header["Cookie"] = []string{"PREF=hl=en&tz=UTC; CONSENT=YES+cb.20210328-17-p0.en+FX+617; GPS=1; YSC=4KzZ7fZygtI; VISITOR_INFO1_LIVE=Jau7HPnhETc"}
   req.Header["Host"] = []string{"www.youtube.com"}
   req.Header["Origin"] = []string{"https://www.youtube.com"}
   req.Header["Sec-Fetch-Mode"] = []string{"navigate"}
   req.Header["User-Agent"] = []string{"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.19 Safari/537.36"}
   req.Header["X-Youtube-Client-Name"] = []string{"85"}
   req.Header["X-Youtube-Client-Version"] = []string{"2.0"}
   req.Method = "POST"
   req.URL = new(url.URL)
   req.URL.Host = "www.youtube.com"
   req.URL.Path = "/youtubei/v1/player"
   req.URL.RawPath = ""
   val := make(url.Values)
   val["key"] = []string{"AIzaSyAO_FJ2SlqU8Q4STEHLGCilw_Y9_11qcW8"}
   val["prettyPrint"] = []string{"false"}
   req.URL.RawQuery = val.Encode()
   req.URL.Scheme = "https"
   for name, version := range clients {
      req.Body = io.NopCloser(strings.NewReader(fmt.Sprintf(`
      {
     "context": {
       "client": {
         "clientName": %q,
         "clientVersion": %q,
         "hl": "en",
         "timeZone": "UTC",
         "utcOffsetMinutes": 0
       },
       "thirdParty": {
         "embedUrl": "https://www.youtube.com/"
       }
     },
     "videoId": "HsUATh_Nc2U",
     "playbackContext": {
       "contentPlaybackContext": {
         "html5Preference": "HTML5_PREF_WANTS",
         "signatureTimestamp": 19088
       }
     },
     "contentCheckOk": true,
     "racyCheckOk": true
   }
      `, name, version)))
      res, err := new(http.Transport).RoundTrip(&req)
      if err != nil {
         t.Fatal(err)
      }
      var play player
      if err := json.NewDecoder(res.Body).Decode(&play); err != nil {
         t.Fatal(err)
      }
      if play.PlayabilityStatus.Status == "OK" {
         fmt.Println("pass", name)
      } else {
         fmt.Println(".")
      }
      if err := res.Body.Close(); err != nil {
         t.Fatal(err)
      }
      time.Sleep(time.Second)
   }
}
