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

type Meta struct {
   Property string `xml:"property,attr"`
   Content string `xml:"content,attr"`
}

// 1. year
// 2. image
func NewMeta(id int64) (*Meta, error) {
   req, err := http.NewRequest("GET", "https://www.facebook.com/video.php", nil)
   if err != nil {
      return nil, err
   }
   req.URL.RawQuery = "v=" + strconv.FormatInt(id, 10)
   LogLevel.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   scan, err := xml.NewScanner(res.Body)
   if err != nil {
      return nil, err
   }
   scan.Split = []byte("<meta ")
   for scan.Scan() {
      var meta Meta
      err := scan.Decode(&meta)
      if err != nil {
         return nil, err
      }
      if meta.Property == "og:title" {
         return &meta, nil
      }
   }
   return nil, notFound{"og:title"}
}

type notFound struct {
   value string
}

func (n notFound) Error() string {
   return strconv.Quote(n.value) + " is not found"
}

// facebook.com/video/video_data?video_id=309868367063220
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

type Video struct {
   Hd_Src string
   Sd_Src string
}

var LogLevel format.LogLevel

type Login struct {
   Datr *http.Cookie
   Lsd Input
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

type Regular struct {
   C_User *http.Cookie
   Xs *http.Cookie
}

type Input struct {
   Name string `xml:"name,attr"`
   Value string `xml:"value,attr"`
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
   var login Login
   for _, cook := range res.Cookies() {
      if cook.Name == "datr" {
         login.Datr = cook
      }
   }
   scan, err := xml.NewScanner(res.Body)
   if err != nil {
      return nil, err
   }
   scan.Split = []byte("<input ")
   for scan.Scan() {
      var input Input
      err := scan.Decode(&input)
      if err != nil {
         return nil, err
      }
      if input.Name == "lsd" {
         login.Lsd = input
      }
   }
   return &login, nil
}
