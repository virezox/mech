package apple

import (
   "encoding/json"
   "github.com/89z/format"
   "github.com/89z/format/http"
   "github.com/89z/mech/widevine"
   "os"
)

type Auth []*http.Cookie

func (a Auth) Create(name string) error {
   file, err := format.Create(name)
   if err != nil {
      return err
   }
   defer file.Close()
   return json.NewEncoder(file).Encode(a)
}

func Open_Auth(name string) (Auth, error) {
   file, err := os.Open(name)
   if err != nil {
      return nil, err
   }
   defer file.Close()
   var auth Auth
   if err := json.NewDecoder(file).Decode(&auth); err != nil {
      return nil, err
   }
   return auth, nil
}

func (s Signin) Auth() (Auth, error) {
   req, err := http.NewRequest(
      "POST", "https://buy.tv.apple.com/account/web/auth", nil,
   )
   if err != nil {
      return nil, err
   }
   req.AddCookie(s.my_ac_info())
   req.Header.Set("Origin", "https://tv.apple.com")
   res, err := Client.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   return res.Cookies(), nil
}

func (a Auth) media_user_token() *http.Cookie {
   for _, cookie := range a {
      if cookie.Name == "media-user-token" {
         return cookie
      }
   }
   return nil
}

func (a Auth) Request(client widevine.Client) (*Request, error) {
   var (
      err error
      req Request
   )
   req.auth = a
   req.Module, err = client.PSSH()
   if err != nil {
      return nil, err
   }
   req.body.Challenge, err = req.Marshal()
   if err != nil {
      return nil, err
   }
   req.body.Key_System = "com.widevine.alpha"
   req.body.URI = client.Raw
   return &req, nil
}
