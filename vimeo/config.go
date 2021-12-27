package vimeo

import (
   "encoding/json"
   "github.com/89z/mech"
   "net/http"
   "net/url"
   "strconv"
   "time"
)

var LogLevel mech.LogLevel

func Parse(id string) (uint64, error) {
   return strconv.ParseUint(id, 10, 64)
}

type DASH struct {
   Base_URL string
   Video []struct {
      Width int
      Height int
      Base_URL string
   }
}

func (c Config) DASH() (*DASH, error) {
   // remove query_string_ranges=1
   loc, err := url.Parse(c.Request.Files.DASH.CDNs.Fastly_Skyfire.URL)
   if err != nil {
      return nil, err
   }
   loc.RawQuery = ""
   req, err := http.NewRequest("GET", loc.String(), nil)
   if err != nil {
      return nil, err
   }
   LogLevel.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   dash := new(DASH)
   if err := json.NewDecoder(res.Body).Decode(dash); err != nil {
      return nil, err
   }
   return dash, nil
}

type Config struct {
   Request struct {
      Files struct {
         DASH struct {
            CDNs struct {
               Fastly_Skyfire struct {
                  URL string
               }
            }
         }
         Progressive []struct {
            Width int
            Height int
            URL string
         }
      }
      Timestamp int64 // this is just the current time
   }
   Video struct {
      Duration Duration
      Owner struct {
         Name string
      }
      Title string
   }
}

func NewConfig(id uint64) (*Config, error) {
   addr := []byte("https://player.vimeo.com/video/")
   addr = strconv.AppendUint(addr, id, 10)
   addr = append(addr, "/config"...)
   req, err := http.NewRequest("GET", string(addr), nil)
   if err != nil {
      return nil, err
   }
   LogLevel.Dump(req)
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
