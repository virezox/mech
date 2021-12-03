package reddit

import (
   "github.com/89z/parse/m3u"
   "net/http"
   "path"
   "sort"
)

type HLS struct {
   ID int
   Resolution string
   URI string
}

func (l Link) HLS() ([]HLS, error) {
   req, err := http.NewRequest("GET", l.Media.Reddit_Video.HLS_URL, nil)
   if err != nil {
      return nil, err
   }
   req.URL.RawQuery = ""
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   prefix, _ := path.Split(l.Media.Reddit_Video.HLS_URL)
   var hlss []HLS
   for key, val := range m3u.NewPlaylist(res.Body) {
      hlss = append(hlss, HLS{
         Resolution: val["RESOLUTION"], URI: prefix + key,
      })
   }
   sort.Slice(hlss, func(a, b int) bool {
      return hlss[a].Resolution < hlss[b].Resolution
   })
   for i := range hlss {
      hlss[i].ID = i
   }
   return hlss, nil
}
