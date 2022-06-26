package apple

import (
   "bytes"
   "encoding/json"
   "github.com/89z/format/http"
   "github.com/89z/format/xml"
   "github.com/89z/mech/widevine"
   "io"
   "net/url"
   "strconv"
)

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
   var sign Signin
   sign.Response, err = Client.Do(req)
   if err != nil {
      return nil, err
   }
   return &sign, nil
}

type Signin struct {
   *http.Response
}

func (s Signin) my_ac_info() *http.Cookie {
   for _, cook := range s.Cookies() {
      if cook.Name == "myacinfo" {
         return cook
      }
   }
   return nil
}

// FIXME add Open and Create. Also change underlying type?
type Auth struct {
   *http.Response
}

func (s Signin) Auth() (*Auth, error) {
   req, err := http.NewRequest(
      "POST", "https://buy.tv.apple.com/account/web/auth", nil,
   )
   if err != nil {
      return nil, err
   }
   req.AddCookie(s.my_ac_info())
   req.Header.Set("Origin", "https://tv.apple.com")
   var auth Auth
   auth.Response, err = Client.Do(req)
   if err != nil {
      return nil, err
   }
   return &auth, nil
}

func (a Auth) media_user_token() *http.Cookie {
   for _, cook := range a.Cookies() {
      if cook.Name == "media-user-token" {
         return cook
      }
   }
   return nil
}

func (a *Auth) Request(client widevine.Client) (*Request, error) {
   var (
      err error
      req Request
   )
   req.auth = a
   req.Module, err = client.PSSH()
   if err != nil {
      return nil, err
   }
   req.body.Challenge, err = req.Marshal()
   if err != nil {
      return nil, err
   }
   req.body.Key_System = "com.widevine.alpha"
   req.body.URI = client.Raw
   return &req, nil
}
const (
   sf_max = 143499
   sf_min = 143441
   v_max = 58
   v_min = 50
)

var Client = http.Default_Client

func (r Request) License(env *Environment, ep *Episode) (*License, error) {
   asset := ep.Asset()
   r.body.Server_Parameters = asset.FpsKeyServerQueryParameters
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
      "X-Apple-Music-User-Token": {r.auth.media_user_token().Value},
   }
   res, err := Client.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   lic := License{Module: r.Module}
   if err := json.NewDecoder(res.Body).Decode(&lic.body); err != nil {
      return nil, err
   }
   return &lic, nil
}

type Episode struct {
   Data struct {
      Playables map[string]struct {
         Assets Asset
      }
   }
}

func New_Episode(content_ID string) (*Episode, error) {
   req, err := http.NewRequest(
      "GET", "https://tv.apple.com/api/uts/v3/episodes/" + content_ID, nil,
   )
   if err != nil {
      return nil, err
   }
   req.URL.RawQuery = url.Values{
      "caller": {"web"},
      "locale": {"en-US"},
      "pfm": {"web"},
      "sf": {strconv.Itoa(sf_max)},
      "v": {strconv.Itoa(v_max)},
   }.Encode()
   res, err := Client.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   epi := new(Episode)
   if err := json.NewDecoder(res.Body).Decode(epi); err != nil {
      return nil, err
   }
   return epi, nil
}

func (e Episode) Asset() *Asset {
   for _, play := range e.Data.Playables {
      return &play.Assets
   }
   return nil
}

type Config struct {
   WebBag struct {
      AppIdKey string
   }
}

func New_Config() (*Config, error) {
   res, err := Client.Get("https://amp-account.tv.apple.com/account/web/config")
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

type Asset struct {
   FpsKeyServerQueryParameters Server_Parameters
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
      Server_Parameters Server_Parameters `json:"extra-server-parameters"`
      Key_System string `json:"key-system"`
      URI string `json:"uri"`
   }
}

func (l License) Content() (*widevine.Content, error) {
   keys, err := l.Unmarshal(l.body.License)
   if err != nil {
      return nil, err
   }
   return keys.Content(), nil
}

type Environment struct {
   Media_API struct {
      Token string // authorization: Bearer eyJhbGciOiJFUzI1NiIsInR5cCI6IkpXV...
   }
}

func New_Environment() (*Environment, error) {
   res, err := Client.Get("https://tv.apple.com")
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   var scan xml.Scanner
   scan.Data, err = io.ReadAll(res.Body)
   if err != nil {
      return nil, err
   }
   scan.Sep = []byte(`"web-tv-app/config/environment"`)
   scan.Scan()
   scan.Sep = []byte("<meta")
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

type Server_Parameters struct {
   Adam_ID string `json:"adamId"`
   Svc_ID string `json:"svcId"`
}
