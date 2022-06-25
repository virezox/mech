package apple

import (
   "bytes"
   "encoding/json"
   "net/http"
)

// FIXME add Open and Create. Also change underlying type?
type Auth struct {
   *http.Response
}

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
   var sign Signin
   sign.Response, err = Client.Do(req)
   if err != nil {
      return nil, err
   }
   return &sign, nil
}

type Signin struct {
   *http.Response
}

func (s Signin) my_ac_info() *http.Cookie {
   for _, cook := range s.Cookies() {
      if cook.Name == "myacinfo" {
         return cook
      }
   }
   return nil
}

func (s Signin) Auth() (*Auth, error) {
   req, err := http.NewRequest(
      "POST", "https://buy.tv.apple.com/account/web/auth", nil,
   )
   if err != nil {
      return nil, err
   }
   req.AddCookie(s.my_ac_info())
   req.Header.Set("Origin", "https://tv.apple.com")
   var auth Auth
   auth.Response, err = Client.Do(req)
   if err != nil {
      return nil, err
   }
   return &auth, nil
}

func (a Auth) media_user_token() *http.Cookie {
   for _, cook := range a.Cookies() {
      if cook.Name == "media-user-token" {
         return cook
      }
   }
   return nil
}
