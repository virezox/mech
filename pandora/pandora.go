package pandora

import (
   "bytes"
   "encoding/hex"
   "encoding/json"
   "golang.org/x/crypto/blowfish" //lint:ignore SA1019 reason
   "net/http"
   "net/url"
   "strings"
)

const (
   origin = "http://android-tuner.pandora.com"
   partnerPassword = "AC7IBG09A3DTSYM4R41UJWL07VLN8JI7"
   syncTime = 0x7FFF_FFFF
)

var blowfishKey = []byte("6#26FRL$ZWD")

type PartnerLogin struct {
   Result struct {
      PartnerAuthToken string
   }
}

func NewPartnerLogin() (*PartnerLogin, error) {
   body := map[string]string{
      "deviceModel": "android-generic",
      "password": partnerPassword,
      "username": "android",
      "version": "5",
   }
   buf := new(bytes.Buffer)
   err := json.NewEncoder(buf).Encode(body)
   if err != nil {
      return nil, err
   }
   req, err := http.NewRequest("POST", origin + "/services/json/", buf)
   if err != nil {
      return nil, err
   }
   req.URL.RawQuery = "method=auth.partnerLogin"
   LogLevel.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      return nil, errorString(res.Status)
   }
   part := new(PartnerLogin)
   if err := json.NewDecoder(res.Body).Decode(part); err != nil {
      return nil, err
   }
   return part, nil
}

type playbackInfoRequest struct {
   // this can be empty, but must be included:
   DeviceCode string `json:"deviceCode"`
   IncludeAudioToken bool `json:"includeAudioToken"`
   PandoraID string `json:"pandoraId"`
   SyncTime int `json:"syncTime"`
   UserAuthToken string `json:"userAuthToken"`
}

type errorString string

func (e errorString) Error() string {
   return string(e)
}

type userLoginRequest struct {
   LoginType string `json:"loginType"`
   PartnerAuthToken string `json:"partnerAuthToken"`
   Password string `json:"password"`
   SyncTime int `json:"syncTime"`
   Username string `json:"username"`
}

type valueExchangeRequest struct {
   OfferName string `json:"offerName"`
   SyncTime int `json:"syncTime"`
   UserAuthToken string `json:"userAuthToken"`
}

type Cipher []byte

func (c Cipher) Encrypt() (Cipher, error) {
   block, err := blowfish.NewCipher(blowfishKey)
   if err != nil {
      return nil, err
   }
   for low := 0; low < len(c); low += blowfish.BlockSize {
      block.Encrypt(c[low:], c[low:])
   }
   return c, nil
}

func (c Cipher) Pad() Cipher {
   cLen := blowfish.BlockSize - len(c) % blowfish.BlockSize
   for high := byte(cLen); cLen >= 1; cLen-- {
      c = append(c, high)
   }
   return c
}

func (p PartnerLogin) UserLogin(username, password string) (*UserLogin, error) {
   buf, err := json.Marshal(userLoginRequest{
      LoginType: "user",
      PartnerAuthToken: p.Result.PartnerAuthToken,
      Password: password,
      SyncTime: syncTime,
      Username: username,
   })
   if err != nil {
      return nil, err
   }
   buf, err = Cipher.Pad(buf).Encrypt()
   if err != nil {
      return nil, err
   }
   req, err := http.NewRequest(
      "POST", origin + "/services/json/",
      strings.NewReader(hex.EncodeToString(buf)),
   )
   if err != nil {
      return nil, err
   }
   // auth_token can be empty, but must be included:
   req.URL.RawQuery = url.Values{
      "auth_token": {""},
      "method": {"auth.userLogin"},
      "partner_id": {"42"},
   }.Encode()
   LogLevel.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   user := new(UserLogin)
   if err := json.NewDecoder(res.Body).Decode(user); err != nil {
      return nil, err
   }
   return user, nil
}

