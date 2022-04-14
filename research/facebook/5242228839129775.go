package main

import (
   "io"
   "net/http"
   "net/http/httputil"
   "net/url"
   "os"
   "strings"
)

// 135408
func main() {
   body := url.Values{
      "doc_id":[]string{"5242228839129775"},
      "variables":[]string{`{"currentID":"2883317948625723"}`},
      "__spin_b":[]string{"trunk"},
      "__spin_r":[]string{"1005355623"},
      "__spin_t":[]string{"1649974522"},
   }.Encode()
   var req http.Request
   req.Body = io.NopCloser(strings.NewReader(body))
   req.Header = make(http.Header)
   req.Header["Content-Type"] = []string{"application/x-www-form-urlencoded"}
   req.Method = "POST"
   req.URL = new(url.URL)
   req.URL.Host = "www.facebook.com"
   req.URL.Path = "/api/graphql/"
   req.URL.Scheme = "https"
   res, err := new(http.Transport).RoundTrip(&req)
   if err != nil {
      panic(err)
   }
   defer res.Body.Close()
   buf, err := httputil.DumpResponse(res, true)
   if err != nil {
      panic(err)
   }
   println(len(buf))
   return
   os.Stdout.Write(buf)
}
