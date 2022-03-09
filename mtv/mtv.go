package mtv

import (
   "encoding/json"
   "github.com/89z/format"
   "net/http"
   "net/url"
   "strings"
   "time"
)

func (i Item) GetDuration() time.Duration {
   return time.Duration(i.Duration.Milliseconds) * time.Millisecond
}

func (i Item) String() string {
   var buf strings.Builder
   buf.WriteString("Date: ")
   buf.WriteString(i.AirDate.DateString)
   buf.WriteString("\nDuration: ")
   buf.WriteString(i.GetDuration().String())
   buf.WriteString("\nType: ")
   buf.WriteString(i.EntityType)
   buf.WriteString("\nMGID: ")
   buf.WriteString(i.MGID)
   buf.WriteString("\nParent: ")
   buf.WriteString(i.ParentEntity.Title)
   buf.WriteString("\nShortID: ")
   buf.WriteString(i.ShortID)
   buf.WriteString("\nTitle: ")
   buf.WriteString(i.Title)
   return buf.String()
}

var LogLevel format.LogLevel

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

func (p Property) Base() string {
   var buf strings.Builder
   buf.WriteString(p.Data.Item.ParentEntity.Title)
   buf.WriteByte('-')
   buf.WriteString(p.Data.Item.Title)
   return format.Clean(buf.String())
}

type Property struct {
   Data struct {
      Item Item
   }
}

type Item struct {
   AirDate struct {
      DateString string
   }
   Duration struct {
      Milliseconds int64
   }
   EntityType string
   MGID string
   ParentEntity struct {
      Title string
   }
   ShortID string
   Title string
}

type Topaz struct {
   Content []struct {
      Chapters []struct {
         ID string // MGID
      }
   }
   StitchedStream struct {
      Source string
   }
}

func NewTopaz(mgid string) (*Topaz, error) {
   var buf strings.Builder
   buf.WriteString("https://topaz.viacomcbs.digital/topaz/api/")
   buf.WriteString(mgid)
   buf.WriteString("/mica.json")
   req, err := http.NewRequest("GET", buf.String(), nil)
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
