package channel4

import (
   "encoding/base64"
   "encoding/hex"
   "encoding/json"
   "github.com/89z/format"
   "net/http"
   "strings"
)

const uuid = "edef8ba9-79d6-4ace-a3c8-27dcd51d21ed"

const licenseURL = "https://c4.eme.lp.aws.redbeemedia.com" +
   "/wvlicenceproxy-service/widevine/acquire"

var LogLevel format.LogLevel

func newWidevine(payload string) (*widevine, error) {
   req, err := http.NewRequest(
      "POST", licenseURL, strings.NewReader(payload),
   )
   if err != nil {
      return nil, err
   }
   LogLevel.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   vine := new(widevine)
   if err := json.NewDecoder(res.Body).Decode(vine); err != nil {
      return nil, err
   }
   return vine, nil
}

type widevine struct {
   License string
}

func createPSSH(kid string) (string, error) {
   decode := func(s string) ([]byte, error) {
      s = strings.ReplaceAll(s, "-", "")
      return hex.DecodeString(s)
   }
   dUUID, err := decode(uuid)
   if err != nil {
      return "", err
   }
   dKid, err := decode(kid)
   if err != nil {
      return "", err
   }
   var buf []byte
   buf = append(buf, 0, 0, 0, '2', 'p', 's', 's', 'h', 0, 0, 0, 0)
   buf = append(buf, dUUID...)
   buf = append(buf, 0, 0, 0, 0x12, 0x12, 0x10)
   buf = append(buf, dKid...)
   return base64.StdEncoding.EncodeToString(buf), nil
}
