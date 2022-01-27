package main

import (
   "net/http"
   "net/url"
   "fmt"
)

func main() {
   var req http.Request
   req.Header = make(http.Header)
   req.URL = new(url.URL)
   req.URL.Host = "pandora.app.link"
   val := make(url.Values)
   req.URL.Scheme = "https"
   val["$desktop_url"] = []string{"https://www.pandora.com/artist/the-black-dog/radio-scarecrow/train-by-the-autobahn-part-1/TRddpp5JJ2hqnVV"}
   req.Header["User-Agent"] = []string{"Android Chrome"}
   req.URL.RawQuery = val.Encode()
   res, err := new(http.Transport).RoundTrip(&req)
   if err != nil {
      panic(err)
   }
   defer res.Body.Close()
   addr, err := res.Location()
   if err != nil {
      panic(err)
   }
   fmt.Println(addr.Query())
}
