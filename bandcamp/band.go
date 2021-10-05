package bandcamp

import (
   "bytes"
   "encoding/json"
   "github.com/89z/mech"
   "net/http"
)

const MobileBand = "http://bandcamp.com/api/mobile/24/band_details"

type Band struct {
   Bandcamp_URL string
}

func (b *Band) Post(id int) error {
   body := map[string]int{"band_id": id}
   buf := new(bytes.Buffer)
   if err := json.NewEncoder(buf).Encode(body); err != nil {
      return err
   }
   req, err := http.NewRequest("POST", MobileBand, buf)
   if err != nil {
      return err
   }
   res, err := mech.RoundTrip(req)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   return json.NewDecoder(res.Body).Decode(b)
}
