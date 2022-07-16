package youtube

import (
   "encoding/json"
   "github.com/89z/rosso/os"
   "net/http"
   "net/url"
   "strings"
)

func (self Header) Create(name string) error {
   file, err := os.Create(name)
   if err != nil {
      return err
   }
   defer file.Close()
   return json.NewEncoder(file).Encode(self)
}

func Open_Header(name string) (*Header, error) {
   file, err := os.Open(name)
   if err != nil {
      return nil, err
   }
   defer file.Close()
   head := new(Header)
   if err := json.NewDecoder(file).Decode(head); err != nil {
      return nil, err
   }
   return head, nil
}

const (
   // YouTube on TV
   client_ID =
      "861556708454-d6dlm3lh05idd8npek18k6be8ba3oc68" +
      ".apps.googleusercontent.com"
   client_secret = "SboVhoG9s0rNafixCSGGKXAT"
)

type Header struct {
   Access_Token string
   Error string
   Refresh_Token string
}

func (self *Header) Refresh() error {
   val := url.Values{
      "client_id": {client_ID},
      "client_secret": {client_secret},
      "grant_type": {"refresh_token"},
      "refresh_token": {self.Refresh_Token},
   }
   res, err := http.PostForm("https://oauth2.googleapis.com/token", val)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   return json.NewDecoder(res.Body).Decode(self)
}

type OAuth struct {
   Device_Code string
   User_Code string
   Verification_URL string
}

func New_OAuth() (*OAuth, error) {
   val := url.Values{
      "client_id": {client_ID},
      "scope": {"https://www.googleapis.com/auth/youtube"},
   }
   res, err := http.PostForm("https://oauth2.googleapis.com/device/code", val)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   auth := new(OAuth)
   if err := json.NewDecoder(res.Body).Decode(auth); err != nil {
      return nil, err
   }
   return auth, nil
}

func (self OAuth) Header() (*Header, error) {
   val := url.Values{
      "client_id": {client_ID},
      "client_secret": {client_secret},
      "device_code": {self.Device_Code},
      "grant_type":  {"urn:ietf:params:oauth:grant-type:device_code"},
   }
   res, err := http.PostForm("https://oauth2.googleapis.com/token", val)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   head := new(Header)
   if err := json.NewDecoder(res.Body).Decode(head); err != nil {
      return nil, err
   }
   return head, nil
}

func (self OAuth) String() string {
   var b strings.Builder
   b.WriteString("1. Go to\n")
   b.WriteString(self.Verification_URL)
   b.WriteString("\n\n2. Enter this code\n")
   b.WriteString(self.User_Code)
   b.WriteString("\n\n3. Press Enter to continue")
   return b.String()
}
