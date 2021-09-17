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

func AlbumGet(id string) (*Album, error) {
   req, err := http.NewRequest("GET", Origin + "/api/album/2/info", nil)
   if err != nil {
      return nil, err
   }
   q := req.URL.Query()
   q.Set("album_id", id)
   q.Set("key", key)
   req.URL.RawQuery = q.Encode()
   alb := new(Album)
   if err := roundTrip(req, alb); err != nil {
      return nil, err
   }
   return alb, nil
}

func AlbumPost(id string) (*Album, error) {
   val := url.Values{
      "album_id": {id}, "key": {key},
   }
   req, err := http.NewRequest(
      "POST", "/api/album/2/info", strings.NewReader(val.Encode()),
   )
   if err != nil {
      return nil, err
   }
   alb := new(Album)
   if err := roundTrip(req, alb); err != nil {
      return nil, err
   }
   return alb, nil
}
