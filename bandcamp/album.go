package bandcamp

import (
   "bytes"
   "encoding/json"
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
   ar := albumRequest{
      json.Number(id), key,
   }
   buf := new(bytes.Buffer)
   if err := json.NewEncoder(buf).Encode(ar); err != nil {
      return err
   }
   req, err := http.NewRequest("POST", Origin + "/api/album/2/info", buf)
   if err != nil {
      return err
   }
   return roundTrip(req, a)
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
   return roundTrip(req, a)
}

type albumRequest struct {
   Album_ID json.Number `json:"album_id"`
   Key string `json:"key"`
}
