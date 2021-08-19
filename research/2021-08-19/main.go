package main

import (
   "net/http"
   "net/http/httputil"
   "os"
)

func main() {
   req, err := playground("66531465")
   if err != nil {
      panic(err)
   }
   d, err := httputil.DumpRequest(req, true)
   if err != nil {
      panic(err)
   }
   os.Stdout.Write(d)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      panic(err)
   }
   defer res.Body.Close()
   f, err := os.Create("out.json")
   if err != nil {
      panic(err)
   }
   defer f.Close()
   f.ReadFrom(res.Body)
}
