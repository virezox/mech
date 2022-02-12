package main

import (
   "net/http"
   "net/http/httputil"
   "net/url"
   "os"
   "strings"
   "time"
)

const consumerKey = "BUHsuO5U9DF42uJtc8QTZlOmnUaJmBJGuU1efURxeklbdiLn9L"

func main() {
   var req http.Request
   req.Header = make(http.Header)
   req.URL = new(url.URL)
   req.URL.Host = "api-http2.tumblr.com"
   req.URL.Path = "/v2/blog/lyssafreyguy/posts/187741823636/permalink"
   req.URL.Scheme = "https"
   var key strings.Builder
   key.WriteString(`OAuth oauth_consumer_key="`)
   key.WriteString(consumerKey)
   key.WriteByte('"')
   req.Header.Set("Authorization", key.String())
   time.Sleep(time.Second)
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
