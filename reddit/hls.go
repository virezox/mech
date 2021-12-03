package reddit

import (
   "github.com/89z/mech"
   "github.com/89z/parse/m3u"
   "net/http"
   "path"
)

func (l Link) HLS() ([]m3u.Format, error) {
   req, err := http.NewRequest("GET", l.Media.Reddit_Video.HLS_URL, nil)
   if err != nil {
      return nil, err
   }
   req.URL.RawQuery = ""
   mech.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   dir, _ := path.Split(l.Media.Reddit_Video.HLS_URL)
   return m3u.Decode(res.Body, dir)
}
