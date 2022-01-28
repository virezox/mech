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

const videoType = "bbc.mobile.news.video"

var LogLevel format.LogLevel

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

func NewNewsItem(addr string) (*NewsItem, error) {
   var str strings.Builder
   str.WriteString("http://walter-producer-cdn.api.bbci.co.uk")
   str.WriteString("/content/cps/news/")
   str.WriteString(path.Base(addr))
   req, err := http.NewRequest("GET", str.String(), nil)
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

type notFound struct {
   input string
}

func (n notFound) Error() string {
   return strconv.Quote(n.input) + " not found"
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

////////////////////////////////////////////////////////////////////////////////

func (n NewsItem) Media() (*Media, error) {
   var (
      addr strings.Builder
      found bool
   )
   addr.WriteString("http://open.live.bbc.co.uk")
   addr.WriteString("/mediaselector/6/select/version/2.0/mediaset/pc/vpid/")
   for _, rel := range n.Relations {
      if rel.PrimaryType == videoType {
         addr.WriteString(rel.Content.ExternalID)
         found = true
      }
   }
   if !found {
      return nil, notFound{videoType}
   }
   req, err := http.NewRequest("GET", addr.String(), nil)
   if err != nil {
      return nil, err
   }
   LogLevel.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   var mediaset struct {
      Media []Media
   }
   if err := json.NewDecoder(res.Body).Decode(&mediaset); err != nil {
      return nil, err
   }
   for _, media := range mediaset.Media {
      if media.Kind == "video" {
         return &media, nil
      }
   }
   return nil, notFound{"video"}
}

func (m Media) Streams() ([]Stream, error) {
   var href string
   for _, video := range m.Connection {
      if video.Protocol == "http" &&
      video.Supplier == "mf_akamai" &&
      video.TransferFormat == "hls" {
         href = video.Href
      }
   }
   if href == "" {
      return nil, notFound{"http,hls,mf_akamai"}
   }
   req, err := http.NewRequest("GET", href, nil)
   if err != nil {
      return nil, err
   }
   LogLevel.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   dir, _ := path.Split(href)
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
