package main

import (
   "encoding/json"
   "net/http"
   "os"
   "strings"
)

const body = `
{
   "context": {
      "client": {
         "clientName": "WEB",
         "clientVersion": "2.20210708.06.00",
      }
   },
   "videoId": "bO7PgQ-DtZk"
}
`

func main() {
   f, err := os.Open("secret.json")
   if err != nil {
      panic(err)
   }
   defer f.Close()
   secret := make(map[string]string)
   json.NewDecoder(f).Decode(&secret)
   req, err := http.NewRequest(
      "POST", "https://www.youtube.com/youtubei/v1/player",
      strings.NewReader(body),
   )
   if err != nil {
      panic(err)
   }
   q := req.URL.Query()
   q.Set("key", "AIzaSyAO_FJ2SlqU8Q4STEHLGCilw_Y9_11qcW8")
   req.URL.RawQuery = q.Encode()
   req.Header.Set("X-Origin","https://www.youtube.com")
   req.Header.Set("Authorization", secret["Authorization"])
   req.AddCookie(&http.Cookie{
      Name: "__Secure-3PSID", Value: secret["__Secure-3PSID"],
   })
   req.AddCookie(&http.Cookie{
      Name: "__Secure-3PAPISID", Value: secret["__Secure-3PAPISID"],
   })
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      panic(err)
   }
   defer res.Body.Close()
   os.Stdout.ReadFrom(res.Body)
}
