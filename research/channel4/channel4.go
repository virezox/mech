package channel4

import (
   "bytes"
   "encoding/base64"
   "encoding/hex"
   "encoding/json"
   "github.com/89z/format"
   "io"
   "net/http"
   "strings"
)

const buildInfo = "Xiaomi/nitrogen/nitrogen:10/QKQ1.190910.002" +
   "/V12.0.1.0.QEDMIXM:user/release-keys"

const licenseURL = "https://c4.eme.lp.aws.redbeemedia.com" +
   "/wvlicenceproxy-service/widevine/acquire"

const uuid = "edef8ba9-79d6-4ace-a3c8-27dcd51d21ed"

var LogLevel format.LogLevel

type Header struct {
   BuildInfo string `json:"buildInfo"`
   PSSH string `json:"pssh"`
}

func NewHeader(kid string) (*Header, error) {
   decode := func(s string) ([]byte, error) {
      s = strings.ReplaceAll(s, "-", "")
      return hex.DecodeString(s)
   }
   dUUID, err := decode(uuid)
   if err != nil {
      return nil, err
   }
   dKID, err := decode(kid)
   if err != nil {
      return nil, err
   }
   var pssh []byte
   pssh = append(pssh, 0, 0, 0, '2', 'p', 's', 's', 'h', 0, 0, 0, 0)
   pssh = append(pssh, dUUID...)
   pssh = append(pssh, 0, 0, 0, 0x12, 0x12, 0x10)
   pssh = append(pssh, dKID...)
   var head Header
   head.BuildInfo = buildInfo
   head.PSSH = base64.StdEncoding.EncodeToString(pssh)
   return &head, nil
}

func (h Header) Payload(token []byte) (*Payload, error) {
   pssh := new(bytes.Buffer)
   err := json.NewEncoder(pssh).Encode(h)
   if err != nil {
      return nil, err
   }
   req, err := http.NewRequest("POST", "http://getwvkeys.cc/pssh", pssh)
   if err != nil {
      return nil, err
   }
   req.Header.Set("Content-Type", "application/json")
   LogLevel.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      return nil, errorString(res.Status)
   }
   var pay Payload
   if err := json.Unmarshal(token, &pay); err != nil {
      return nil, err
   }
   message, err := io.ReadAll(res.Body)
   if err != nil {
      return nil, err
   }
   pay.Message = base64.StdEncoding.EncodeToString(message)
   return &pay, nil
}

type Payload struct {
   Message string `json:"message"`
   Token string `json:"token"`
   Video struct {
      Type string `json:"type"`
   } `json:"video"`
}

func (p Payload) Widevine() (*Widevine, error) {
   payload := new(bytes.Buffer)
   err := json.NewEncoder(payload).Encode(p)
   if err != nil {
      return nil, err
   }
   req, err := http.NewRequest("POST", licenseURL, payload)
   if err != nil {
      return nil, err
   }
   req.Header = http.Header{
      "Accept-Language": {"en-US,en;q=0.9"},
      "Cache-Control": {"no-cache"},
      "Connection": {"keep-alive"},
      "Content-Type": {"application/json"},
      "Origin": {"https://www.channel4.com"},
      "Pragma": {"no-cache"},
      "Referer": {"https://www.channel4.com/"},
      "Sec-Fetch-Dest": {"empty"},
      "Sec-Fetch-Mode": {"cors"},
      "Sec-Fetch-Site": {"cross-site"},
      "User-Agent": {"Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/98.0.4758.80 Safari/537.36"},
      "sec-ch-ua-mobile": {"?0"},
      "sec-ch-ua-platform": {"Windows"},
   }
   LogLevel.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      return nil, errorString(res.Status)
   }
   vine := new(Widevine)
   if err := json.NewDecoder(res.Body).Decode(vine); err != nil {
      return nil, err
   }
   return vine, nil
}

type errorString string

func (e errorString) Error() string {
   return string(e)
}

type Widevine struct {
   License string
}

func (w Widevine) decrypt(h Header) ([]byte, error) {
   buf := new(bytes.Buffer)
   err := json.NewEncoder(buf).Encode(map[string]string{
      "buildInfo": buildInfo,
      "headers": "",
      "license": licenseURL,
      "license_response": w.License,
      "pssh": h.PSSH,
   })
   if err != nil {
      return nil, err
   }
   req, err := http.NewRequest("POST", "http://getwvkeys.cc/decrypter", buf)
   if err != nil {
      return nil, err
   }
   req.Header.Set("Content-Type", "application/json")
   LogLevel.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   return io.ReadAll(res.Body)
}
