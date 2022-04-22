package cbc

import (
   "bytes"
   "encoding/json"
   "github.com/89z/format"
   "net/http"
)

var LogLevel format.LogLevel

func login(email, password string) (*http.Response, error) {
   buf := new(bytes.Buffer)
   err := json.NewEncoder(buf).Encode(map[string]string{
      "email": email,
      "password": password,
   })
   if err != nil {
      return nil, err
   }
   req, err := http.NewRequest(
      "POST", "https://api.loginradius.com/identity/v2/auth/login", buf,
   )
   if err != nil {
      return nil, err
   }
   req.Header.Set("Content-Type", "application/json")
   req.URL.RawQuery = "apiKey=3f4beddd-2061-49b0-ae80-6f1f2ed65b37"
   LogLevel.Dump(req)
   return new(http.Transport).RoundTrip(req)
}
