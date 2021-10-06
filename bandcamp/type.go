package bandcamp

import (
   "encoding/json"
   "github.com/89z/mech"
   "net/http"
)

type Info struct {
   Album_ID int
   Band_ID int
   Track_ID int
}

const (
   ApiMobile = "http://bandcamp.com/api/mobile/24/tralbum_details"
   ApiUrl = "http://bandcamp.com/api/url/2/info"
)

// thrjozkaskhjastaurrtygitylpt
// throtaudvinroftignmarkreina
// ullrettkalladrhampa
const key = "veidihundr"

// URL to track_id, album_id or band_id, key
func NewInfo(addr string) (*Info, error) {
   req, err := http.NewRequest("GET", ApiUrl, nil)
   if err != nil {
      return nil, err
   }
   val := req.URL.Query()
   val.Set("key", key)
   val.Set("url", addr)
   req.URL.RawQuery = val.Encode()
   res, err := mech.RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   inf := new(Info)
   if err := json.NewDecoder(res.Body).Decode(inf); err != nil {
      return nil, err
   }
   return inf, nil
}

////////////////////////////////////////////////////////////////////////////////

type Detail struct {
   Band_ID int `json:"band_id"`
   Tralbum_ID int `json:"tralbum_id,omitempty"`
   Tralbum_Type string `json:"tralbum_type,omitempty"`
}

type Tralbum struct {
   Bandcamp_URL string
   Tracks []struct {
      Title string
   }
}
