package pandora

import (
   "encoding/json"
   "github.com/89z/mech"
   "io"
   "net/http"
   "strings"
)

type PlaybackInfo struct {
   Stat string
   Result *struct {
      AudioUrlMap struct {
         HighQuality struct {
            AudioURL string
         }
      }
   }
}

type UserLogin struct {
   Result struct {
      UserID string
      UserAuthToken string
   }
}

// This can be used to decode an existing login response.
func (u *UserLogin) Decode(src io.Reader) error {
   return json.NewDecoder(src).Decode(u)
}

func (u UserLogin) Encode(dst io.Writer) error {
   enc := json.NewEncoder(dst)
   enc.SetIndent("", " ")
   return enc.Encode(u)
}

func (u UserLogin) PlaybackInfo(id string) (*PlaybackInfo, error) {
   rInfo := playbackInfoRequest{
      IncludeAudioToken: true,
      PandoraID: id,
      SyncTime: syncTime,
      UserAuthToken: u.Result.UserAuthToken,
   }
   body, err := hexEncode(rInfo)
   if err != nil {
      return nil, err
   }
   req, err := http.NewRequest(
      "POST", origin + "/services/json/", strings.NewReader(body),
   )
   if err != nil {
      return nil, err
   }
   val := make(mech.Values)
   // this can be empty, but it must be included:
   val["auth_token"] = ""
   val["method"] = "onDemand.getAudioPlaybackInfo"
   val["partner_id"] = "42"
   // this can be empty, but it must be included:
   val["user_id"] = ""
   req.URL.RawQuery = val.Encode()
   LogLevel.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   info := new(PlaybackInfo)
   if err := json.NewDecoder(res.Body).Decode(info); err != nil {
      return nil, err
   }
   return info, nil
}

func (u UserLogin) ValueExchange() error {
   rValue := valueExchangeRequest{
      OfferName: "premium_access",
      SyncTime: syncTime,
      UserAuthToken: u.Result.UserAuthToken,
   }
   body, err := hexEncode(rValue)
   if err != nil {
      return err
   }
   req, err := http.NewRequest(
      "POST", origin + "/services/json/", strings.NewReader(body),
   )
   if err != nil {
      return err
   }
   val := make(mech.Values)
   // this can be empty, but must be included:
   val["auth_token"] = ""
   val["method"] = "user.startValueExchange"
   val["partner_id"] = "42"
   // this can be empty, but it must be included:
   val["user_id"] = ""
   req.URL.RawQuery = val.Encode()
   LogLevel.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return err
   }
   return res.Body.Close()
}

type playbackInfoRequest struct {
   // this can be empty, but must be included:
   DeviceCode string `json:"deviceCode"`
   IncludeAudioToken bool `json:"includeAudioToken"`
   PandoraID string `json:"pandoraId"`
   SyncTime int `json:"syncTime"`
   UserAuthToken string `json:"userAuthToken"`
}

type valueExchangeRequest struct {
   OfferName string `json:"offerName"`
   SyncTime int `json:"syncTime"`
   UserAuthToken string `json:"userAuthToken"`
}
