package cbc

import (
   "bytes"
   "github.com/89z/format/json"
   "net/http"
   "net/url"
)

func (p Profile) Create(name string) error {
   return json.Create(p, name)
}

func Open_Profile(name string) (*Profile, error) {
   return json.Open[Profile](name)
}

const api_key = "3f4beddd-2061-49b0-ae80-6f1f2ed65b37"

type Login struct {
   Access_Token string
   Expires_In string
}

func New_Login(email, password string) (*Login, error) {
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
   req.URL.RawQuery = "apiKey=" + api_key
   Log.Dump(req)
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

func (l Login) Web_Token() (*Web_Token, error) {
   req, err := http.NewRequest(
      "GET", "https://cloud-api.loginradius.com/sso/jwt/api/token", nil,
   )
   if err != nil {
      return nil, err
   }
   req.URL.RawQuery = url.Values{
      "access_token": {l.Access_Token},
      "apikey": {api_key},
      "jwtapp": {"jwt"},
   }.Encode()
   Log.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   web := new(Web_Token)
   if err := json.NewDecoder(res.Body).Decode(web); err != nil {
      return nil, err
   }
   return web, nil
}

type Over_The_Top struct {
   AccessToken string
}

func (o Over_The_Top) Profile() (*Profile, error) {
   req, err := http.NewRequest(
      "GET", "https://services.radio-canada.ca/ott/cbc-api/v2/profile", nil,
   )
   if err != nil {
      return nil, err
   }
   req.Header.Set("OTT-Access-Token", o.AccessToken)
   Log.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   pro := new(Profile)
   if err := json.NewDecoder(res.Body).Decode(pro); err != nil {
      return nil, err
   }
   return pro, nil
}

type Profile struct {
   Tier string
   ClaimsToken string
}

type Web_Token struct {
   Signature string
}

func (w Web_Token) Over_The_Top() (*Over_The_Top, error) {
   buf := new(bytes.Buffer)
   err := json.NewEncoder(buf).Encode(map[string]string{
      "jwt": w.Signature,
   })
   if err != nil {
      return nil, err
   }
   req, err := http.NewRequest(
      "POST", "https://services.radio-canada.ca/ott/cbc-api/v2/token", buf,
   )
   if err != nil {
      return nil, err
   }
   Log.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   top := new(Over_The_Top)
   if err := json.NewDecoder(res.Body).Decode(top); err != nil {
      return nil, err
   }
   return top, nil
}
