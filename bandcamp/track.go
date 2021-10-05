package bandcamp

import (
   "bytes"
   "encoding/json"
   "github.com/89z/mech"
   "net/http"
   "strconv"
)

const (
   MobileBand = "http://bandcamp.com/api/mobile/24/band_details"
   MobileTralbum = "http://bandcamp.com/api/mobile/24/tralbum_details"
)

var Verbose = mech.Verbose

type Track struct {
   Bandcamp_URL string
}

func (t *Track) Post(id int) error {
   body := map[string]string{
      "band_id": "1",
      "tralbum_id": strconv.Itoa(id),
      "tralbum_type": "t",
   }
   buf := new(bytes.Buffer)
   if err := json.NewEncoder(buf).Encode(body); err != nil {
      return err
   }
   req, err := http.NewRequest("POST", MobileTralbum, buf)
   if err != nil {
      return err
   }
   res, err := mech.RoundTrip(req)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   return json.NewDecoder(res.Body).Decode(t)
}
