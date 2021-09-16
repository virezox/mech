package bc

import (
   "net/http"
   "os"
)

func album(band, album string) (*http.Response, error) {
   req, err := http.NewRequest(
      "GET", "http://bandcamp.com/api/mobile/24/tralbum_details", nil,
   )
   if err != nil {
      return nil, err
   }
   q := req.URL.Query()
   q.Set("band_id", band)
   q.Set("tralbum_type", "a")
   q.Set("tralbum_id", album)
   req.URL.RawQuery = q.Encode()
   return new(http.Transport).RoundTrip(req)
}

func track(band, track string) (*http.Response, error) {
   req, err := http.NewRequest(
      "GET", "http://bandcamp.com/api/mobile/24/tralbum_details", nil,
   )
   if err != nil {
      return nil, err
   }
   q := req.URL.Query()
   q.Set("band_id", band)
   q.Set("tralbum_type", "t")
   q.Set("tralbum_id", track)
   req.URL.RawQuery = q.Encode()
   return new(http.Transport).RoundTrip(req)
}
