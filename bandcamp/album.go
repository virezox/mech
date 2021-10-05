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

type Album struct {
   Bandcamp_URL string
   Tracks []struct {
      Title string
   }
}

func (a *Album) Get(id int) error {
   req, err := http.NewRequest("GET", MobileTralbum, nil)
   if err != nil {
      return err
   }
   val := url.Values{
      "band_id": {"1"},
      "tralbum_id": {
         strconv.Itoa(id),
      },
      "tralbum_type": {"a"},
   }
   req.URL.RawQuery = val.Encode()
   res, err := mech.RoundTrip(req)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   return json.NewDecoder(res.Body).Decode(a)
}

func (a *Album) Post(id int) error {
   body := map[string]string{
      "band_id": "1",
      "tralbum_id": strconv.Itoa(id),
      "tralbum_type": "a",
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
   return json.NewDecoder(res.Body).Decode(a)
}

func (a *Album) PostForm(id int) error {
   val := url.Values{
      "band_id": {"1"},
      "tralbum_id": {
         strconv.Itoa(id),
      },
      "tralbum_type": {"a"},
   }
   req, err := http.NewRequest(
      "POST", MobileTralbum, strings.NewReader(val.Encode()),
   )
   if err != nil {
      return err
   }
   res, err := mech.RoundTrip(req)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   return json.NewDecoder(res.Body).Decode(a)
}
