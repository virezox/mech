package apple

import (
   "bytes"
   "encoding/json"
   "errors"
   "github.com/89z/format"
   "github.com/89z/format/xml"
   "github.com/89z/mech/widevine"
   "net/http"
   "net/url"
)

type Asset struct {
   FpsKeyServerQueryParameters ServerParameters
   FpsKeyServerUrl string
   HlsUrl string
}

type License struct {
   *widevine.Module
   body struct {
      License []byte
   }
}

type Request struct {
   *widevine.Module
   auth *Auth
   body struct {
      Challenge []byte `json:"challenge"`
      ExtraServerParameters ServerParameters `json:"extra-server-parameters"`
      KeySystem string `json:"key-system"`
      URI string `json:"uri"`
   }
}

func (a *Auth) Request(client widevine.Client) (*Request, error) {
   var (
      err error
      req Request
   )
   req.auth = a
   req.Module, err = client.Module()
   if err != nil {
      return nil, err
   }
   req.body.Challenge, err = req.Marshal()
   if err != nil {
      return nil, err
   }
   req.body.KeySystem = "com.widevine.alpha"
   req.body.URI = client.RawPSSH
   return &req, nil
}

func (l License) Content() (*widevine.Content, error) {
   keys, err := l.Unmarshal(l.body.License)
   if err != nil {
      return nil, err
   }
   return keys.Content(), nil
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
      "X-Apple-Music-User-Token": {r.auth.Value},
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

const (
   sf_max = 143499
   sf_min = 143441
   v_max = 58
   v_min = 50
)

var LogLevel format.LogLevel

type Auth struct {
   *http.Cookie
}

type Config struct {
   WebBag struct {
      AppIdKey string
   }
}

func NewConfig() (*Config, error) {
   req, err := http.NewRequest(
      "GET", "https://amp-account.tv.apple.com/account/web/config", nil,
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
   if res.StatusCode != http.StatusOK {
      return nil, errors.New(res.Status)
   }
   con := new(Config)
   if err := json.NewDecoder(res.Body).Decode(con); err != nil {
      return nil, err
   }
   return con, nil
}

func (c Config) Signin(email, password string) (*Signin, error) {
   buf := new(bytes.Buffer)
   err := json.NewEncoder(buf).Encode(map[string]string{
     "accountName": email,
     "password": password,
   })
   if err != nil {
      return nil, err
   }
   req, err := http.NewRequest(
      "POST", "https://idmsa.apple.com/appleauth/auth/signin", buf,
   )
   if err != nil {
      return nil, err
   }
   req.Header = http.Header{
      "Content-Type": {"application/json"},
      "X-Apple-Widget-Key": {c.WebBag.AppIdKey},
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
   var sign Signin
   for _, cook := range res.Cookies() {
      if cook.Name == "myacinfo" {
         sign.Cookie = cook
      }
   }
   return &sign, nil
}

type Environment struct {
   Media_API struct {
      Token string // authorization: Bearer eyJhbGciOiJFUzI1NiIsInR5cCI6IkpXV...
   }
}

func NewEnvironment() (*Environment, error) {
   req, err := http.NewRequest("GET", "https://tv.apple.com", nil)
   if err != nil {
      return nil, err
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
   scan, err := xml.NewScanner(res.Body)
   if err != nil {
      return nil, err
   }
   scan.Split = []byte(`"web-tv-app/config/environment"`)
   scan.Scan()
   scan.Split = []byte("<meta")
   var meta struct {
      Content string `xml:"content,attr"`
   }
   if err := scan.Decode(&meta); err != nil {
      return nil, err
   }
   content, err := url.PathUnescape(meta.Content)
   if err != nil {
      return nil, err
   }
   env := new(Environment)
   if err := json.Unmarshal([]byte(content), env); err != nil {
      return nil, err
   }
   return env, nil
}

type ServerParameters struct {
   AdamId string `json:"adamId"`
   SvcId string `json:"svcId"`
}

type Signin struct {
   *http.Cookie
}

func (s Signin) Auth() (*Auth, error) {
   req, err := http.NewRequest(
      "POST", "https://buy.tv.apple.com/account/web/auth", nil,
   )
   if err != nil {
      return nil, err
   }
   req.AddCookie(s.Cookie)
   req.Header.Set("Origin", "https://tv.apple.com")
   LogLevel.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      return nil, errors.New(res.Status)
   }
   var auth Auth
   for _, cook := range res.Cookies() {
      if cook.Name == "media-user-token" {
         auth.Cookie = cook
      }
   }
   return &auth, nil
}
