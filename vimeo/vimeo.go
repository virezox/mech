package vimeo

import (
   "encoding/json"
   "fmt"
   "net/http"
)

var Verbose bool

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
      Owner struct {
         Name string
      }
      Title string
   }
}

func NewConfig(id string) (*Config, error) {
   addr := "https://player.vimeo.com/video/" + id + "/config"
   if Verbose {
      fmt.Println("GET", addr)
   }
   res, err := http.Get(addr)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   cfg := new(Config)
   if err := json.NewDecoder(res.Body).Decode(cfg); err != nil {
      return nil, err
   }
   return cfg, nil
}

type Video struct {
   Title string
   Upload_Date string
   Thumbnail_URL string
}

func NewVideo(id string) (*Video, error) {
   addr := "https://vimeo.com/api/oembed.json?url=//vimeo.com/" + id
   if Verbose {
      fmt.Println("GET", addr)
   }
   res, err := http.Get(addr)
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
