package main

import (
   "bytes"
   "fmt"
   "net/http"
   "net/http/httputil"
   "net/url"
   "os"
   "strconv"
)

const (
   sf_max = 143499
   sf_min = 143441
   v_max = 58
   v_min = 50
)

func main() {
   var req http.Request
   req.Header = make(http.Header)
   req.URL = new(url.URL)
   req.URL.Host = "tv.apple.com"
   req.URL.Path = "/api/uts/v3/episodes/umc.cmc.45cu44369hb2qfuwr3fihnr8e"
   req.URL.Scheme = "https"
   val := make(url.Values)
   val["caller"] = []string{"web"}
   val["locale"] = []string{"en-US"}
   val["pfm"] = []string{"web"}
   val["sf"] = []string{strconv.Itoa(sf_max)}
   val["v"] = []string{strconv.Itoa(v_max)}
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
   if bytes.Contains(buf, []byte(`"adamId"`)) {
      fmt.Println("pass")
   } else {
      fmt.Println("fail")
   }
   fmt.Println(len(buf))
}
