package facebook

import (
   "github.com/89z/format"
   "github.com/89z/format/json"
   "github.com/89z/format/xml"
   "net/http"
   "net/url"
   "strconv"
   "strings"
)

func (v Video) String() string {
   var buf strings.Builder
   buf.WriteString("Title: ")
   buf.WriteString(v.Title.Text)
   buf.WriteString("\nDate: ")
   buf.WriteString(v.Date.DateCreated)
   buf.WriteString("\nImage: ")
   buf.WriteString(v.Media.Preferred_Thumbnail.Image.URI)
   buf.WriteString("\nVideo: ")
   buf.WriteString(v.Media.Playable_URL_Quality_HD)
   return buf.String()
}

type Video struct {
   Title struct {
      Text string
   }
   Date struct {
      DateCreated string
   }
   Media struct {
      Preferred_Thumbnail struct {
         Image struct {
            URI string
         }
      }
      Playable_URL_Quality_HD string
   }
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

func NewVideo(id int64) (*Video, error) {
   req, err := http.NewRequest("GET", "https://www.facebook.com/video.php", nil)
   if err != nil {
      return nil, err
   }
   req.Header.Set("Accept", "text/html")
   req.URL.RawQuery = "v=" + strconv.FormatInt(id, 10)
   LogLevel.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   scan, err := json.NewScanner(res.Body)
   if err != nil {
      return nil, err
   }
   var vid Video
   scan.Split = []byte(`{"\u0040context"`)
   scan.Scan()
   if err := scan.Decode(&vid.Date); err != nil {
      return nil, err
   }
   scan.Split = []byte(`{"__typename"`)
   scan.Scan()
   if err := scan.Decode(&vid.Media); err != nil {
      return nil, err
   }
   scan.Split = []byte(`{"delight_ranges"`)
   scan.Scan()
   if err := scan.Decode(&vid.Title); err != nil {
      return nil, err
   }
   return &vid, nil
}
