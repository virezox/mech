package bandcamp

import (
   "encoding/json"
   "net/http"
   "net/http/httputil"
   "net/url"
   "os"
   "strings"
)

type Album struct {
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
   return albumRT(req)
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
   return albumRT(req)
}

func albumRT(req *http.Request) (*Album, error) {
   if Verbose {
      d, err := httputil.DumpRequest(req, true)
      if err != nil {
         return nil, err
      }
      os.Stdout.Write(d)
   }
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   alb := new(Album)
   if err := json.NewDecoder(res.Body).Decode(alb); err != nil {
      return nil, err
   }
   return alb, nil
}
