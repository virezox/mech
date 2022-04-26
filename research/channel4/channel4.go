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

type PSSH struct {
   base64 string
}

func NewPSSH(kid string) (*PSSH, error) {
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
   var buf []byte
   buf = append(buf, 0, 0, 0, '2', 'p', 's', 's', 'h', 0, 0, 0, 0)
   buf = append(buf, dUUID...)
   buf = append(buf, 0, 0, 0, 0x12, 0x12, 0x10)
   buf = append(buf, dKID...)
   var pssh PSSH
   pssh.base64 = base64.StdEncoding.EncodeToString(buf)
   return &pssh, nil
}

func (p PSSH) post() ([]byte, error) {
   buf := new(bytes.Buffer)
   err := json.NewEncoder(buf).Encode(map[string]string{
      "buildInfo": buildInfo,
      "pssh": p.base64,
   })
   if err != nil {
      return nil, err
   }
   req, err := http.NewRequest("POST", "http://getwvkeys.cc/pssh", buf)
   if err != nil {
      return nil, err
   }
   LogLevel.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   return io.ReadAll(res.Body)
}

type Widevine struct {
   License string
}

func NewWidevine(payload string) (*Widevine, error) {
   req, err := http.NewRequest(
      "POST", licenseURL, strings.NewReader(payload),
   )
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
   vine := new(Widevine)
   if err := json.NewDecoder(res.Body).Decode(vine); err != nil {
      return nil, err
   }
   return vine, nil
}
