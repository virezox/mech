package main

import (
   "net/http"
   "net/url"
   "fmt"
   "encoding/json"
)

func main() {
   var req http.Request
   req.Header = make(http.Header)
   req.URL = new(url.URL)
   req.URL.Host = "api.twitter.com"
   req.URL.Path = "/2/search/adaptive.json"
   val := make(url.Values)
   req.Header["Authorization"] = []string{"Bearer AAAAAAAAAAAAAAAAAAAAANRILgAAAAAAnNwIzUejRCOuH5E6I8xnZz4puTs=1Zv7ttfk8LF81IUq16cHjhLTvJu4FA33AGWWjCpTnA"}
   val["q"] = []string{"filter:spaces"}
   req.URL.RawQuery = val.Encode()
   req.URL.Scheme = "https"
   res, err := new(http.Transport).RoundTrip(&req)
   if err != nil {
      panic(err)
   }
   defer res.Body.Close()
   var search struct {
      GlobalObjects struct {
         Tweets map[int]struct{}
      }
   }
   if err := json.NewDecoder(res.Body).Decode(&search); err != nil {
      panic(err)
   }
   fmt.Println(len(search.GlobalObjects.Tweets))
}
