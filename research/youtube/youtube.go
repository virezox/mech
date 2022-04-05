package youtube

import (
   "bytes"
   "github.com/89z/format"
   "github.com/89z/format/json"
   "github.com/89z/googleplay"
   "net/http"
   "os"
   stdjson "encoding/json"
)

type token struct {
   Access_Token string
}

var logLevel format.LogLevel

func appVersion(app string, tv bool) (string, error) {
   cache, err := os.UserCacheDir()
   if err != nil {
      return "", err
   }
   token, err := googleplay.OpenToken(cache, "googleplay/token.json")
   if err != nil {
      return "", err
   }
   elem := "googleplay/phone.json"
   if tv {
      elem = "googleplay/tv.json"
   }
   phone, err := googleplay.OpenDevice(cache, elem)
   if err != nil {
      return "", err
   }
   head, err := token.Header(phone)
   if err != nil {
      return "", err
   }
   detail, err := head.Details(app)
   if err != nil {
      return "", err
   }
   return string(detail.VersionString), nil
}

func post(name, version string) (*http.Response, error) {
   var play player
   play.VideoID = "eZHsmb4ezEk"
   play.Context.Client.ClientName = name
   play.Context.Client.ClientVersion = version
   buf := new(bytes.Buffer)
   if err := stdjson.NewEncoder(buf).Encode(play); err != nil {
      return nil, err
   }
   req, err := http.NewRequest(
      "POST", "https://www.youtube.com/youtubei/v1/player", buf,
   )
   if err != nil {
      return nil, err
   }
   // AIzaSy
   req.Header.Set("X-Goog-Api-Key", "AIzaSyAO_FJ2SlqU8Q4STEHLGCilw_Y9_11qcW8")
   logLevel.Dump(req)
   return new(http.Transport).RoundTrip(req)
}

type player struct {
   VideoID string `json:"videoId"`
   Context struct {
      Client struct {
         ClientName string `json:"clientName"`
         ClientVersion string `json:"clientVersion"`
      } `json:"client"`
   } `json:"context"`
}

func newVersion(addr, agent string) (string, error) {
   req, err := http.NewRequest("GET", addr, nil)
   if err != nil {
      return "", err
   }
   if agent != "" {
      req.Header.Set("User-Agent", agent)
   }
   logLevel.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return "", err
   }
   defer res.Body.Close()
   sep := []byte(`"client":`)
   var client struct {
      ClientVersion string
   }
   if err := json.Decode(res.Body, sep, &client); err != nil {
      return "", err
   }
   return client.ClientVersion, nil
}
