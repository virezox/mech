package facebook

import (
   "encoding/json"
   "github.com/89z/format"
   "github.com/89z/format/xml"
   "net/http"
   "net/url"
   "strconv"
   "strings"
)

type Regular struct {
   C_User *http.Cookie
   Xs *http.Cookie
}

func (r Regular) Video(id int64) (*Video, error) {
   buf := []byte("https://www.facebook.com/video/video_data?video_id=")
   buf = strconv.AppendInt(buf, id, 10)
   req, err := http.NewRequest("GET", string(buf), nil)
   if err != nil {
      return nil, err
   }
   req.AddCookie(r.C_User)
   req.AddCookie(r.Xs)
   LogLevel.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   vid := new(Video)
   if err := json.NewDecoder(res.Body).Decode(vid); err != nil {
      return nil, err
   }
   return vid, nil
}

var LogLevel format.LogLevel

type Input struct {
   Name string `xml:"name,attr"`
   Value string `xml:"value,attr"`
}

type Login struct {
   Datr *http.Cookie
   Lsd Input
}

func NewLogin() (*Login, error) {
   req, err := http.NewRequest("GET", "https://m.facebook.com/login.php", nil)
   if err != nil {
      return nil, err
   }
   LogLevel.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   sep := []byte(`<div class="t">`)
   var form struct {
      Input []Input `xml:"input"`
   }
   if err := xml.Decode(res.Body, sep, &form); err != nil {
      return nil, err
   }
   var login Login
   for _, input := range form.Input {
      if input.Name == "lsd" {
         login.Lsd = input
      }
   }
   for _, cook := range res.Cookies() {
      if cook.Name == "datr" {
         login.Datr = cook
      }
   }
   return &login, nil
}

func (l Login) Regular(email, password string) (*Regular, error) {
   body := url.Values{
      "email": {email},
      "lsd": {l.Lsd.Value},
      "pass": {password},
   }.Encode()
   req, err := http.NewRequest(
      "POST", "https://m.facebook.com/login/device-based/regular/login/",
      strings.NewReader(body),
   )
   if err != nil {
      return nil, err
   }
   req.AddCookie(l.Datr)
   req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
   LogLevel.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   var reg Regular
   for _, cook := range res.Cookies() {
      switch cook.Name {
      case "c_user":
         reg.C_User = cook
      case "xs":
         reg.Xs = cook
      }
   }
   return &reg, nil
}

type Video struct {
   Hd_Src string
   Sd_Src string
}
