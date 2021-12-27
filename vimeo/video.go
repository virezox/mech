package vimeo

import (
   "encoding/json"
   "net/http"
   "strconv"
)

type Video struct {
   Title string
   Upload_Date string
   Thumbnail_URL string
}

func NewVideo(id uint64) (*Video, error) {
   req, err := http.NewRequest("GET", "https://vimeo.com/api/oembed.json", nil)
   if err != nil {
      return nil, err
   }
   req.URL.RawQuery = "url=//vimeo.com/" + strconv.FormatUint(id, 10)
   LogLevel.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   vid := new(Video)
   if err := json.NewDecoder(res.Body).Decode(vid); err != nil {
      return nil, err
   }
   return vid, nil
}
