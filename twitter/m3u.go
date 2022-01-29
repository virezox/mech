package twitter

import (
   "github.com/89z/format/m3u"
   "net/http"
   "path"
)

func (s Stream) Chunks() ([]m3u.Format, error) {
   req, err := http.NewRequest("GET", s.Source.Location, nil)
   if err != nil {
      return nil, err
   }
   LogLevel.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   dir, _ := path.Split(s.Source.Location)
   return m3u.Decode(res.Body, dir)
}
