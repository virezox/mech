package bandcamp

import (
   "encoding/json"
   "github.com/89z/mech"
   "net/http"
   "net/url"
   "strings"
)

const key = "veidihundr"

type Album struct {
   Tracks []struct {
      URL string
   }
   URL string
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

type Band struct {
   Band_ID json.Number
   URL string
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
