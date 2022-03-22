package instagram

import (
   "bytes"
   "encoding/json"
   "github.com/89z/format"
   "net/http"
   "strconv"
   "strings"
)

func OpenLogin(elem ...string) (*Login, error) {
   return format.Open[Login](elem...)
}

func (l Login) Create(elem ...string) error {
   return format.Create(l, elem...)
}

var LogLevel format.LogLevel

func Shortcode(address string) string {
   var prev string
   for _, split := range strings.Split(address, "/") {
      if prev == "p" {
         return split
      }
      prev = split
   }
   return ""
}

type Login struct {
   Authorization string
   Date string
}

func NewLogin(username, password string) (*Login, error) {
   buf := bytes.NewBufferString("signed_body=SIGNATURE.")
   sig := map[string]string{
      "device_id": "device_id",
      "enc_password": "#PWD_INSTAGRAM:0:0:" + password,
      "username": username,
   }
   if err := json.NewEncoder(buf).Encode(sig); err != nil {
      return nil, err
   }
   req, err := http.NewRequest(
      "POST", "https://i.instagram.com/api/v1/accounts/login/", buf,
   )
   if err != nil {
      return nil, err
   }
   req.Header = http.Header{
      "Content-Type": {"application/x-www-form-urlencoded"},
      "User-Agent": {Android.String()},
   }
   LogLevel.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   var login Login
   login.Authorization = res.Header.Get("Ig-Set-Authorization")
   login.Date = res.Header.Get("Date")
   return &login, nil
}

func (l Login) mediaInfo(id int64) (*http.Response, error) {
   buf := []byte("https://i.instagram.com/api/v1/media/")
   buf = strconv.AppendInt(buf, id, 10)
   buf = append(buf, "/info/"...)
   req, err := http.NewRequest("GET", string(buf), nil)
   if err != nil {
      return nil, err
   }
   req.Header = http.Header{
      "Authorization": {l.Authorization},
      "User-Agent": {Android.String()},
   }
   LogLevel.Dump(req)
   return new(http.Transport).RoundTrip(req)
}

// I noticed that even with the posts that have `video_dash_manifest`, you have
// to request with a correct User-Agent. If you use wrong agent, you will get a
// normal response, but the `video_dash_manifest` will be missing.
type UserAgent struct {
   API int64
   Brand string
   Density string
   Device string
   Instagram string
   Model string
   Platform string
   Release int64
   Resolution string
}

var Android = UserAgent{
   API: 99,
   Brand: "brand",
   Density: "density",
   Device: "device",
   Instagram: "222.0.0.15.114",
   Model: "model",
   Platform: "platform",
   Release: 9,
   Resolution: "9999x9999",
}

func (u UserAgent) String() string {
   buf := []byte("Instagram ")
   buf = append(buf, u.Instagram...)
   buf = append(buf, " Android ("...)
   buf = strconv.AppendInt(buf, u.API, 10)
   buf = append(buf, '/')
   buf = strconv.AppendInt(buf, u.Release, 10)
   buf = append(buf, "; "...)
   buf = append(buf, u.Density...)
   buf = append(buf, "; "...)
   buf = append(buf, u.Resolution...)
   buf = append(buf, "; "...)
   buf = append(buf, u.Brand...)
   buf = append(buf, "; "...)
   buf = append(buf, u.Model...)
   buf = append(buf, "; "...)
   buf = append(buf, u.Device...)
   buf = append(buf, "; "...)
   buf = append(buf, u.Platform...)
   return string(buf)
}

type errorString string

func (e errorString) Error() string {
   return string(e)
}
