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

var names = map[string]string{
   "ANDROID": "17.11.34",
   "ANDROID_CASUAL": "",
   "ANDROID_CREATOR": "22.11.100",
   "ANDROID_EMBEDDED_PLAYER": "17.11.34",
   "ANDROID_INSTANT": "",
   "ANDROID_KIDS": "7.10.3",
   "ANDROID_LITE": "3.25.54",
   "ANDROID_MUSIC": "4.70.50",
   "ANDROID_PRODUCER": "",
   "ANDROID_SPORTS": "",
   "ANDROID_TESTSUITE": "1.9",
   "ANDROID_TV": "2.16.032",
   "ANDROID_TV_KIDS": "1.15.03",
   "ANDROID_UNPLUGGED": "6.12.1",
   "ANDROID_VR": "1.28.63",
   "ANDROID_WITNESS": "",
   "CLIENTX": "",
   "GOOGLE_ASSISTANT": "0.1",
   "GOOGLE_MEDIA_ACTIONS": "0.1",
   "IOS": "17.11.34",
   "IOS_CREATOR": "22.11.100",
   "IOS_DIRECTOR": "",
   "IOS_EMBEDDED_PLAYER": "2.0",
   "IOS_INSTANT": "",
   "IOS_KIDS": "7.10.3",
   "IOS_LIVE_CREATION_EXTENSION": "17.11.34",
   "IOS_MESSAGES_EXTENSION": "17.11.34",
   "IOS_MUSIC": "4.70.50",
   "IOS_PILOT_STUDIO": "",
   "IOS_PRODUCER": "0.1",
   "IOS_SPORTS": "",
   "IOS_TABLOID": "",
   "IOS_UNPLUGGED": "6.12.1",
   "IOS_UPTIME": "1.0",
   "IOS_WITNESS": "",
   "MUSIC_INTEGRATIONS": "0.1",
   "MWEB": "2.20220405.01.00",
   "MWEB_TIER_2": "9.20220325",
   "TVANDROID": "1.0",
   "TVAPPLE": "1.0",
   "TVHTML5": "7.20220404.09.00",
   "TVHTML5_AUDIO": "",
   "TVHTML5_CAST": "1.1.426206631",
   "TVHTML5_FOR_KIDS": "7.20220325",
   "TVHTML5_KIDS": "3.20220325",
   "TVHTML5_SIMPLY": "1.0",
   "TVHTML5_SIMPLY_EMBEDDED_PLAYER": "2.0",
   "TVHTML5_UNPLUGGED": "6.12.1",
   "TVHTML5_VR": "0.1",
   "TVHTML5_YONGLE": "0.1",
   "TVLITE": "2",
   "TV_UNPLUGGED_ANDROID": "1.13.02",
   "TV_UNPLUGGED_CAST": "0.1",
   "UNKNOWN_INTERFACE": "",
   "WEB": "2.20220405.00.00",
   "WEB_CREATOR": "1.20220405.02.00",
   "WEB_EMBEDDED_PLAYER": "1.20220403.00.00",
   "WEB_EXPERIMENTS": "1",
   "WEB_HEROES": "0.1",
   "WEB_INTERNAL_ANALYTICS": "0.1",
   "WEB_KIDS": "2.20220405.00.00",
   "WEB_LIVE_STREAMING": "",
   "WEB_MUSIC": "1.0",
   "WEB_MUSIC_ANALYTICS": "0.2",
   "WEB_MUSIC_EMBEDDED_PLAYER": "",
   "WEB_PARENT_TOOLS": "1.20220330",
   "WEB_PHONE_VERIFICATION": "1.0.0",
   "WEB_REMIX": "1.20220330.01.00",
   "WEB_UNPLUGGED": "1.20220403.00.00",
   "WEB_UNPLUGGED_ONBOARDING": "0.1",
   "WEB_UNPLUGGED_OPS": "0.1",
   "WEB_UNPLUGGED_PUBLIC": "0.1",
   "XBOX": "",
   "XBOXONEGUIDE": "1.0",
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

func post(name, version string) error {
   var play player
   play.VideoID = "eZHsmb4ezEk"
   play.Context.Client.ClientName = name
   play.Context.Client.ClientVersion = version
   buf := new(bytes.Buffer)
   if err := stdjson.NewEncoder(buf).Encode(play); err != nil {
      return err
   }
   req, err := http.NewRequest(
      "POST", "https://www.youtube.com/youtubei/v1/player", buf,
   )
   if err != nil {
      return err
   }
   // AIzaSy
   req.Header.Set("X-Goog-Api-Key", "AIzaSyAO_FJ2SlqU8Q4STEHLGCilw_Y9_11qcW8")
   logLevel.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      return errorString(res.Status)
   }
   return nil
}

type errorString string

func (e errorString) Error() string {
   return string(e)
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
