package bandcamp

import (
   "bytes"
   "encoding/json"
   "github.com/89z/mech"
   "net/http"
)

func (d *Discography) Get(id string) error {
   req, err := http.NewRequest(
      "GET", Origin + "/api/mobile/24/band_details", nil,
   )
   if err != nil {
      return err
   }
   val := req.URL.Query()
   val.Set("band_id", id)
   req.URL.RawQuery = val.Encode()
   res, err := mech.RoundTrip(req)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   return json.NewDecoder(res.Body).Decode(d)
}

func (d *Discography) Post(id string) error {
   body := map[string]string{"band_id": id}
   buf := new(bytes.Buffer)
   if err := json.NewEncoder(buf).Encode(body); err != nil {
      return err
   }
   req, err := http.NewRequest(
      "POST", Origin + "/api/mobile/24/band_details", buf,
   )
   if err != nil {
      return err
   }
   res, err := mech.RoundTrip(req)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   return json.NewDecoder(res.Body).Decode(d)
}
