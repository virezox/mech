package youtube

import (
   "github.com/89z/format/json"
   "net/http"
   "net/url"
   "strings"
)

func (x Exchange) Create(name string) error {
   return json.Create(x, name)
}

func OpenExchange(name string) (*Exchange, error) {
   return json.Open[Exchange](name)
}

const (
   // YouTube on TV
   client_id =
      "861556708454-d6dlm3lh05idd8npek18k6be8ba3oc68" +
      ".apps.googleusercontent.com"
   client_secret = "SboVhoG9s0rNafixCSGGKXAT"
)

type Exchange struct {
   Access_Token string
   Error string
   Refresh_Token string
}

func (x *Exchange) Refresh() error {
   val := url.Values{
      "client_id": {client_id},
      "client_secret": {client_secret},
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

type Oauth struct {
   Device_Code string
   User_Code string
   Verification_URL string
}

func NewOauth() (*Oauth, error) {
   val := url.Values{
      "client_id": {client_id},
      "scope": {"https://www.googleapis.com/auth/youtube"},
   }
   res, err := http.PostForm("https://oauth2.googleapis.com/device/code", val)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   auth := new(Oauth)
   if err := json.NewDecoder(res.Body).Decode(auth); err != nil {
      return nil, err
   }
   return auth, nil
}

func (o Oauth) Exchange() (*Exchange, error) {
   val := url.Values{
      "client_id": {client_id},
      "client_secret": {client_secret},
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

func (o Oauth) String() string {
   var buf strings.Builder
   buf.WriteString("1. Go to\n")
   buf.WriteString(o.Verification_URL)
   buf.WriteString("\n\n2. Enter this code\n")
   buf.WriteString(o.User_Code)
   buf.WriteString("\n\n3. Press Enter to continue")
   return buf.String()
}
