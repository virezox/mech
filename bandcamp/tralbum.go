package bandcamp

import (
   "bytes"
   "encoding/json"
   "github.com/89z/mech"
   "net/http"
)

const MobileTralbum = "http://bandcamp.com/api/mobile/24/tralbum_details"

var Verbose = mech.Verbose

type Tralbum struct {
   Bandcamp_URL string
   Tracks []struct {
      Title string
   }
}

func (d Detail) Tralbum() (*Tralbum, error) {
   buf := new(bytes.Buffer)
   if err := json.NewEncoder(buf).Encode(d); err != nil {
      return nil, err
   }
   req, err := http.NewRequest("POST", MobileTralbum, buf)
   if err != nil {
      return nil, err
   }
   res, err := mech.RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   tra := new(Tralbum)
   if err := json.NewDecoder(res.Body).Decode(tra); err != nil {
      return nil, err
   }
   return tra, nil
}
