package pbs

import (
   "github.com/89z/format/json"
   "net/http"
   "net/url"
   "strconv"
   "strings"
)

func NewNova(addr string) (*Nova, error) {
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
   var (
      nova = new(Nova)
      sep = []byte(` id="__NEXT_DATA__" type="application/json">`)
   )
   if err := json.Decode(res.Body, sep, nova); err != nil {
      return nil, err
   }
   return nova, nil
}

type notFound struct {
   value string
}

func (n notFound) Error() string {
   return strconv.Quote(n.value) + " is not found"
}

type Asset struct {
   Object_Type string
   Slug string
   Player_Code string
}

type Nova struct {
   Props struct {
      PageProps struct {
         Data struct {
            Episodes []struct {
               Episode struct {
                  Assets []Asset
               }
               Slug string
            }
         }
      }
   }
   Query struct {
      Video string
   }
}

func (n Nova) Asset() *Asset {
   for _, episode := range n.Props.PageProps.Data.Episodes {
      if episode.Slug == n.Query.Video {
         for _, asset := range episode.Episode.Assets {
            if asset.Object_Type == "full_length" {
               return &asset
            }
         }
      }
   }
   return nil
}

func (a Asset) Widget() (*Widget, error) {
   for _, split := range strings.Split(a.Player_Code, "'") {
      if strings.Contains(split, "/partnerplayer/") {
         addr, err := url.Parse(split)
         if err != nil {
            return nil, err
         }
         addr.Scheme = "https"
         return NewWidget(addr)
      }
   }
   return nil, notFound{"/partnerplayer/"}
}
