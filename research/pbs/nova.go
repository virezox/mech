package pbs

import (
   "encoding/json"
   "fmt"
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

func (a Asset) Format(f fmt.State, verb rune) {
   fmt.Fprintln(f, "Object_Type:", a.Object_Type)
   fmt.Fprint(f, "Slug: ", a.Slug)
   if verb == 'a' {
      fmt.Fprint(f, "\nPlayer_Code: ", a.Player_Code)
   }
}

type NextData struct {
   Props struct {
      PageProps struct {
         IsSeries []struct {
            Episode struct {
               Assets []Asset
            }
            Slug string
         }
      }
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
   data := new(NextData)
   if err := json.NewDecoder(res.Body).Decode(data); err != nil {
      return nil, err
   }
   return data, nil
}

func (n NextData) Asset() *Asset {
   for _, episode := range n.Props.PageProps.IsSeries {
      if episode.Slug == "nova-universe-revealed-milky-way" {
         for _, asset := range episode.Episode.Assets {
            if asset.Object_Type == "full_length" {
               return &asset
            }
         }
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
