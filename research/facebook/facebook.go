package facebook

import (
   "github.com/89z/format"
   "io"
   "net/http"
   "net/http/httputil"
   "net/url"
   "os"
   "strings"
)

var LogLevel format.LogLevel

type Login struct {
   Datr string
   Lsd string
}

func NewLogin() (*Login, error) {
   req, err := http.NewRequest("GET", "https://m.facebook.com/login.php", nil)
   if err != nil {
      return nil, err
   }
   LogLevel.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   type Form struct {
      Input []struct {
         Name string `xml:"name,attr"`
         Value string `xml:"value,attr"`
      } `xml:"input"`
   }
}

func getLogin() (*http.Response, error) {
   var req http.Request
   req.Header = make(http.Header)
   req.Method = "POST"
   req.URL = new(url.URL)
   req.URL.Host = "m.facebook.com"
   req.URL.Path = "/login/device-based/regular/login/"
   req.Header["Content-Type"] = []string{"application/x-www-form-urlencoded"}
   req.Header["Cookie"] = []string{"datr=MxJfYrG9o7FrP2k9iHQ2uhm9"}
   req.URL.Scheme = "https"
   body := url.Values{
      "email":[]string{email},
      "pass":[]string{password},
      "lsd":[]string{"AVoiuEvJgiA"},
   }
   req.Body = io.NopCloser(strings.NewReader(body.Encode()))
   format.LogLevel(1).Dump(&req)
   return new(http.Transport).RoundTrip(&req)
}
