package mtv

import (
   "encoding/json"
   "github.com/89z/format"
   "net/http"
   "net/url"
)

var LogLevel format.LogLevel

type Property struct {
   Data struct {
      Item struct {
         VideoServiceURL string
      }
   }
}

func NewProperty(typ, shortID string) (*Property, error) {
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
      "shortId": {shortID},
      "type": {typ},
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
