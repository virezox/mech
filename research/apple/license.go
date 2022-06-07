package apple

import (
   "bytes"
   "encoding/json"
   "errors"
   "github.com/89z/mech/widevine"
   "net/http"
)

func (a *Auth) Request(key, client []byte, pssh string) (*Request, error) {
   keyID, err := widevine.KeyID(pssh)
   if err != nil {
      return nil, err
   }
   var req Request
   req.Auth = a
   req.Module, err = widevine.NewModule(key, client, keyID)
   if err != nil {
      return nil, err
   }
   req.body.Challenge, err = req.Marshal()
   if err != nil {
      return nil, err
   }
   req.body.KeySystem = "com.widevine.alpha"
   req.body.URI = pssh
   return &req, nil
}

type Asset struct {
   FpsKeyServerQueryParameters ServerParameters
   FpsKeyServerUrl string
   HlsUrl string
}

func (r Request) License(env *Environment, ep *Episode) (*License, error) {
   asset := ep.Asset()
   r.body.ExtraServerParameters = asset.FpsKeyServerQueryParameters
   buf := new(bytes.Buffer)
   err := json.NewEncoder(buf).Encode(r.body)
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
      "X-Apple-Music-User-Token": {r.Value},
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
   lic := License{Module: r.Module}
   if err := json.NewDecoder(res.Body).Decode(&lic.body); err != nil {
      return nil, err
   }
   return &lic, nil
}

type Request struct {
   *Auth
   *widevine.Module
   body struct {
      Challenge []byte `json:"challenge"`
      ExtraServerParameters ServerParameters `json:"extra-server-parameters"`
      KeySystem string `json:"key-system"`
      URI string `json:"uri"`
   }
}

type License struct {
   *widevine.Module
   body struct {
      License []byte
   }
}

func (l License) Key() ([]byte, error) {
   keys, err := l.Unmarshal(l.body.License)
   if err != nil {
      return nil, err
   }
   return keys.Content().Key, nil
}
