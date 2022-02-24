package soundcloud

import (
   "encoding/json"
   "net/http"
   "net/url"
)

type resolve struct {
   Kind string
}

func Tracks(addr string) ([]Track, error) {
   req, err := http.NewRequest(
      "GET", "https://api-v2.soundcloud.com/resolve", nil,
   )
   if err != nil {
      return nil, err
   }
   req.URL.RawQuery = url.Values{
      "client_id": {clientID},
      "url": {addr},
   }.Encode()
   LogLevel.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   var tra Track
   if err := json.NewDecoder(res.Body).Decode(&tra); err != nil {
      return nil, err
   }
   return []Track{tra}, nil
}
