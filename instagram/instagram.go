package instagram

import (
   "bytes"
   "encoding/json"
   "fmt"
   "net/http"
   "net/http/httputil"
   "os"
)

const (
   Origin = "https://i.instagram.com"
   userAgent = "Instagram 207.0.0.39.120 Android"
)

var Verbose bool

type Item struct {
   Media struct {
      Video_Versions []struct {
         URL string
      }
   }
}

type Login struct {
   http.Header
}

func NewLogin(username, password string) (*Login, error) {
   buf := bytes.NewBufferString("signed_body=SIGNATURE.")
   sig := map[string]string{
      "device_id": userAgent,
      "enc_password": "#PWD_INSTAGRAM:0:0:" + password,
      "username": username,
   }
   if err := json.NewEncoder(buf).Encode(sig); err != nil {
      return nil, err
   }
   req, err := http.NewRequest("POST", Origin + "/api/v1/accounts/login/", buf)
   if err != nil {
      return nil, err
   }
   req.Header = http.Header{
      "Content-Type": {"application/x-www-form-urlencoded"},
      "User-Agent": {userAgent},
   }
   if Verbose {
      dum, err := httputil.DumpRequest(req, true)
      if err != nil {
         return nil, err
      }
      os.Stdout.Write(dum)
   }
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      return nil, fmt.Errorf("status %q", res.Status)
   }
   return &Login{res.Header}, nil
}

func (l Login) Item(code string) (*Item, error) {
   req, err := http.NewRequest("GET", Origin + "/api/v1/clips/item/", nil)
   if err != nil {
      return nil, err
   }
   val := req.URL.Query()
   val.Set("clips_media_shortcode", code)
   req.URL.RawQuery = val.Encode()
   req.Header.Set("User-Agent", userAgent)
   req.Header.Set("Authorization", l.Get("Ig-Set-Authorization"))
   if Verbose {
   }
   return nil, nil
}
