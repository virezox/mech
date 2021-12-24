package main

import (
   "net/http"
   "net/http/httputil"
   "os"
)

func main() {
   req, err := http.NewRequest(
      "GET", "https://api2.musical.ly/aweme/v1/aweme/detail/", nil,
   )
   if err != nil {
      panic(err)
   }
   req.URL.RawQuery = "aweme_id=7038818332270808325"
   res, err := new(http.Transport).RoundTrip(req)
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
