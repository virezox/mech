package main

import (
   "net/http"
   "net/http/httputil"
   "net/url"
   "os"
)

func main() {
   var req http.Request
   req.Header = make(http.Header)
   req.URL = new(url.URL)
   req.URL.Host = "player.vimeo.com"
   req.URL.Path = "/video/581039021/config"
   req.URL.Scheme = "https"
   val := make(url.Values)
   val["context"] = []string{"Vimeo\\Controller\\Api\\Resources\\VideoController."}
   val["email"] = []string{"0"}
   val["force_embed"] = []string{"1"}
   val["h"] = []string{"9603038895"}
   val["s"] = []string{"7cbe9e2d4127a912c0f380bb8ad208e843b6a5ea_1642220050"}
   req.URL.RawQuery = val.Encode()
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
