package bandcamp

import (
   "bytes"
   "encoding/json"
   "github.com/89z/mech"
   "net/http"
   "net/url"
   "strings"
)

type Band struct {
   Band_ID json.Number
   URL string
}

func (b *Band) Get(id string) error {
   req, err := http.NewRequest("GET", Origin + "/api/band/3/info", nil)
   if err != nil {
      return err
   }
   q := req.URL.Query()
   q.Set("band_id", id)
   q.Set("key", key)
   req.URL.RawQuery = q.Encode()
   res, err := mech.RoundTrip(req)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   return json.NewDecoder(res.Body).Decode(b)
}

func (b *Band) Post(id string) error {
   bReq := bandRequest{
      json.Number(id), key,
   }
   buf := new(bytes.Buffer)
   if err := json.NewEncoder(buf).Encode(bReq); err != nil {
      return err
   }
   req, err := http.NewRequest("POST", Origin + "/api/band/3/info", buf)
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
      "band_id": {id}, "key": {key},
   }
   req, err := http.NewRequest(
      "POST", Origin + "/api/band/3/info", strings.NewReader(val.Encode()),
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

type Discography struct {
   Discography []struct {
      URL string
   }
}

func (d *Discography) PostForm(id string) error {
   val := url.Values{
      "band_id": {id}, "key": {key},
   }
   req, err := http.NewRequest(
      "POST", Origin + "/api/band/3/discography",
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
   return json.NewDecoder(res.Body).Decode(d)
}

type bandRequest struct {
   Band_ID json.Number `json:"band_id"`
   Key string `json:"key"`
}
