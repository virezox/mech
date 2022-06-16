package apple

import (
   "bytes"
   "encoding/json"
   "errors"
   "github.com/89z/format"
   "github.com/89z/format/xml"
   "github.com/89z/mech/widevine"
   "io"
   "net/http"
   "net/url"
   "strconv"
)

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
      "X-Apple-Music-User-Token": {r.auth.mediaUserToken().Value},
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

type Episode struct {
   Data struct {
      Playables map[string]struct {
         Assets Asset
      }
   }
}

func NewEpisode(contentID string) (*Episode, error) {
   req, err := http.NewRequest(
      "GET", "https://tv.apple.com/api/uts/v3/episodes/" + contentID, nil,
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
   LogLevel.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      return nil, errors.New(res.Status)
   }
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

func (l License) Content() (*widevine.Content, error) {
   keys, err := l.Unmarshal(l.body.License)
   if err != nil {
      return nil, err
   }
   return keys.Content(), nil
}

const (
   sf_max = 143499
   sf_min = 143441
   v_max = 58
   v_min = 50
)

var LogLevel format.LogLevel

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

type ServerParameters struct {
   AdamId string `json:"adamId"`
   SvcId string `json:"svcId"`
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
