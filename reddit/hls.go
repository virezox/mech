package reddit

import (
   "github.com/89z/mech"
   "github.com/89z/parse/m3u"
   "net/http"
   "path"
)

type Link struct {
   Media struct {
      Reddit_Video struct {
         DASH_URL string // v.redd.it/16cqbkev2ci51/DASHPlaylist.mpd
         HLS_URL string // v.redd.it/16cqbkev2ci51/HLSPlaylist.m3u8
      }
   }
   Subreddit string
   Title string
   URL string // https://v.redd.it/pjn0j2z4v6o71
}

func (l Link) HLS() (m3u.Playlist, error) {
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
   return m3u.NewPlaylist(res.Body, prefix), nil
}
