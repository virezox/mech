package instagram

import (
   "bytes"
   "encoding/json"
   "github.com/89z/format"
   "net/http"
   "os"
   "path/filepath"
   "strconv"
   "strings"
)

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
   return &login, nil
}

func OpenLogin(name string) (*Login, error) {
   file, err := os.Open(name)
   if err != nil {
      return nil, err
   }
   defer file.Close()
   log := new(Login)
   if err := json.NewDecoder(file).Decode(log); err != nil {
      return nil, err
   }
   return log, nil
}

func (l Login) Create(name string) error {
   err := os.MkdirAll(filepath.Dir(name), os.ModeDir)
   if err != nil {
      return err
   }
   file, err := os.Create(name)
   if err != nil {
      return err
   }
   defer file.Close()
   enc := json.NewEncoder(file)
   enc.SetIndent("", " ")
   return enc.Encode(l)
}

func (l Login) User(username string) (*User, error) {
   var buf strings.Builder
   buf.WriteString("https://www.instagram.com/")
   buf.WriteString(username)
   buf.WriteByte('/')
   req, err := http.NewRequest("GET", buf.String(), nil)
   if err != nil {
      return nil, err
   }
   req.Header.Set("User-Agent", Android.String())
   if l.Authorization != "" {
      req.Header.Set("Authorization", l.Authorization)
   }
   req.URL.RawQuery = "__a=1"
   LogLevel.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   var profile struct {
      GraphQL struct {
         User User
      }
   }
   if err := json.NewDecoder(res.Body).Decode(&profile); err != nil {
      return nil, err
   }
   return &profile.GraphQL.User, nil
}

type User struct {
   Edge_Followed_By struct {
      Count int64
   }
   Edge_Follow struct {
      Count int64
   }
   Edge_Owner_To_Timeline_Media EdgeURL
}

type EdgeURL struct {
   Edges []struct {
      Node struct {
         Display_URL string
         Video_URL string
      }
   }
}

func (u User) String() string {
   buf := []byte("Followers: ")
   buf = strconv.AppendInt(buf, u.Edge_Followed_By.Count, 10)
   buf = append(buf, "\nFollowing: "...)
   buf = strconv.AppendInt(buf, u.Edge_Follow.Count, 10)
   return string(buf)
}

func NewUser(username string) (*User, error) {
   return Login{}.User(username)
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
