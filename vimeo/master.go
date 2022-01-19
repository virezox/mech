package vimeo

import (
   "path"
)

type Video struct {
   ID string
   Width int
   Height int
   Init_Segment string
   Base_URL string
}

func (v Video) URL() string {
   return v.Base_URL + "/" + v.ID + path.Ext(v.Init_Segment)
}
