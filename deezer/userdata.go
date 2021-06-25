package deezer

import (
   "encoding/json"
   "net/http"
)


type UserData struct {
   Results struct {
      CheckForm string
      User struct {
         Options struct {
            License_Token string
         }
      }
   }
   SID string
}

func NewUserData(arl string) (*UserData, error) {
   req, err := http.NewRequest("GET", GatewayWWW, nil)
   if err != nil {
      return nil, err
   }
   val := req.URL.Query()
   val.Set("api_version", "1.0")
   val.Set("api_token", "")
   val.Set("input", "3")
   val.Set("method", "deezer.getUserData")
   req.URL.RawQuery = val.Encode()
   req.Header.Set("Cookie", "arl=" + arl)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   user := new(UserData)
   for _, c := range res.Cookies() {
      if c.Name == "sid" {
         user.SID = c.Value
      }
   }
   json.NewDecoder(res.Body).Decode(user)
   return user, nil
}
