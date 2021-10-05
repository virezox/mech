package bandcamp

import (
   "bytes"
   "encoding/json"
   "github.com/89z/mech"
   "net/http"
   "net/url"
   "strings"
)

const Origin = "http://bandcamp.com"

var Verbose = mech.Verbose

type Album struct {
   Bandcamp_URL string
   Tracks []struct {
      Title string
   }
}

func (a *Album) Get(id string) error {
   req, err := http.NewRequest(
      "GET", Origin + "/api/mobile/24/tralbum_details", nil,
   )
   if err != nil {
      return err
   }
   val := req.URL.Query()
   val.Set("band_id", "1")
   val.Set("tralbum_id", id)
   val.Set("tralbum_type", "a")
   req.URL.RawQuery = val.Encode()
   res, err := mech.RoundTrip(req)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   return json.NewDecoder(res.Body).Decode(a)
}

func (a *Album) Post(id string) error {
   body := map[string]string{
      "band_id": "1", "tralbum_type": "a", "tralbum_id": id,
   }
   buf := new(bytes.Buffer)
   if err := json.NewEncoder(buf).Encode(body); err != nil {
      return err
   }
   req, err := http.NewRequest(
      "POST", Origin + "/api/mobile/24/tralbum_details", buf,
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

func (a *Album) PostForm(id string) error {
   val := url.Values{
      "band_id": {"1"}, "tralbum_id": {id}, "tralbum_type": {"a"},
   }
   req, err := http.NewRequest(
      "POST", Origin + "/api/mobile/24/tralbum_details",
      strings.NewReader(val.Encode()),
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

type Band struct {
   Bandcamp_URL string
}

func (b *Band) Get(id string) error {
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
   return json.NewDecoder(res.Body).Decode(b)
}

func (b *Band) Post(id string) error {
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
   return json.NewDecoder(res.Body).Decode(b)
}

func (b *Band) PostForm(id string) error {
   val := url.Values{
      "band_id": {id},
   }
   req, err := http.NewRequest(
      "POST", Origin + "/api/mobile/24/band_details",
      strings.NewReader(val.Encode()),
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

type Track struct {
   Bandcamp_URL string
}

func (t *Track) Get(id string) error {
   req, err := http.NewRequest(
      "GET", Origin + "/api/mobile/24/tralbum_details", nil,
   )
   if err != nil {
      return err
   }
   val := req.URL.Query()
   val.Set("band_id", "1")
   val.Set("tralbum_id", id)
   val.Set("tralbum_type", "t")
   req.URL.RawQuery = val.Encode()
   res, err := mech.RoundTrip(req)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   return json.NewDecoder(res.Body).Decode(t)
}

func (t *Track) Post(id string) error {
   body := map[string]string{
      "band_id": "1",
      "tralbum_id": id,
      "tralbum_type": "t",
   }
   buf := new(bytes.Buffer)
   if err := json.NewEncoder(buf).Encode(body); err != nil {
      return err
   }
   req, err := http.NewRequest(
      "POST", Origin + "/api/mobile/24/tralbum_details", buf,
   )
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

func (t *Track) PostForm(id string) error {
   val := url.Values{
      "band_id": {"1"}, "tralbum_id": {id}, "tralbum_type": {"t"},
   }
   req, err := http.NewRequest(
      "POST", Origin + "/api/mobile/24/tralbum_details",
      strings.NewReader(val.Encode()),
   )
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
