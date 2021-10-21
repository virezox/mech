package main

import (
   "net/http"
   "net/http/httputil"
   "os"
)

func main() {
   req, err := http.NewRequest(
      "GET", "https://www.amazon.com/dp/B07K5214NZ", nil,
   )
   req.Header = http.Header{
      "Accept": {"*/*"},
      "User-Agent": {"Mozilla"},
   }
   a, err := httputil.DumpRequestOut(req, false)
   if err != nil {
      panic(err)
   }
   os.Stdout.Write(a)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      panic(err)
   }
   defer res.Body.Close()
   os.Stdout.ReadFrom(res.Body)
}
