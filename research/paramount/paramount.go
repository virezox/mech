package main

import (
   "io"
   "net/http"
   "net/http/httputil"
   "net/url"
   "os"
   "strings"
)

func main() {
   var req http.Request
   req.Body = io.ReadCloser(nil)
   req.Header = make(http.Header)
   req.Header["Accept"] = []string{"text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8"}
   req.Header["Accept-Encoding"] = []string{"gzip, deflate"}
   req.Header["Accept-Language"] = []string{"en-us,en;q=0.5"}
   req.Header["Connection"] = []string{"close"}
   req.Header["Content-Length"] = []string{"0"}
   req.Header["Host"] = []string{"link.theplatform.com"}
   req.Header["Sec-Fetch-Mode"] = []string{"navigate"}
   req.Header["User-Agent"] = []string{"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/97.0.4692.20 Safari/537.36"}
   req.Method = "GET"
   req.URL = new(url.URL)
   req.URL.Host = "link.theplatform.com"
   req.URL.Path = "/s/dJ5BDC/media/guid/2198311517/3htV4fvVt4Z8gDZHqlzPOGLSMgcGc_vy"
   val := make(url.Values)
   val["format"] = []string{"SMIL"}
   val["formats"] = []string{"MPEG4,M3U"}
   req.URL.RawQuery = val.Encode()
   req.URL.Scheme = "http"
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

var body = strings.NewReader("")
