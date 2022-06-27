package amc

import (
   "bytes"
   "github.com/89z/format/http"
   "github.com/89z/format/json"
   "github.com/89z/mech"
   "github.com/89z/mech/widevine"
   "io"
   "strconv"
   "strings"
)

func (p Playback) Content(c widevine.Client) (*widevine.Content, error) {
   source := p.DASH()
   module, err := c.Key_ID()
   if err != nil {
      return nil, err
   }
   buf, err := module.Marshal()
   if err != nil {
      return nil, err
   }
   req, err := http.NewRequest(
      "POST", source.Key_Systems.Widevine.License_URL, bytes.NewReader(buf),
   )
   if err != nil {
      return nil, err
   }
   req.Header.Set("bcov-auth", p.BC_JWT)
   res, err := Client.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   buf, err = io.ReadAll(res.Body)
   if err != nil {
      return nil, err
   }
   keys, err := module.Unmarshal(buf)
   if err != nil {
      return nil, err
   }
   return keys.Content(), nil
}

type Playback struct {
   PlaybackJsonData Playback_JSON_Data
   BC_JWT string
}

func (p Playback) Base() string {
   var buf strings.Builder
   buf.WriteString(p.PlaybackJsonData.Custom_Fields.Show)
   buf.WriteByte('-')
   buf.WriteString(p.PlaybackJsonData.Custom_Fields.Season)
   buf.WriteByte('-')
   buf.WriteString(p.PlaybackJsonData.Custom_Fields.Episode)
   buf.WriteByte('-')
   buf.WriteString(p.PlaybackJsonData.Name)
   return buf.String()
}

type Playback_JSON_Data struct {
   Custom_Fields struct {
      Show string // 1
      Season string // 2
      Episode string // 3
   }
   Name string // 4
   Sources []Source
}

type Source struct {
   Key_Systems *struct {
      Widevine struct {
         License_URL string
      } `json:"com.widevine.alpha"`
   }
   Src string
   Type string
}

func (p Playback) DASH() *Source {
   for _, source := range p.PlaybackJsonData.Sources {
      if source.Type == "application/dash+xml" {
         return &source
      }
   }
   return nil
}

var Client = http.Default_Client

func Get_NID(input string) (int64, error) {
   _, nID, found := strings.Cut(input, "--")
   if found {
      input = nID
   }
   return strconv.ParseInt(input, 10, 64)
}

type playback_request struct {
   Ad_Tags struct {
      Lat int `json:"lat"`
      Mode string `json:"mode"`
      PPID int `json:"ppid"`
      Player_Height int `json:"playerHeight"`
      Player_Width int `json:"playerWidth"`
      URL string `json:"url"`
   } `json:"adtags"`
}

type playback_response struct {
   Data struct {
      PlaybackJsonData Playback_JSON_Data
   }
}

func Unauth() (*Auth, error) {
   req, err := http.NewRequest(
      "POST", "https://gw.cds.amcn.com/auth-orchestration-id/api/v1/unauth", nil,
   )
   if err != nil {
      return nil, err
   }
   req.Header = http.Header{
      "X-Amcn-Device-Id": {"!"},
      "X-Amcn-Language": {"en"},
      "X-Amcn-Network": {"amcplus"},
      "X-Amcn-Platform": {"web"},
      "X-Amcn-Tenant": {"amcn"},
   }
   res, err := Client.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   auth := new(Auth)
   if err := json.NewDecoder(res.Body).Decode(auth); err != nil {
      return nil, err
   }
   return auth, nil
}

func (a *Auth) Login(email, password string) error {
   buf := new(bytes.Buffer)
   err := json.NewEncoder(buf).Encode(map[string]string{
      "email": email,
      "password": password,
   })
   if err != nil {
      return err
   }
   req, err := http.NewRequest(
      "POST", "https://gw.cds.amcn.com/auth-orchestration-id/api/v1/login", buf,
   )
   if err != nil {
      return err
   }
   req.Header = http.Header{
      "Authorization": {"Bearer " + a.Data.Access_Token},
      "Content-Type": {"application/json"},
      "X-Amcn-Device-Ad-Id": {"!"},
      "X-Amcn-Language": {"en"},
      "X-Amcn-Network": {"amcplus"},
      "X-Amcn-Platform": {"web"},
      "X-Amcn-Service-Group-Id": {"10"},
      "X-Amcn-Service-Id": {"amcplus"},
      "X-Amcn-Tenant": {"amcn"},
      "X-Ccpa-Do-Not-Sell": {"doNotPassData"},
   }
   res, err := Client.Do(req)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   return json.NewDecoder(res.Body).Decode(a)
}

func (a *Auth) Refresh() error {
   req, err := http.NewRequest(
      "POST",
      "https://gw.cds.amcn.com/auth-orchestration-id/api/v1/refresh",
      nil,
   )
   if err != nil {
      return err
   }
   req.Header.Set("Authorization", "Bearer " + a.Data.Refresh_Token)
   res, err := Client.Do(req)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   return json.NewDecoder(res.Body).Decode(a)
}

func (a Auth) Playback(nID int64) (*Playback, error) {
   var (
      addr []byte
      p playback_request
   )
   addr = append(addr, "https://gw.cds.amcn.com/playback-id/api/v1/playback/"...)
   addr = strconv.AppendInt(addr, nID, 10)
   p.Ad_Tags.Mode = "on-demand"
   p.Ad_Tags.URL = "!"
   buf, err := mech.Encode(p)
   if err != nil {
      return nil, err
   }
   req, err := http.NewRequest("POST", string(addr), buf)
   if err != nil {
      return nil, err
   }
   req.Header = http.Header{
      "Authorization": {"Bearer " + a.Data.Access_Token},
      "Content-Type": {"application/json"},
      "X-Amcn-Device-Ad-Id": {"!"},
      "X-Amcn-Language": {"en"},
      "X-Amcn-Network": {"amcplus"},
      "X-Amcn-Platform": {"web"},
      "X-Amcn-Service-Id": {"amcplus"},
      "X-Amcn-Tenant": {"amcn"},
      "X-Ccpa-Do-Not-Sell": {"doNotPassData"},
   }
   res, err := Client.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   var (
      out playback_response
      play Playback
   )
   if err := json.NewDecoder(res.Body).Decode(&out); err != nil {
      return nil, err
   }
   play.PlaybackJsonData = out.Data.PlaybackJsonData
   play.BC_JWT = res.Header.Get("X-AMCN-BC-JWT")
   return &play, nil
}

type Auth struct {
   Data struct {
      Access_Token string
      Refresh_Token string
   }
}

func Open_Auth(name string) (*Auth, error) {
   return json.Open[Auth](name)
}

func (a Auth) Create(name string) error {
   return json.Create(a, name)
}
