package main

import (
   "bytes"
   "fmt"
   "io"
   "net/http"
   "net/url"
   "time"
)

func main() {
   time.Sleep(time.Second)
   var req http.Request
   req.Header = make(http.Header)
   req.URL = new(url.URL)
   req.URL.Host = "twitter.com"
   req.URL.Path = "/i/api/2/search/adaptive.json"
   val := make(url.Values)
   val["q"] = []string{"filter:spaces"}
   req.Header["Authorization"] = []string{"Bearer AAAAAAAAAAAAAAAAAAAAANRILgAAAAAAnNwIzUejRCOuH5E6I8xnZz4puTs%3D1Zv7ttfk8LF81IUq16cHjhLTvJu4FA33AGWWjCpTnA"}
   req.Header["X-Guest-Token"] = []string{"1496955309839220736"}
   val["count"] = []string{"1"}
   req.URL.RawQuery = val.Encode()
   req.URL.Scheme = "https"
   res, err := new(http.Transport).RoundTrip(&req)
   if err != nil {
      panic(err)
   }
   defer res.Body.Close()
   fmt.Printf("%+v\n", res)
   buf, err := io.ReadAll(res.Body)
   if err != nil {
      panic(err)
   }
   if bytes.Contains(buf, []byte(`\/i\/spaces\/`)) {
      fmt.Println("pass")
   } else {
      fmt.Println("fail")
   }
}
