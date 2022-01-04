package apple

import (
   "encoding/json"
   "github.com/89z/format"
   "github.com/89z/format/net"
   "net/http"
   "net/url"
   "os"
   "strconv"
   "strings"
   "time"
)

const podcast = "\uF8FF.v1.catalog."

var Log = format.Log{Writer: os.Stdout}

type AssetURL string

func (a AssetURL) String() string {
   str := string(a)
   addr, err := url.Parse(str)
   if err != nil {
      return str
   }
   addr.RawQuery = ""
   return addr.String()
}

type Attributes struct {
   ArtistName string
   Name string
   ReleaseDateTime string
   Duration Duration `json:"durationInMilliseconds"`
   AssetURL AssetURL
}

type Audio struct {
   D []struct {
      Attributes Attributes
   }
}

func NewAudio(addr string) (*Audio, error) {
   req, err := http.NewRequest("GET", addr, nil)
   if err != nil {
      return nil, err
   }
   Log.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   for _, node := range net.ReadHTML(res.Body, "script") {
      if node.Attr["id"] == "shoebox-media-api-cache-amp-podcasts" {
         var raw map[string]json.RawMessage
         if err := json.Unmarshal(node.Data, &raw); err != nil {
            return nil, err
         }
         for key, val := range raw {
            if strings.HasPrefix(key, podcast) {
               unq, err := strconv.Unquote(string(val))
               if err != nil {
                  return nil, err
               }
               aud := new(Audio)
               if err := json.Unmarshal([]byte(unq), aud); err != nil {
                  return nil, err
               }
               return aud, nil
            }
         }
      }
   }
   return nil, notFound{podcast}
}

type Duration int64

func (d Duration) String() string {
   dur := time.Duration(d) * time.Millisecond
   return dur.String()
}

type notFound struct {
   input string
}

func (n notFound) Error() string {
   return strconv.Quote(n.input) + " not found"
}
