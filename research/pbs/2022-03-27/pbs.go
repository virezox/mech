package pbs

import (
   "github.com/89z/format"
   "github.com/89z/format/json"
   "net/http"
   "strconv"
)

var LogLevel format.LogLevel

type VideoBridge struct {
   Encodings []string
}

func NewVideoBridge(id int64) (*VideoBridge, error) {
   buf := []byte("https://player.pbs.org/widget/partnerplayer/")
   buf = strconv.AppendInt(buf, id, 10)
   buf = append(buf, '/')
   req, err := http.NewRequest("GET", string(buf), nil)
   if err != nil {
      return nil, err
   }
   LogLevel.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   var (
      sep = []byte("\twindow.videoBridge = ")
      video = new(VideoBridge)
   )
   if err := json.Decode(res.Body, sep, video); err != nil {
      return nil, err
   }
   return video, nil
}
