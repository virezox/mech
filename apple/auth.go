package apple

import (
   "bufio"
   "bytes"
   "encoding/json"
   "errors"
   "github.com/89z/format"
   "net/http"
   "os"
)

func (a Auth) Create(name string) error {
   file, err := format.Create(name)
   if err != nil {
      return err
   }
   defer file.Close()
   return a.Write(file)
}

type Auth struct {
   *http.Response
}

func OpenAuth(name string) (*Auth, error) {
   file, err := os.Open(name)
   if err != nil {
      return nil, err
   }
   defer file.Close()
   var auth Auth
   auth.Response, err = http.ReadResponse(bufio.NewReader(file), nil)
   if err != nil {
      return nil, err
   }
   return &auth, nil
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
   Log.Dump(req)
   var sign Signin
   sign.Response, err = new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   if sign.StatusCode != http.StatusOK {
      return nil, errors.New(sign.Status)
   }
   return &sign, nil
}

type Signin struct {
   *http.Response
}

func (s Signin) myAcInfo() *http.Cookie {
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
   req.AddCookie(s.myAcInfo())
   req.Header.Set("Origin", "https://tv.apple.com")
   Log.Dump(req)
   var auth Auth
   auth.Response, err = new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   if auth.StatusCode != http.StatusOK {
      return nil, errors.New(auth.Status)
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
