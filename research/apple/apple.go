package apple

import (
   "bytes"
   "encoding/json"
   "errors"
   "github.com/89z/format"
   "net/http"
   "net/url"
   "strconv"
)

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

func (a Auth) Create(elem ...string) error {
   return format.Create(a, elem...)
}

func OpenAuth(elem ...string) (*Auth, error) {
   return format.Open[Auth](elem...)
}

type Auth struct {
   *http.Cookie
}

type Signin struct {
   *http.Cookie
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

type Config struct {
   WebBag struct {
      AppIdKey string
   }
}

var LogLevel format.LogLevel

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

const (
   sf_max = 143499
   sf_min = 143441
   v_max = 58
   v_min = 50
)

func Episodes() (*http.Response, error) {
   req := new(http.Request)
   req.Header = make(http.Header)
   req.URL = new(url.URL)
   req.URL.Host = "tv.apple.com"
   req.URL.Path = "/api/uts/v3/episodes/umc.cmc.45cu44369hb2qfuwr3fihnr8e"
   req.URL.Scheme = "https"
   val := make(url.Values)
   val["caller"] = []string{"web"}
   val["locale"] = []string{"en-US"}
   val["pfm"] = []string{"web"}
   val["sf"] = []string{strconv.Itoa(sf_max)}
   val["v"] = []string{strconv.Itoa(v_max)}
   req.URL.RawQuery = val.Encode()
   LogLevel.Dump(req)
   // "adamId"
   return new(http.Transport).RoundTrip(req)
}
