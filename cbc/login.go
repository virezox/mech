package cbc

import (
   "bytes"
   "encoding/json"
   "github.com/89z/format"
   "net/http"
   "net/url"
)

const apiKey = "3f4beddd-2061-49b0-ae80-6f1f2ed65b37"

var LogLevel format.LogLevel

type Login struct {
   Access_Token string
   Expires_In string
}

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

func OpenLogin(elem ...string) (*Login, error) {
   return format.Open[Login](elem...)
}

func (l Login) Create(elem ...string) error {
   return format.Create(l, elem...)
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
