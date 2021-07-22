package main

import (
   "encoding/json"
   "fmt"
   "net/http"
   "net/url"
)

const (
   clientID =
      "861556708454-d6dlm3lh05idd8npek18k6be8ba3oc68" +
      ".apps.googleusercontent.com"
   clientSecret = "SboVhoG9s0rNafixCSGGKXAT"
)

type authorization struct {
   Device_Code string
   User_Code string
   Verification_URL string
}

func newAuthorization() (*authorization, error) {
   data := url.Values{
      "client_id": {clientID},
      "scope": {"https://www.googleapis.com/auth/youtube"},
   }
   res, err := http.PostForm("https://oauth2.googleapis.com/device/code", data)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   auth := new(authorization)
   if err := json.NewDecoder(res.Body).Decode(auth); err != nil {
      return nil, err
   }
   return auth, nil
}

func (a authorization) exchange() (*exchange, error) {
   data := url.Values{
      "client_id": {clientID},
      "client_secret": {clientSecret},
      "code": {a.Device_Code},
      "grant_type": {"http://oauth.net/grant_type/device/1.0"},
   }
   res, err := http.PostForm("https://oauth2.googleapis.com/token", data)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   exch := new(exchange)
   if err := json.NewDecoder(res.Body).Decode(exch); err != nil {
      return nil, err
   }
   return exch, nil
}

type exchange struct {
   Access_Token string
   Refresh_Token string
}

func (x *exchange) refresh() error {
   data := url.Values{
      "client_id": {clientID},
      "client_secret": {clientSecret},
      "grant_type": {"refresh_token"},
      "refresh_token": {x.Refresh_Token},
   }
   res, err := http.PostForm("https://oauth2.googleapis.com/token", data)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   return json.NewDecoder(res.Body).Decode(x)
}

func main() {
   a, err := newAuthorization()
   if err != nil {
      panic(err)
   }
   fmt.Printf(`1. Go to
%v

2. Enter this code
%v

3. Sign in to your Google Account

4. Press Enter to continue`, a.Verification_URL, a.User_Code)
   fmt.Scanln()
   x, err := a.exchange()
   if err != nil {
      panic(err)
   }
   fmt.Printf("%+v\n", x)
   if err := x.refresh(); err != nil {
      panic(err)
   }
   fmt.Printf("%+v\n", x)
}
