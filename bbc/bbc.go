package bbc

import (
   "encoding/json"
   "github.com/89z/format"
   "github.com/89z/format/m3u"
   "net/http"
   "path"
   "strconv"
   "strings"
)

var LogLevel format.LogLevel

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

func (m Media) Streams() ([]Stream, error) {
   addr, err := m.address()
   if err != nil {
      return nil, err
   }
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
   dir, _ := path.Split(addr)
   forms, err := m3u.Decode(res.Body, dir)
   if err != nil {
      return nil, err
   }
   var streams []Stream
   for i, form := range forms {
      var stream Stream
      stream.Bandwidth, err = strconv.ParseInt(form["BANDWIDTH"], 10, 64)
      if err != nil {
         return nil, err
      }
      stream.Codecs = form["CODECS"]
      stream.ID = int64(i)
      stream.Resolution = form["RESOLUTION"]
      stream.URI = form["URI"]
      streams = append(streams, stream)
   }
   return streams, nil
}

// #EXT-X-STREAM-INF
type Stream struct {
   ID int64
   Resolution string
   Bandwidth int64
   Codecs string
   URI string
}

func (s Stream) Information() ([]string, error) {
   req, err := http.NewRequest("GET", s.URI, nil)
   if err != nil {
      return nil, err
   }
   LogLevel.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   dir, _ := path.Split(s.URI)
   forms, err := m3u.Decode(res.Body, dir)
   if err != nil {
      return nil, err
   }
   var infos []string
   for _, form := range forms {
      infos = append(infos, form["URI"])
   }
   return infos, nil
}

func (s Stream) String() string {
   buf := []byte("ID:")
   buf = strconv.AppendInt(buf, s.ID, 10)
   buf = append(buf, " Resolution:"...)
   buf = append(buf, s.Resolution...)
   buf = append(buf, " Bandwidth:"...)
   buf = strconv.AppendInt(buf, s.Bandwidth, 10)
   buf = append(buf, " Codecs:"...)
   buf = append(buf, s.Codecs...)
   return string(buf)
}

type NewsItem struct {
   ShortName string
   IstatsLabels struct {
      CPS_Asset_ID string
   }
   Relations []struct {
      PrimaryType string
      Content struct {
         ExternalID string
      }
   }
}

type Media struct {
   Kind string
   Type string
   Connection []struct {
      Protocol string
      Supplier string
      TransferFormat string
      Href string
   }
}

func (m Media) Name(item *NewsItem) (string, error) {
   ext, err := format.ExtensionByType(m.Type)
   if err != nil {
      return "", err
   }
   return item.ShortName + "-" + item.IstatsLabels.CPS_Asset_ID + ext, nil
}
