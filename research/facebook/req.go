package main

import (
   "github.com/89z/format"
   "io"
   "net/http"
   "net/http/httputil"
   "net/url"
   "os"
   "strings"
)

func main() {
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
   return
   res, err := new(http.Transport).RoundTrip(&req)
   if err != nil {
      panic(err)
   }
   defer res.Body.Close()
   buf, err := httputil.DumpResponse(res, true)
   if err != nil {
      panic(err)
   }
   os.Stdout.Write(buf)
}
