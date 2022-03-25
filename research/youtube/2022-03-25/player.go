package main

import (
   "io"
   "net/http"
   "net/http/httputil"
   "net/url"
   "os"
   "strings"
)

func main() {
   var req http.Request
   req.Body = io.NopCloser(body)
   req.Header = make(http.Header)
   req.Header["Accept"] = []string{"text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8"}
   req.Header["Accept-Language"] = []string{"en-us,en;q=0.5"}
   req.Header["Connection"] = []string{"close"}
   req.Header["Content-Type"] = []string{"application/json"}
   req.Header["Cookie"] = []string{"PREF=hl=en&tz=UTC; CONSENT=YES+cb.20210328-17-p0.en+FX+854; GPS=1; YSC=K2BdHNo_Yys; VISITOR_INFO1_LIVE=kK46uI5xLYc"}
   req.Header["Host"] = []string{"www.youtube.com"}
   req.Header["Origin"] = []string{"https://www.youtube.com"}
   req.Header["Sec-Fetch-Mode"] = []string{"navigate"}
   req.Header["User-Agent"] = []string{"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.115 Safari/537.36"}
   req.Header["X-Youtube-Client-Name"] = []string{"55"}
   req.Header["X-Youtube-Client-Version"] = []string{"16.49"}
   req.Method = "POST"
   req.URL = new(url.URL)
   req.URL.Host = "www.youtube.com"
   req.URL.Path = "/youtubei/v1/player"
   req.URL.RawPath = ""
   val := make(url.Values)
   val["key"] = []string{"AIzaSyCjc_pVEDi4qsv5MtC2dMXzpIaDoRFLsxw"}
   val["prettyPrint"] = []string{"false"}
   req.URL.RawQuery = val.Encode()
   req.URL.Scheme = "https"
   res, err := new(http.Transport).RoundTrip(&req)
   if err != nil {
      panic(err)
   }
   defer res.Body.Close()
   buf, err := httputil.DumpResponse(res, true)
   if err != nil {
      panic(err)
   }
   os.Stdout.Write(buf)
}

var body = strings.NewReader(`
{
  "context": {
    "client": {
      "clientName": "ANDROID_EMBEDDED_PLAYER",
      "clientVersion": "16.49",
      "hl": "en",
      "timeZone": "UTC",
      "utcOffsetMinutes": 0
    },
    "thirdParty": {
      "embedUrl": "https://google.com"
    }
  },
  "videoId": "HtVdAasjOgU",
  "playbackContext": {
    "contentPlaybackContext": {
      "html5Preference": "HTML5_PREF_WANTS"
    }
  },
  "contentCheckOk": true,
  "racyCheckOk": true
}
`)
