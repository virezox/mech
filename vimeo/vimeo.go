package vimeo

import (
   "encoding/json"
   "net/http"
   "strconv"
   "time"
)

const origin = "https://player.vimeo.com"

// vimeo.com/7350260
// vimeo.com/66531465
// vimeo.com/196937578
func Valid(id string) bool {
   switch len(id) {
   case 7, 8, 9:
      return true
   }
   return false
}

type Config struct {
   Request struct {
      Files struct {
         Progressive []struct {
            Height int
            URL string
         }
      }
   }
   Video struct {
      Duration Duration
      Owner struct {
         Name string
      }
      Title string
   }
}

func NewConfig(id string) (*Config, error) {
   req, err := http.NewRequest("GET", origin + "/video/" + id + "/config", nil)
   if err != nil {
      return nil, err
   }
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   con := new(Config)
   if err := json.NewDecoder(res.Body).Decode(con); err != nil {
      return nil, err
   }
   return con, nil
}

type Duration int64

func (d Duration) String() string {
   dur := time.Duration(d) * time.Second
   return dur.String()
}

type Video struct {
   Title string
   Upload_Date string
   Thumbnail_URL string
}

func NewVideo(id int) (*Video, error) {
   req, err := http.NewRequest("GET", "https://vimeo.com/api/oembed.json", nil)
   if err != nil {
      return nil, err
   }
   req.URL.RawQuery = "url=//vimeo.com/" + strconv.Itoa(id)
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
