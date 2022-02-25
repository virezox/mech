package mtv

import (
   "encoding/json"
   "github.com/89z/format"
   "net/http"
   "net/url"
   "strings"
)

var LogLevel format.LogLevel

type property struct {
   Data struct {
      Item struct {
         VideoServiceURL string
      }
   }
}

func newProperty(typ, shortID string) (*property, error) {
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
   prop := new(property)
   if err := json.NewDecoder(res.Body).Decode(prop); err != nil {
      return nil, err
   }
   return prop, nil
}

////////////////////////////////////////////////////////////////////////////////

func topaz() (*http.Response, error) {
   var buf strings.Builder
   buf.WriteString("https://topaz.viacomcbs.digital/topaz/api/")
   buf.WriteString("mgid:arc:episode:android.playplex.mtv.com:7d923439-a492-11ea-9225-70df2f866ace/")
   buf.WriteString("mica.json")
   req, err := http.NewRequest("GET", buf.String(), nil)
   if err != nil {
      return nil, err
   }
   req.URL.RawQuery = "clientPlatform=android"
   LogLevel.Dump(req)
   return new(http.Transport).RoundTrip(req)
}
