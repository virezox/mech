package pandora

import (
   "encoding/hex"
   "encoding/json"
   "fmt"
   "github.com/89z/mech"
   "io"
   "net/http"
   "strings"
)

type PlaybackInfo struct {
   Result struct {
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
   dec := fmt.Sprintf(`
   {
      "deviceCode": "",
      "includeAudioToken": true,
      "pandoraId": "%v",
      "syncTime": 2222222222,
      "userAuthToken": "%v"
   }
   `, id, u.Result.UserAuthToken)
   enc, err := Encrypt([]byte(dec))
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
   body := fmt.Sprintf(`
   {
      "offerName": "premium_access",
      "syncTime": 2222222222,
      "userAuthToken": "%v"
   }
   `, u.Result.UserAuthToken)
   enc, err := Encrypt([]byte(body))
   if err != nil {
      return err
   }
   req, err := http.NewRequest(
      "POST", origin + "/services/json/",
      strings.NewReader(hex.EncodeToString(enc)),
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
