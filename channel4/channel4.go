package channel4

import (
   "encoding/json"
   "github.com/89z/format"
   "net/http"
)

var LogLevel format.LogLevel

type Stream struct {
   VideoProfiles []VideoProfile
}

func NewStream(id string) (*Stream, error) {
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
   str := new(Stream)
   if err := json.NewDecoder(res.Body).Decode(str); err != nil {
      return nil, err
   }
   return str, nil
}

func (s Stream) Widevine() *VideoProfile {
   for _, profile := range s.VideoProfiles {
      if profile.Name == "dashwv-dyn-stream-1" {
         return &profile
      }
   }
   return nil
}

type VideoProfile struct {
   Name string
   Streams []struct {
      URI string
   }
}
