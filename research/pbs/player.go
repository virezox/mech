package pbs

import (
   "fmt"
   "github.com/89z/format/json"
   "net/http"
   "net/url"
   "strconv"
   "strings"
)

type Asset struct {
   Object_Type string
   Slug string
   Player_Code string
}

func (a Asset) Format(f fmt.State, verb rune) {
   fmt.Fprintln(f, "Object_Type:", a.Object_Type)
   fmt.Fprint(f, "Slug: ", a.Slug)
   if verb == 'a' {
      fmt.Fprint(f, "\nPlayer_Code: ", a.Player_Code)
   }
}

func (a Asset) VideoBridge() (*VideoBridge, error) {
   for _, split := range strings.Split(a.Player_Code, "'") {
      if strings.Contains(split, "/partnerplayer/") {
         addr, err := url.Parse(split)
         if err != nil {
            return nil, err
         }
         addr.Scheme = "https"
         return NewVideoBridge(addr)
      }
   }
   return nil, notFound{"/partnerplayer/"}
}

type Episode struct {
   Assets []Asset
}

func (e Episode) Asset() *Asset {
   for _, asset := range e.Assets {
      if asset.Object_Type == "full_length" {
         return &asset
      }
   }
   return nil
}

type NextData struct {
   Props struct {
      PageProps struct {
         IsEpisode Episode
         IsSeries []struct {
            Episode Episode
            Slug string
         }
      }
   }
   Query struct {
      Video string
   }
}

func NewNextData(addr string) (*NextData, error) {
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
      data = new(NextData)
      sep = []byte(` id="__NEXT_DATA__" type="application/json">`)
   )
   if err := json.Decode(res.Body, sep, data); err != nil {
      return nil, err
   }
   return data, nil
}

func (n NextData) Episode() *Episode {
   if n.Props.PageProps.IsSeries == nil {
      return &n.Props.PageProps.IsEpisode
   }
   for _, episode := range n.Props.PageProps.IsSeries {
      if episode.Slug == n.Query.Video {
         return &episode.Episode
      }
   }
   return nil
}

type notFound struct {
   value string
}

func (n notFound) Error() string {
   return strconv.Quote(n.value) + " is not found"
}
