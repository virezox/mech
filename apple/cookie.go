package apple

import (
   "bytes"
   "encoding/json"
   "errors"
   "net/http"
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
   LogLevel.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      return nil, errors.New(res.Status)
   }
   var sign Signin
   for _, cook := range res.Cookies() {
      if cook.Name == "myacinfo" {
         sign.Cookie = cook
      }
   }
   return &sign, nil
}

func (s Signin) Auth() (*Auth, error) {
   req, err := http.NewRequest(
      "POST", "https://buy.tv.apple.com/account/web/auth", nil,
   )
   if err != nil {
      return nil, err
   }
   req.AddCookie(s.Cookie)
   req.Header.Set("Origin", "https://tv.apple.com")
   LogLevel.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      return nil, errors.New(res.Status)
   }
   var auth Auth
   for _, cook := range res.Cookies() {
      if cook.Name == "media-user-token" {
         auth.Cookie = cook
      }
   }
   return &auth, nil
}

type Signin struct { // 1
   *http.Cookie
}

type Auth struct { // 2
   *http.Cookie
}
