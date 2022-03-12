package vimeo

import (
   "net/url"
)

type Clip struct {
   ID int64
   UnlistedHash string
}

func NewClip(address string) (*Clip, error) {
   addr, err := url.Parse(address)
   if err != nil {
      return nil, err
   }
   // player.vimeo.com/video/412573977?h=f7f2d6fcb7
   h := addr.Query().Get("h")
   return &Clip{UnlistedHash: h}, nil
   // player.vimeo.com/video/412573977?unlisted_hash=f7f2d6fcb7
   // vimeo.com/477957994/2282452868
   // vimeo.com/477957994?unlisted_hash=2282452868
   // vimeo.com/66531465
}
