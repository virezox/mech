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

var Names = map[string]string{
   "ANDROID": "",
   "ANDROID_CASUAL": "",
   "ANDROID_CREATOR": "",
   "ANDROID_EMBEDDED_PLAYER": "",
   "ANDROID_GAMING": "",
   "ANDROID_INSTANT": "",
   "ANDROID_KIDS": "",
   "ANDROID_LITE": "",
   "ANDROID_MUSIC": "",
   "ANDROID_PRODUCER": "",
   "ANDROID_SPORTS": "",
   "ANDROID_TESTSUITE": "",
   "ANDROID_TV": "",
   "ANDROID_TV_KIDS": "",
   "ANDROID_UNPLUGGED": "",
   "ANDROID_VR": "",
   "ANDROID_WITNESS": "",
   "CLIENTX": "",
   "GOOGLE_ASSISTANT": "",
   "GOOGLE_MEDIA_ACTIONS": "",
   "IOS": "",
   "IOS_CREATOR": "",
   "IOS_DIRECTOR": "",
   "IOS_EMBEDDED_PLAYER": "",
   "IOS_GAMING": "",
   "IOS_INSTANT": "",
   "IOS_KIDS": "",
   "IOS_LIVE_CREATION_EXTENSION": "",
   "IOS_MESSAGES_EXTENSION": "",
   "IOS_MUSIC": "",
   "IOS_PILOT_STUDIO": "",
   "IOS_PRODUCER": "",
   "IOS_SPORTS": "",
   "IOS_TABLOID": "",
   "IOS_UNPLUGGED": "",
   "IOS_UPTIME": "",
   "IOS_WITNESS": "",
   "MUSIC_INTEGRATIONS": "",
   "MWEB": "",
   "MWEB_TIER_2": "",
   "TVANDROID": "",
   "TVAPPLE": "",
   "TVHTML5": "",
   "TVHTML5_AUDIO": "",
   "TVHTML5_CAST": "",
   "TVHTML5_FOR_KIDS": "",
   "TVHTML5_KIDS": "",
   "TVHTML5_SIMPLY": "",
   "TVHTML5_SIMPLY_EMBEDDED_PLAYER": "",
   "TVHTML5_UNPLUGGED": "",
   "TVHTML5_VR": "",
   "TVHTML5_YONGLE": "",
   "TVLITE": "",
   "TV_UNPLUGGED_ANDROID": "",
   "TV_UNPLUGGED_CAST": "",
   "UNKNOWN_INTERFACE": "",
   "WEB": "",
   "WEB_CREATOR": "",
   "WEB_EMBEDDED_PLAYER": "",
   "WEB_EXPERIMENTS": "",
   "WEB_GAMING": "",
   "WEB_HEROES": "",
   "WEB_INTERNAL_ANALYTICS": "",
   "WEB_KIDS": "",
   "WEB_LIVE_STREAMING": "",
   "WEB_MUSIC": "",
   "WEB_MUSIC_ANALYTICS": "",
   "WEB_MUSIC_EMBEDDED_PLAYER": "",
   "WEB_PARENT_TOOLS": "",
   "WEB_PHONE_VERIFICATION": "",
   "WEB_REMIX": "",
   "WEB_UNPLUGGED": "",
   "WEB_UNPLUGGED_ONBOARDING": "",
   "WEB_UNPLUGGED_OPS": "",
   "WEB_UNPLUGGED_PUBLIC": "",
   "XBOX": "",
   "XBOXONEGUIDE": "",
}

const (
   phone = "googleplay/phone.json"
   tablet = "googleplay/tablet.json"
   tv = "googleplay/tv.json"
)

func appVersion(app, elem string) (string, error) {
   cache, err := os.UserCacheDir()
   if err != nil {
      return "", err
   }
   token, err := googleplay.OpenToken(cache, "googleplay/token.json")
   if err != nil {
      return "", err
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

type token struct {
   Access_Token string
}

var logLevel format.LogLevel

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
