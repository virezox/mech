package pandora

import (
   "encoding/hex"
   "encoding/json"
   "net/http"
   "net/url"
   "strings"
   "time"
)

type playbackInfo struct {
   Result struct {
      AudioUrlMap struct {
         HighQuality struct {
            AudioURL string
         }
      }
   }
}

type playbackInfoRequest struct {
   DeviceCode string `json:"deviceCode"`
   IncludeAudioToken bool `json:"includeAudioToken"`
   PandoraID string `json:"pandoraId"`
   SyncTime int64 `json:"syncTime"`
   UserAuthToken string `json:"userAuthToken"`
}

func (u userLogin) playbackInfo() (*playbackInfo, error) {
   rInfo := playbackInfoRequest{
      DeviceCode: "b2946eba-c594-497b-9126-49a551346f1c",
      IncludeAudioToken: true,
      PandoraID: "TR:1168891",
      SyncTime: time.Now().Unix(),
      UserAuthToken: u.Result.UserAuthToken,
   }
   dec, err := json.Marshal(rInfo)
   if err != nil {
      return nil, err
   }
   //enc, err := encrypt(audioDec)
   enc, err := encrypt(dec)
   if err != nil {
      return nil, err
   }
   req, err := http.NewRequest(
      "POST", origin + "/services/json/",
      strings.NewReader(hex.EncodeToString(enc)),
   )
   if err != nil {
      return nil, err
   }
   req.URL.RawQuery =
      "partner_id=42&method=onDemand.getAudioPlaybackInfo&user_id=" +
      u.Result.UserID + "&auth_token=" + url.QueryEscape(u.Result.UserAuthToken)
   LogLevel.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   info := new(playbackInfo)
   if err := json.NewDecoder(res.Body).Decode(info); err != nil {
      return nil, err
   }
   return info, nil
}
