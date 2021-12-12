package vimeo

import (
   "encoding/json"
   "github.com/89z/mech"
   "net/http"
   "strconv"
   "time"
)

// This should succeed if ID is passed, and fail is URL is passed.
func Parse(id string) (int64, error) {
   return strconv.ParseInt(id, 10, 64)
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

func NewConfig(id int64) (*Config, error) {
   addr := []byte("https://player.vimeo.com/video/")
   addr = strconv.AppendInt(addr, id, 10)
   addr = append(addr, "/config"...)
   req, err := http.NewRequest("GET", string(addr), nil)
   if err != nil {
      return nil, err
   }
   mech.Dump(req)
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
   mech.Dump(req)
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
