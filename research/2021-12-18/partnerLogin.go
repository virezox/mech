package main

import (
   "encoding/json"
   "fmt"
   "github.com/89z/mech"
   "net/http"
   "strings"
)

var LogLevel mech.LogLevel

func main() {
   body := strings.NewReader(`
   {
      "deviceModel": "android-generic_x86",
      "password": "AC7IBG09A3DTSYM4R41UJWL07VLN8JI7",
      "username": "android",
      "version": "5"
   }
   `)
   req, err := http.NewRequest(
      "POST", "http://android-tuner.pandora.com/services/json/", body,
   )
   if err != nil {
      panic(err)
   }
   req.Header.Set("User-Agent", "Pandora/2110.1 Android/7.0 generic_x86")
   req.URL.RawQuery = "method=auth.partnerLogin"
   LogLevel.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      panic(err)
   }
   defer res.Body.Close()
   var part partnerLogin
   if err := json.NewDecoder(res.Body).Decode(&part); err != nil {
      panic(err)
   }
   fmt.Printf("%+v\n", part.Result)
}

type partnerLogin struct {
   Result struct {
      PartnerAuthToken string
   }
}
