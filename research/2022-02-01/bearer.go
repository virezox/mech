package main

import (
   "net/http"
   "net/http/httputil"
   "net/url"
   "os"
   "strings"
)

func main() {
   var req http.Request
   req.Header = make(http.Header)
   req.URL = new(url.URL)
   req.URL.Host = "i.instagram.com"
   req.URL.Path = "/api/v1/media/2506147657383710114/info/"
   req.URL.Scheme = "https"
   req.Header["Authorization"] = []string{"Bearer " + bearer}
   req.Header["User-Agent"] = []string{
      "Instagram 219.0.0.12.117 Android (" +
      strings.Join([]string{
         "24/7.0", // keep
         "420dpi", // keep
         "1080x1794", // keep
         "Google/google", // keep
         "Android SDK built for x86", // keep
         "generic_x86", // keep
         "ranchu", // keep
      }, "; ") +
      ")",
   }
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
