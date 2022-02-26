package bbc

import (
   "github.com/89z/format/m3u"
   "net/http"
   "path"
   "strconv"
)

// #EXT-X-STREAM-INF
type Stream struct {
   ID int64
   Resolution string
   Bandwidth int64
   Codecs string
   URI string
}

func (c Connection) Streams() ([]Stream, error) {
   req, err := http.NewRequest("GET", c.Href, nil)
   if err != nil {
      return nil, err
   }
   LogLevel.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   dir, _ := path.Split(c.Href)
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
