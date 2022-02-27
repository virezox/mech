package nbc

import (
   "github.com/89z/format/hls"
   "net/http"
   "strconv"
)

func (a AccessVOD) Streams() ([]Stream, error) {
   req, err := http.NewRequest("GET", a.ManifestPath, nil)
   if err != nil {
      return nil, err
   }
   LogLevel.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   forms, err := hls.Decode(res.Body, "")
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

// #EXTINF
type Information struct {
   FrameRate float64
   URI string
}

// #EXT-X-STREAM-INF
type Stream struct {
   ID int64
   Resolution string
   Bandwidth int64
   Codecs string
   URI string
}

func (s Stream) Information() ([]Information, error) {
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
   forms, err := hls.Decode(res.Body, "")
   if err != nil {
      return nil, err
   }
   var infos []Information
   for _, form := range forms {
      var info Information
      info.FrameRate, err = strconv.ParseFloat(form["frame-rate"], 64)
      if err != nil {
         return nil, err
      }
      info.URI = form["URI"]
      infos = append(infos, info)
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
