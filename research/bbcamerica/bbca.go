package main

import (
   "net/http"
   "net/http/httputil"
   "net/url"
   "os"
   "time"
)

func main() {
   var req http.Request
   req.Header = make(http.Header)
   req.URL = new(url.URL)
   req.URL.Host = "gw.cds.amcn.com"
   req.URL.Scheme = "https"
   val := make(url.Values)
   val["device"] = []string{"mobile"}
   req.URL.RawQuery = val.Encode()
   req.Header["X-Amcn-Cache-Hash"] = []string{"f2d50469102e2f5aa19bceefa107147385fc26f35913874d27a75526d7c629bb"}
   //pass
   //req.URL.Path = "/content-compiler-cr/api/v1/content/amcn/bbca/type/season-episodes/id/1010621"
   //fail
   //req.URL.Path = "/content-compiler-cr/api/v1/content/amcn/bbca/type/season-episodes/id/1011072"
   //req.URL.Path = "/content-compiler-cr/api/v1/content/amcn/bbca/type/season-episodes/id/1051589"
   time.Sleep(2 * time.Second)
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
