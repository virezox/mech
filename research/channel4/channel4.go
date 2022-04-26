package channel4

import (
   "bytes"
   "encoding/json"
   "github.com/89z/format"
   "io"
   "net/http"
)

type errorString string

func (e errorString) Error() string {
   return string(e)
}

func NewPayload(token []byte) (*Payload, error) {
   buf := new(bytes.Buffer)
   err := json.NewEncoder(buf).Encode(header{PSSH: pssh})
   if err != nil {
      return nil, err
   }
   req, err := http.NewRequest("POST", "http://getwvkeys.cc/pssh", buf)
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
   message, err := io.ReadAll(res.Body)
   if err != nil {
      return nil, err
   }
   var pay Payload
   if err := json.Unmarshal(token, &pay); err != nil {
      return nil, err
   }
   pay.Message = message
   return &pay, nil
}

var pssh = []byte{
   0, 0, 0, '2', 'p', 's', 's', 'h', 0, 0, 0, 0,   //
   0xed, 0xef, 0x8b, 0xa9,                         // Widevine UUID
   0x79, 0xd6,                                     // Widevine UUID
   0x4a, 0xce,                                     // Widevine UUID
   0xa3, 0xc8,                                     // Widevine UUID
   0x27, 0xdc, 0xd5, 0x1d, 0x21, 0xed,             // Widevine UUID
   0, 0, 0, 0x12, 0x12, 0x10,                      //
   0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, // KID
}

type decrypter struct {
   BuildInfo string `json:"buildInfo"`
   Headers string `json:"headers"`
   License string `json:"license"`
   License_Response string `json:"license_response"`
   PSSH []byte `json:"pssh"`
}

const licenseURL = "https://c4.eme.lp.aws.redbeemedia.com" +
   "/wvlicenceproxy-service/widevine/acquire"

var LogLevel format.LogLevel

type Widevine struct {
   License string
}

type Payload struct {
   Message []byte `json:"message"`
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

type header struct {
   // this can be empty, but must be included
   BuildInfo string `json:"buildInfo"`
   PSSH []byte `json:"pssh"`
}

func (w Widevine) Decrypt() ([]byte, error) {
   buf := new(bytes.Buffer)
   err := json.NewEncoder(buf).Encode(decrypter{
      License: licenseURL,
      License_Response: w.License,
      PSSH: pssh,
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
   if res.StatusCode != http.StatusOK {
      return nil, errorString(res.Status)
   }
   return io.ReadAll(res.Body)
}
