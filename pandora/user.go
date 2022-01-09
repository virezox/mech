package pandora

import (
   "encoding/json"
   "github.com/89z/format"
   "net/http"
   "net/url"
   "os"
   "path/filepath"
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

func OpenUserLogin(name string) (*UserLogin, error) {
   file, err := os.Open(name)
   if err != nil {
      return nil, err
   }
   defer file.Close()
   user := new(UserLogin)
   if err := json.NewDecoder(file).Decode(user); err != nil {
      return nil, err
   }
   return user, nil
}

func (u UserLogin) Create(name string) error {
   err := os.MkdirAll(filepath.Dir(name), os.ModeDir)
   if err != nil {
      return err
   }
   file, err := os.Create(name)
   if err != nil {
      return err
   }
   defer file.Close()
   enc := json.NewEncoder(file)
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
   // auth_token and user_Id can be empty, but they must be included
   req.URL.RawQuery = url.Values{
      "auth_token": {""},
      "method": {"onDemand.getAudioPlaybackInfo"},
      "partner_id": {"42"},
      "user_id": {""},
   }.Encode()
   format.Log.Dump(req)
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

// Token is good for 30 minutes.
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
   // auth_token and user_Id can be empty, but they must be included
   req.URL.RawQuery = url.Values{
      "auth_token": {""},
      "method": {"user.startValueExchange"},
      "partner_id": {"42"},
      "user_id": {""},
   }.Encode()
   format.Log.Dump(req)
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
