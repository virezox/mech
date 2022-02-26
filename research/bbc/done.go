package bbc

import (
   "encoding/json"
   "github.com/89z/format"
   "net/http"
   "path"
   "strings"
)

var LogLevel format.LogLevel

type Connection struct {
   Protocol string
   Supplier string
   TransferFormat string
   Href string
}

type NewsItem struct {
   ShortName string
   IstatsLabels struct {
      CPS_Asset_ID string
   }
   Relations []Relation
}

func NewNewsItem(addr string) (*NewsItem, error) {
   var buf strings.Builder
   buf.WriteString("http://walter-producer-cdn.api.bbci.co.uk")
   buf.WriteString("/content/cps/news/")
   buf.WriteString(path.Base(addr))
   req, err := http.NewRequest("GET", buf.String(), nil)
   if err != nil {
      return nil, err
   }
   LogLevel.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   item := new(NewsItem)
   if err := json.NewDecoder(res.Body).Decode(item); err != nil {
      return nil, err
   }
   return item, nil
}
