package cbc

import (
   "bytes"
   "encoding/json"
   "github.com/89z/format"
   "net/http"
   "net/url"
)

type Profile struct {
   Tier string
   ClaimsToken string
}

func (o OverTheTop) Profile() (*Profile, error) {
   req.Header["Host"] = []string{"services.radio-canada.ca"}
   req.Header["Ott-Access-Token"] = []string{o.AccessToken}
   req.URL.Host = "services.radio-canada.ca"
   req.URL.Path = "/ott/cbc-api/v2/profile"
   req.URL.Scheme = "https"
   res, err := new(http.Transport).RoundTrip(&req)
   if err != nil {
      panic(err)
   }
   defer res.Body.Close()
}

type OverTheTop struct {
   AccessToken string
}

func (w WebToken) OverTheTop() (*OverTheTop, error) {
   buf := new(bytes.Buffer)
   err := json.NewEncoder(buf).Encode(map[string]string{
      "jwt": w.Signature,
   })
   req, err := http.NewRequest(
      "POST", "https://services.radio-canada.ca/ott/cbc-api/v2/token", buf,
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
   top := new(OverTheTop)
   if err := json.NewDecoder(res.Body).Decode(top); err != nil {
      return nil, err
   }
   return top, nil
}

type WebToken struct {
   Signature string
}

func (l Login) WebToken() (*WebToken, error) {
   req, err := http.NewRequest(
      "GET", "https://cloud-api.loginradius.com/sso/jwt/api/token", nil,
   )
   if err != nil {
      return nil, err
   }
   req.URL.RawQuery = url.Values{
      "access_token": {l.Access_Token},
      "apikey": {apiKey},
      "jwtapp": {"jwt"},
   }.Encode()
   LogLevel.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   web := new(WebToken)
   if err := json.NewDecoder(res.Body).Decode(web); err != nil {
      return nil, err
   }
   return web, nil
}

type Login struct {
   Access_Token string
   Expires_In string
}

const apiKey = "3f4beddd-2061-49b0-ae80-6f1f2ed65b37"

var LogLevel format.LogLevel

func NewLogin(email, password string) (*Login, error) {
   buf := new(bytes.Buffer)
   err := json.NewEncoder(buf).Encode(map[string]string{
      "email": email,
      "password": password,
   })
   if err != nil {
      return nil, err
   }
   req, err := http.NewRequest(
      "POST", "https://api.loginradius.com/identity/v2/auth/login", buf,
   )
   if err != nil {
      return nil, err
   }
   req.Header.Set("Content-Type", "application/json")
   req.URL.RawQuery = "apiKey=" + apiKey
   LogLevel.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   login := new(Login)
   if err := json.NewDecoder(res.Body).Decode(login); err != nil {
      return nil, err
   }
   return login, nil
}
