package vimeo

import (
   "encoding/json"
   "github.com/89z/format"
   "net/http"
   "net/url"
   "strconv"
   "strings"
   "time"
)

var LogLevel format.LogLevel

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
   Video ConfigVideo
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

// These are segmented, but you can actually get the full videos like this:
// skyfire.vimeocdn.com/1640649881-0xc62066ffa3260c57af3d58b6b788399c3f8a52ef/
// 64a97917-f2a3-46b6-a4cc-3e55e3dd07a8/parcel/video/fb8654f4.mp4
// Its only advertised for 426x240, but it seems to work with all of them.
// Careful, URLs like above are timestamped, so they only work for a short time.
// Also, even though it says Video, audio is included too.
func (c Config) Master() (*Master, error) {
   addr, err := c.URL()
   if err != nil {
      return nil, err
   }
   req, err := http.NewRequest("GET", addr, nil)
   if err != nil {
      return nil, err
   }
   LogLevel.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   mas := new(Master)
   if err := json.NewDecoder(res.Body).Decode(mas); err != nil {
      return nil, err
   }
   return mas, nil
}

func (c Config) URL() (string, error) {
   addr, err := url.Parse(c.Request.Files.DASH.CDNs.Fastly_Skyfire.URL)
   if err != nil {
      return "", err
   }
   addr.RawQuery = ""
   return addr.String(), nil
}

type ConfigVideo struct {
   Owner struct {
      Name string
   }
   Title string
   Duration int64
}

func (c ConfigVideo) String() string {
   var buf strings.Builder
   buf.WriteString("Owner: ")
   buf.WriteString(c.Owner.Name)
   buf.WriteString("\nTitle: ")
   buf.WriteString(c.Title)
   buf.WriteString("\nDuration: ")
   buf.WriteString(c.Time().String())
   return buf.String()
}

func (c ConfigVideo) Time() time.Duration {
   return time.Duration(c.Duration) * time.Second
}

type Master struct {
   Video []MasterVideo
}

type MasterVideo struct {
   Base_URL string
   Height int
   ID string
   Mime_Type string
   Width int
}

func (m MasterVideo) URL(con *Config) (string, error) {
   addr, err := con.URL()
   if err != nil {
      return "", err
   }
   ind := strings.Index(addr, "/sep/")
   if ind == -1 {
      return "", notFound{"/sep/"}
   }
   ext, err := format.ExtensionByType(m.Mime_Type)
   if err != nil {
      return "", err
   }
   return addr[:ind] + "/parcel/video/" + m.ID + ext, nil
}

type Oembed struct {
   Title string
   Upload_Date string
   Thumbnail_URL string
}

func NewOembed(id uint64) (*Oembed, error) {
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
   embed := new(Oembed)
   if err := json.NewDecoder(res.Body).Decode(embed); err != nil {
      return nil, err
   }
   return embed, nil
}

type notFound struct {
   input string
}

func (n notFound) Error() string {
   return strconv.Quote(n.input) + " not found"
}
