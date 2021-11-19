package reddit

import (
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

type HLS struct{}

func (l Link) HLS() *HLS {
   return nil
}
