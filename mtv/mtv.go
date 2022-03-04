package mtv

import (
   "encoding/json"
   "github.com/89z/format"
   "net/http"
   "net/url"
   "strings"
)

var LogLevel format.LogLevel

type Item struct {
   EntityType string
   ParentEntity struct {
      Title string
   }
   ShortID string
   Title string
   VideoServiceURL string
}

func NewItem(addr string) Item {
   var (
      item Item
      prev string
   )
   for _, split := range strings.Split(addr, "/") {
      switch prev {
      case "episodes":
         item.EntityType = "episode"
         item.ShortID = split
      case "video-clips":
         item.EntityType = "showvideo"
         item.ShortID = split
      }
      prev = split
   }
   return item
}

func (i Item) Property() (*Property, error) {
   req, err := http.NewRequest(
      "GET", "https://neutron-api.viacom.tech/api/2.9/property", nil,
   )
   if err != nil {
      return nil, err
   }
   req.URL.RawQuery = url.Values{
      "brand": {"mtv"},
      "platform": {"web"},
      "region": {"US-PHASE1"},
      "shortId": {i.ShortID},
      "type": {i.EntityType},
   }.Encode()
   LogLevel.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   prop := new(Property)
   if err := json.NewDecoder(res.Body).Decode(prop); err != nil {
      return nil, err
   }
   return prop, nil
}

type Property struct {
   Data struct {
      Item Item
   }
}

func (p Property) Base() string {
   var buf strings.Builder
   buf.WriteString(p.Data.Item.ParentEntity.Title)
   buf.WriteByte('-')
   buf.WriteString(p.Data.Item.Title)
   return format.Clean(buf.String())
}

func (p Property) Topaz() (*Topaz, error) {
   req, err := http.NewRequest("GET", p.Data.Item.VideoServiceURL, nil)
   if err != nil {
      return nil, err
   }
   req.URL.RawQuery = "clientPlatform=android"
   LogLevel.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   top := new(Topaz)
   if err := json.NewDecoder(res.Body).Decode(top); err != nil {
      return nil, err
   }
   return top, nil
}

type Topaz struct {
   StitchedStream struct {
      Source string
   }
}
