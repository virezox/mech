package apple

import (
   "bytes"
   "encoding/json"
   "github.com/89z/format"
   "net/http"
   "net/url"
   "strconv"
   "strings"
)

func (s Signin) Auth() (*http.Response, error) {
   req, err := http.NewRequest(
      "POST", "http://buy.tv.apple.com/account/web/auth",
      strings.NewReader(`{"webAuthorizationFlowContext":"tv"}`),
   )
   if err != nil {
      return nil, err
   }
   req.AddCookie(s.Cookie)
   req.Header["Accept"] = []string{"*/*"}
   req.Header["Accept-Language"] = []string{"en-US,en;q=0.5"}
   req.Header["Connection"] = []string{"keep-alive"}
   req.Header["Content-Type"] = []string{"application/json"}
   req.Header["Dnt"] = []string{"1"}
   req.Header["Host"] = []string{"buy.tv.apple.com"}
   req.Header["Origin"] = []string{"https://tv.apple.com"}
   req.Header["Referer"] = []string{"https://tv.apple.com/"}
   req.Header["User-Agent"] = []string{"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:88.0) Gecko/20100101 Firefox/88.0"}
   LogLevel.Dump(req)
   return new(http.Transport).RoundTrip(req)
}

func (s Signin) Create(elem ...string) error {
   return format.Create(s, elem...)
}

func OpenSignin(elem ...string) (*Signin, error) {
   return format.Open[Signin](elem...)
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
