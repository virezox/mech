package apple

import (
   "bytes"
   "encoding/base64"
   "encoding/json"
   "errors"
   "github.com/89z/mech/widevine"
   "io"
   "net/http"
   "strings"
)

func (a *Auth) License(key, client []byte, pssh string) (*License, error) {
   keyID, err := getKeyID(pssh)
   if err != nil {
      return nil, err
   }
   var lic License
   lic.Auth = a
   lic.Module, err = widevine.NewModule(key, client, keyID)
   if err != nil {
      return nil, err
   }
   lic.body.Challenge, err = lic.Marshal()
   if err != nil {
      return nil, err
   }
   lic.body.KeySystem = "com.widevine.alpha"
   lic.body.URI = pssh
   return &lic, nil
}

func getKeyID(rawKey string) ([]byte, error) {
   _, after, ok := strings.Cut(rawKey, "data:text/plain;base64,")
   if ok {
      rawKey = after
   }
   key, err := base64.StdEncoding.DecodeString(rawKey)
   if err != nil {
      return nil, err
   }
   return widevine.KeyID(key)
}

type License struct {
   *Auth
   *widevine.Module
   body struct {
      Challenge []byte `json:"challenge"`
      ExtraServerParameters ServerParameters `json:"extra-server-parameters"`
      KeySystem string `json:"key-system"`
      URI string `json:"uri"`
   }
}

func (l License) Key(env *Environment, ep *Episode) ([]byte, error) {
   asset := ep.Asset()
   l.body.ExtraServerParameters = asset.FpsKeyServerQueryParameters
   buf := new(bytes.Buffer)
   err := json.NewEncoder(buf).Encode(l.body)
   if err != nil {
      return nil, err
   }
   req, err := http.NewRequest("POST", asset.FpsKeyServerUrl, buf)
   if err != nil {
      return nil, err
   }
   req.Header = http.Header{
      "Authorization": {"Bearer " + env.Media_API.Token},
      "Content-Type": {"application/json"},
      "X-Apple-Music-User-Token": {l.Value},
   }
   LogLevel.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      return nil, errors.New(res.Status)
   }
   out, err := io.ReadAll(res.Body)
   if err != nil {
      return nil, err
   }
   keys, err := l.Unmarshal(out)
   if err != nil {
      return nil, err
   }
   return keys.Content().Key, nil
}
