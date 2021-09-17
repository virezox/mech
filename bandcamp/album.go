package bandcamp

import (
   "net/http"
   "net/url"
   "strings"
)

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
   q := req.URL.Query()
   q.Set("album_id", id)
   q.Set("key", key)
   req.URL.RawQuery = q.Encode()
   return roundTrip(req, a)
}

func (a *Album) Post(id string) error {
   val := url.Values{
      "album_id": {id}, "key": {key},
   }
   req, err := http.NewRequest(
      "POST", "/api/album/2/info", strings.NewReader(val.Encode()),
   )
   if err != nil {
      return err
   }
   return roundTrip(req, a)
}
