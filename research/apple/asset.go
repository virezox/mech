package apple

import (
   "bytes"
   "encoding/base64"
   "encoding/json"
   "github.com/89z/mech/widevine"
   "net/http"
   "strings"
)

func (a Asset) Request(key, client []byte, pssh string) (*http.Request, error) {
   keyID, err := getKeyID(pssh)
   if err != nil {
      return nil, err
   }
   mod, err := widevine.NewModule(key, client, keyID)
   if err != nil {
      return nil, err
   }
   var lic licenseRequest
   lic.Challenge, err = mod.Marshal()
   if err != nil {
      return nil, err
   }
   lic.ExtraServerParameters = a.FpsKeyServerQueryParameters
   lic.KeySystem = "com.widevine.alpha"
   lic.URI = pssh
   buf := new(bytes.Buffer)
   if err := json.NewEncoder(buf).Encode(lic); err != nil {
      return nil, err
   }
   req, err := http.NewRequest("POST", a.FpsKeyServerUrl, buf)
   if err != nil {
      return nil, err
   }
   req.Header.Set("Content-Type", "application/json")
   return req, nil
}

type Asset struct {
   FpsKeyServerQueryParameters ServerParameters
   FpsKeyServerUrl string
   HlsUrl string
}

type ServerParameters struct {
   AdamId string `json:"adamId"`
   SvcId string `json:"svcId"`
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
