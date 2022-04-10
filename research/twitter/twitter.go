package twitter

import (
   "encoding/json"
   "github.com/89z/format"
   "github.com/klaidas/go-oauth1"
   "net/http"
   "net/url"
   "strings"
)

const bearer =
   "AAAAAAAAAAAAAAAAAAAAANRILgAAAAAAnNwIzUejRCOuH5E6I8xnZz4puTs=" +
   "1Zv7ttfk8LF81IUq16cHjhLTvJu4FA33AGWWjCpTnA"

var LogLevel format.LogLevel

func (g Guest) xauth(identifier, password string) (*xauth, error) {
   body := url.Values{
      "x_auth_identifier": {identifier},
      "x_auth_password": {password},
   }.Encode()
   req, err := http.NewRequest(
      "POST", "https://api.twitter.com/auth/1/xauth_password.json",
      strings.NewReader(body),
   )
   if err != nil {
      return nil, err
   }
   req.Header = http.Header{
      "Authorization": {"Bearer " + bearer},
      "Content-Type": {"application/x-www-form-urlencoded"},
      "X-Guest-Token": {g.Guest_Token},
   }
   LogLevel.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   auth := new(xauth)
   if err := json.NewDecoder(res.Body).Decode(auth); err != nil {
      return nil, err
   }
   return auth, nil
}
