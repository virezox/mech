package youtube

import (
   "bytes"
   "encoding/json"
   "github.com/89z/format"
   "net/http"
   "os"
   gp "github.com/89z/googleplay"
)

var logLevel format.LogLevel

func appVersion(app string) (string, error) {
   cache, err := os.UserCacheDir()
   if err != nil {
      return "", err
   }
   token, err := gp.OpenToken(cache, "googleplay/token.json")
   if err != nil {
      return "", err
   }
   device, err := gp.OpenDevice(cache, "googleplay/device.json")
   if err != nil {
      return "", err
   }
   head, err := token.Header(device)
   if err != nil {
      return "", err
   }
   det, err := head.Details(app)
   if err != nil {
      return "", err
   }
   return string(det.VersionString), nil
}

func post(name, version string) (*http.Response, error) {
   var play player
   play.VideoID = "eZHsmb4ezEk"
   play.Context.Client.ClientName = name
   play.Context.Client.ClientVersion = version
   buf := new(bytes.Buffer)
   if err := json.NewEncoder(buf).Encode(play); err != nil {
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
