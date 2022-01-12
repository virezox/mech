package youtube

import (
   "encoding/json"
   "net/http"
   "net/url"
   "os"
   "path/filepath"
   "strings"
)

const (
   // YouTube on TV
   clientID =
      "861556708454-d6dlm3lh05idd8npek18k6be8ba3oc68" +
      ".apps.googleusercontent.com"
   clientSecret = "SboVhoG9s0rNafixCSGGKXAT"
)

type Exchange struct {
   Access_Token string
   Error string
   Refresh_Token string
}

func OpenExchange(name string) (*Exchange, error) {
   file, err := os.Open(name)
   if err != nil {
      return nil, err
   }
   defer file.Close()
   exc := new(Exchange)
   if err := json.NewDecoder(file).Decode(exc); err != nil {
      return nil, err
   }
   return exc, nil
}

func (x Exchange) Create(name string) error {
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
   return enc.Encode(x)
}

func (x *Exchange) Refresh() error {
   val := url.Values{
      "client_id": {clientID},
      "client_secret": {clientSecret},
      "grant_type": {"refresh_token"},
      "refresh_token": {x.Refresh_Token},
   }
   res, err := http.PostForm("https://oauth2.googleapis.com/token", val)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   return json.NewDecoder(res.Body).Decode(x)
}

type OAuth struct {
   Device_Code string
   User_Code string
   Verification_URL string
}

func NewOAuth() (*OAuth, error) {
   val := url.Values{
      "client_id": {clientID},
      "scope": {"https://www.googleapis.com/auth/youtube"},
   }
   res, err := http.PostForm("https://oauth2.googleapis.com/device/code", val)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   oau := new(OAuth)
   if err := json.NewDecoder(res.Body).Decode(oau); err != nil {
      return nil, err
   }
   return oau, nil
}

func (o OAuth) Exchange() (*Exchange, error) {
   val := url.Values{
      "client_id": {clientID},
      "client_secret": {clientSecret},
      "device_code": {o.Device_Code},
      "grant_type":  {"urn:ietf:params:oauth:grant-type:device_code"},
   }
   res, err := http.PostForm("https://oauth2.googleapis.com/token", val)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   exc := new(Exchange)
   if err := json.NewDecoder(res.Body).Decode(exc); err != nil {
      return nil, err
   }
   return exc, nil
}

func (o OAuth) String() string {
   var str strings.Builder
   str.WriteString("1. Go to\n")
   str.WriteString(o.Verification_URL)
   str.WriteString("\n\n2. Enter this code\n")
   str.WriteString(o.User_Code)
   str.WriteString("\n\n3. Press Enter to continue")
   return str.String()
}
