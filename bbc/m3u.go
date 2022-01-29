package bbc

import (
   "github.com/89z/format/m3u"
   "net/http"
   "path"
   "strconv"
)

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
