package vimeo

import (
   "encoding/json"
   "github.com/89z/mech"
   "net/http"
   "net/url"
   "path"
   "strconv"
   "strings"
   "time"
)

var LogLevel mech.LogLevel

func Parse(id string) (uint64, error) {
   return strconv.ParseUint(id, 10, 64)
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

func (c Config) DASH() (string, error) {
   loc, err := url.Parse(c.Request.Files.DASH.CDNs.Fastly_Skyfire.URL)
   if err != nil {
      return "", err
   }
   loc.RawQuery = ""
   return loc.String(), nil
}

// These are segmented, but you can actually get the full videos like this:
// skyfire.vimeocdn.com/1640649881-0xc62066ffa3260c57af3d58b6b788399c3f8a52ef/
// 64a97917-f2a3-46b6-a4cc-3e55e3dd07a8/parcel/video/fb8654f4.mp4
// Its only advertised for 426x240, but it seems to work with all of them.
// Careful, URLs like above are timestamped, so they only work for a short time.
// Also, even though it says Video, audio is included too.
func (c Config) Videos() ([]Video, error) {
   loc, err := c.DASH()
   if err != nil {
      return nil, err
   }
   ind := strings.Index(loc, "/sep/")
   if ind == -1 {
      return nil, mech.NotFound{"/sep/"}
   }
   req, err := http.NewRequest("GET", loc, nil)
   if err != nil {
      return nil, err
   }
   LogLevel.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   var dash struct {
      Video []Video
   }
   if err := json.NewDecoder(res.Body).Decode(&dash); err != nil {
      return nil, err
   }
   var vids []Video
   for _, vid := range dash.Video {
      if vid.Init_Segment != "" {
         vid.Base_URL = loc[:ind] + "/parcel/video"
         vids = append(vids, vid)
      }
   }
   return vids, nil
}

type Duration int64

func (d Duration) String() string {
   dur := time.Duration(d) * time.Second
   return dur.String()
}

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
