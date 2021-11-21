package reddit

import (
   "github.com/89z/mech"
   "github.com/89z/parse/m3u"
   "net/http"
   "path"
   "sort"
)

func (l Link) HLS() (SliceHLS, error) {
   req, err := http.NewRequest("GET", l.Media.Reddit_Video.HLS_URL, nil)
   if err != nil {
      return nil, err
   }
   req.URL.RawQuery = ""
   res, err := mech.RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   prefix, _ := path.Split(l.Media.Reddit_Video.HLS_URL)
   var hlss SliceHLS
   for key, val := range m3u.NewPlaylist(res.Body, prefix) {
      hlss = append(hlss, HLS{
         Resolution: val["RESOLUTION"], URI: key,
      })
   }
   hlss.Sort()
   for i := range hlss {
      hlss[i].ID = i
   }
   return hlss, nil
}

type HLS struct {
   ID int
   Resolution string
   URI string
}

type SliceHLS []HLS

func (h SliceHLS) Sort() {
   funs := []func(a, b HLS) bool{
      func(a, b HLS) bool {
         return a.Resolution < b.Resolution
      },
      func(a, b HLS) bool {
         return a.URI < b.URI
      },
   }
   sort.Slice(h, func(a, b int) bool {
      ha, hb := h[a], h[b]
      for _, fun := range funs {
         if fun(ha, hb) {
            return true
         }
         if fun(hb, ha) {
            break
         }
      }
      return false
   })
}
