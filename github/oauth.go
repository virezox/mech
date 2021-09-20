package github

import (
   "encoding/json"
   "fmt"
   "net/http"
   "net/url"
   "strings"
)

const HtmlOrigin = "https://github.com"

// GitHub Android
const clientID = "3f8b8834a91f0caad392"

type Exchange struct {
   Access_Token string
   Error string
}

type OAuth struct {
   Device_Code string
   User_Code string
   Verification_URI string
}

func NewOAuth() (*OAuth, error) {
   val := url.Values{
      "client_id": {clientID},
   }
   req, err := http.NewRequest(
      "POST", HtmlOrigin + "/login/device/code",
      strings.NewReader(val.Encode()),
   )
   req.Header.Set("Accept", "application/json")
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      return nil, fmt.Errorf("status %q", res.Status)
   }
   oau := new(OAuth)
   if err := json.NewDecoder(res.Body).Decode(oau); err != nil {
      return nil, err
   }
   return oau, nil
}

func (o OAuth) Exchange() (*Exchange, error) {
   val := url.Values{
      "client_id": {clientID},
      "device_code": {o.Device_Code},
      "grant_type":  {"urn:ietf:params:oauth:grant-type:device_code"},
   }
   req, err := http.NewRequest(
      "POST", HtmlOrigin + "/login/oauth/access_token",
      strings.NewReader(val.Encode()),
   )
   req.Header.Set("Accept", "application/json")
   res, err := new(http.Transport).RoundTrip(req)
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
