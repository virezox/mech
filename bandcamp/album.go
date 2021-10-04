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

// thrjozkaskhjastaurrtygitylpt
// throtaudvinroftignmarkreina
// ullrettkalladrhampa
const key = "veidihundr"

func Verbose(v bool) {
   mech.Verbose(v)
}

type Album struct {
   Tracks []struct {
      URL string
   }
   URL string
}

func (a *Album) Get(id string) error {
   req, err := http.NewRequest("GET", Origin + "/api/album/2/info", nil)
   if err != nil {
      return err
   }
   val := req.URL.Query()
   val.Set("album_id", id)
   val.Set("key", key)
   req.URL.RawQuery = val.Encode()
   res, err := mech.RoundTrip(req)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   return json.NewDecoder(res.Body).Decode(a)
}

func (a *Album) Post(id string) error {
   aReq := albumRequest{
      json.Number(id), key,
   }
   buf := new(bytes.Buffer)
   if err := json.NewEncoder(buf).Encode(aReq); err != nil {
      return err
   }
   req, err := http.NewRequest("POST", Origin + "/api/album/2/info", buf)
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
      "album_id": {id}, "key": {key},
   }
   req, err := http.NewRequest(
      "POST", Origin + "/api/album/2/info", strings.NewReader(val.Encode()),
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

type albumRequest struct {
   Album_ID json.Number `json:"album_id"`
   Key string `json:"key"`
}
