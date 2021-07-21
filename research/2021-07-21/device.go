package main

import (
   "encoding/json"
   "fmt"
   "net/http"
   "net/url"
   "os"
)

const (
   clientID =
      "861556708454-d6dlm3lh05idd8npek18k6be8ba3oc68" +
      ".apps.googleusercontent.com"
   clientSecret = "SboVhoG9s0rNafixCSGGKXAT"
)

func main() {
   data := url.Values{
      "client_id": {clientID},
      "scope": {"https://www.googleapis.com/auth/youtube"},
   }
   res, err := http.PostForm("https://oauth2.googleapis.com/device/code", data)
   if err != nil {
      panic(err)
   }
   defer res.Body.Close()
   var auth struct {
      Device_Code string
      User_Code string
      Verification_URL string
   }
   json.NewDecoder(res.Body).Decode(&auth)
   fmt.Printf(`1. Go to
%v

2. Enter this code
%v

3. Sign in to your Google Account

4. Press Enter to continue`, auth.Verification_URL, auth.User_Code)
   fmt.Scanln()
   data = url.Values{
      "client_id": {clientID},
      "client_secret": {clientSecret},
      "code": {auth.Device_Code},
      "grant_type": {"http://oauth.net/grant_type/device/1.0"},
   }
   if res, err := http.PostForm(
      "https://oauth2.googleapis.com/token", data,
   ); err != nil {
      panic(err)
   } else {
      defer res.Body.Close()
      os.Stdout.ReadFrom(res.Body)
   }
}
