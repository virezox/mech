package amc

import (
   "bytes"
   "encoding/json"
   "errors"
   "github.com/89z/format"
   "github.com/89z/mech"
   "io"
   "net/http"
   "os"
   "strconv"
   "strings"
)

var LogLevel format.LogLevel

func GetNID(input string) (int64, error) {
   _, nid, found := strings.Cut(input, "--")
   if found {
      input = nid
   }
   return strconv.ParseInt(input, 10, 64)
}

type playbackRequest struct {
   AdTags struct {
      Lat int `json:"lat"`
      Mode string `json:"mode"`
      PPID int `json:"ppid"`
      PlayerHeight int `json:"playerHeight"`
      PlayerWidth int `json:"playerWidth"`
      URL string `json:"url"`
   } `json:"adtags"`
}

type playbackResponse struct {
   Data struct {
      PlaybackJsonData PlaybackJsonData
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
   LogLevel.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      return nil, errors.New(res.Status)
   }
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
   LogLevel.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      return errors.New(res.Status)
   }
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
   LogLevel.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      return errors.New(res.Status)
   }
   return json.NewDecoder(res.Body).Decode(a)
}

func (a Auth) Playback(nID int64) (*Playback, error) {
   var (
      addr []byte
      src playbackRequest
   )
   addr = append(addr, "https://gw.cds.amcn.com/playback-id/api/v1/playback/"...)
   addr = strconv.AppendInt(addr, nID, 10)
   src.AdTags.Mode = "on-demand"
   src.AdTags.URL = "!"
   buf, err := mech.Encode(src)
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
   LogLevel.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      return nil, errors.New(res.Status)
   }
   var (
      dst playbackResponse
      play Playback
   )
   if err := json.NewDecoder(res.Body).Decode(&dst); err != nil {
      return nil, err
   }
   play.PlaybackJsonData = dst.Data.PlaybackJsonData
   play.BC_JWT = res.Header.Get("X-AMCN-BC-JWT")
   return &play, nil
}

type Auth struct {
   Data struct {
      Access_Token string
      Refresh_Token string
   }
}

func OpenAuth(name string) (*Auth, error) {
   file, err := os.Open(name)
   if err != nil {
      return nil, err
   }
   defer file.Close()
   auth := new(Auth)
   if err := json.NewDecoder(file).Decode(auth); err != nil {
      return nil, err
   }
   return auth, nil
}

func (a Auth) WriteTo(w io.Writer) (int64, error) {
   wr := format.Writer{W: w}
   err := json.NewEncoder(&wr).Encode(a)
   if err != nil {
      return 0, err
   }
   return int64(wr.N), nil
}
