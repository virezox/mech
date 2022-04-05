package youtube

import (
   "os"
   gp "github.com/89z/googleplay"
)

type player struct {
   VideoID string `json:"videoId"`
   Context struct {
      Client struct {
         ClientName string `json:"clientName"`
         ClientVersion string `json:"clientVersion"`
      } `json:"client"`
   } `json:"context"`
}

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
