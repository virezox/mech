package channel4

import (
   "bytes"
   "encoding/base64"
   "encoding/json"
   "io"
   "net/http"
)

const licenseURL = "https://c4.eme.lp.aws.redbeemedia.com" +
   "/wvlicenceproxy-service/widevine/acquire"

var pssh = []byte{
   0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
   // Widevine UUID:
   0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
   0, 0, 0, 0,
   // length + KID:
   8, 0, 0, 0, 0, 0, 0, 0,
}

type Payload struct {
   Message string `json:"message"`
   Token string `json:"token"`
   Video struct {
      Type string `json:"type"`
   } `json:"video"`
}

func NewPayload(token io.Reader) (*Payload, error) {
   body := new(bytes.Buffer)
   err := json.NewEncoder(body).Encode(map[string]string{
      "buildInfo": "",
      "pssh": base64.StdEncoding.EncodeToString(pssh),
   })
   if err != nil {
      return nil, err
   }
   req, err := http.NewRequest("POST", "http://getwvkeys.cc/pssh", body)
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
   var pay Payload
   if err := json.NewDecoder(token).Decode(&pay); err != nil {
      return nil, err
   }
   message, err := io.ReadAll(res.Body)
   if err != nil {
      return nil, err
   }
   pay.Message = base64.StdEncoding.EncodeToString(message)
   return &pay, nil
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
   vine := new(Widevine)
   if err := json.NewDecoder(res.Body).Decode(vine); err != nil {
      return nil, err
   }
   return vine, nil
}

type Widevine struct {
   License string
}

func (w Widevine) Decrypt() ([]byte, error) {
   buf := new(bytes.Buffer)
   err := json.NewEncoder(buf).Encode(map[string]string{
      "buildInfo": "",
      "headers": "",
      "license": licenseURL,
      "license_response": w.License,
      "pssh": base64.StdEncoding.EncodeToString(pssh),
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

type errorString string

func (e errorString) Error() string {
   return string(e)
}
