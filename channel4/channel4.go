package channel4

import (
   "encoding/json"
   "github.com/89z/format"
   "net/http"
)

var LogLevel format.LogLevel

type Stream struct {
   URI string
}

type Video struct {
   VideoProfiles []struct {
      Name string
      Streams []Stream
   }
}

func NewVideo(id string) (*Video, error) {
   req, err := http.NewRequest(
      "GET", "https://www.channel4.com/vod/stream/" + id, nil,
   )
   if err != nil {
      return nil, err
   }
   req.Header.Set("X-Forwarded-For", "25.0.0.0")
   LogLevel.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   vid := new(Video)
   if err := json.NewDecoder(res.Body).Decode(vid); err != nil {
      return nil, err
   }
   return vid, nil
}

func (v Video) Widevine() *Stream {
   for _, profile := range v.VideoProfiles {
      if profile.Name == "dashwv-dyn-stream-1" {
         for _, stream := range profile.Streams {
            return &stream
         }
      }
   }
   return nil
}
