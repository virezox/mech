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

func (m Media) address() (string, error) {
   for _, video := range m.Connection {
      if video.Protocol == "http" {
         if video.TransferFormat == "hls" {
            if video.Supplier == "mf_akamai" {
               return video.Href, nil
            }
         }
      }
   }
   return "", notFound{"http,hls,mf_akamai"}
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

func (n NewsItem) Media() (*Media, error) {
   addr, err := n.address()
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

func (n NewsItem) address() (string, error) {
   var buf strings.Builder
   buf.WriteString("http://open.live.bbc.co.uk")
   buf.WriteString("/mediaselector/6/select/version/2.0/mediaset/pc/vpid/")
   for _, rel := range n.Relations {
      if rel.PrimaryType == "bbc.mobile.news.video" {
         buf.WriteString(rel.Content.ExternalID)
         return buf.String(), nil
      }
   }
   return "", notFound{"bbc.mobile.news.video"}
}

type notFound struct {
   value string
}

func (n notFound) Error() string {
   return strconv.Quote(n.value) + " not found"
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
