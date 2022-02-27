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
   Href string
   TransferFormat string
}

type Media struct {
   Kind string
   Type string
   Connection []Connection
}

func (m Media) GetConnection() *Connection {
   for _, con := range m.Connection {
      if con.TransferFormat == "hls" {
         return &con
      }
   }
   return nil
}

func (m Media) Name(item *NewsItem) (string, error) {
   ext, err := format.ExtensionByType(m.Type)
   if err != nil {
      return "", err
   }
   return item.base() + ext, nil
}

type Mediaset struct {
   Media []Media
}

func (m Mediaset) GetMedia() *Media {
   for _, med := range m.Media {
      if med.Kind == "video" {
         return &med
      }
   }
   return nil
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

func (n NewsItem) base() string {
   return n.ShortName + "-" + n.IstatsLabels.CPS_Asset_ID
}

func (n NewsItem) Relation() *Relation {
   for _, rel := range n.Relations {
      if rel.PrimaryType == "bbc.mobile.news.video" {
         return &rel
      }
   }
   return nil
}

type Relation struct {
   PrimaryType string
   Content struct {
      ExternalID string
   }
}

func (r Relation) Mediaset() (*Mediaset, error) {
   var buf strings.Builder
   buf.WriteString("http://open.live.bbc.co.uk")
   buf.WriteString("/mediaselector/6/select/version/2.0/mediaset/pc/vpid/")
   buf.WriteString(r.Content.ExternalID)
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
   set := new(Mediaset)
   if err := json.NewDecoder(res.Body).Decode(set); err != nil {
      return nil, err
   }
   return set, nil
}
