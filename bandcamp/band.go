package bandcamp

import (
   "bytes"
   "encoding/json"
   "github.com/89z/mech"
   "net/http"
   "net/url"
   "strconv"
   "strings"
)

type Band struct {
   Bandcamp_URL string
}

func (b *Band) Get(id int) error {
   req, err := http.NewRequest("GET", MobileBand, nil)
   if err != nil {
      return err
   }
   val := url.Values{
      "band_id": {
         strconv.Itoa(id),
      },
   }
   req.URL.RawQuery = val.Encode()
   res, err := mech.RoundTrip(req)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   return json.NewDecoder(res.Body).Decode(b)
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

func (b *Band) PostForm(id int) error {
   val := url.Values{
      "band_id": {
         strconv.Itoa(id),
      },
   }
   req, err := http.NewRequest(
      "POST", MobileBand, strings.NewReader(val.Encode()),
   )
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
