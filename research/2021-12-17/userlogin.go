package pandora

import (
   "encoding/hex"
   "encoding/json"
   "net/http"
   "strings"
   "time"
)

type userLoginRequest struct {
   LoginType string `json:"loginType"`
   PartnerAuthToken string `json:"partnerAuthToken"`
   Password string `json:"password"`
   SyncTime int64 `json:"syncTime"`
   Username string `json:"username"`
}

type userLogin struct {
   Result struct {
      UserID string
      UserAuthToken string
   }
}

// For some reason the UserAuthToken being returned by this doesnt actually
// work.
func (p partnerLogin) userLogin(username, password string) (*userLogin, error) {
   rUser := userLoginRequest{
      LoginType: "user",
      PartnerAuthToken: p.Result.PartnerAuthToken,
      Password: password,
      SyncTime: time.Now().Unix(),
      Username: username,
   }
   dec, err := json.Marshal(rUser)
   if err != nil {
      return nil, err
   }
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
   req.Header.Set("User-Agent", "Pandora/2110.1 Android/7.0 generic_x86")
   req.URL.RawQuery = values{
      "auth_token": p.Result.PartnerAuthToken,
      "method": "auth.userLogin",
      "partner_id": "42",
   }.encode()
   LogLevel.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   user := new(userLogin)
   if err := json.NewDecoder(res.Body).Decode(user); err != nil {
      return nil, err
   }
   return user, nil
}
