package channel4

import (
   "encoding/json"
   "github.com/89z/format"
   "net/http"
   "path"
)

var LogLevel format.LogLevel

// channel4.com/programmes/frasier/on-demand/18926-001
func ProgramID(addr string) string {
   return path.Base(addr)
}

func (s Stream) Widevine() string {
   for _, profile := range s.VideoProfiles {
      if profile.Name == "dashwv-dyn-stream-1" {
         for _, stream := range profile.Streams {
            return stream.URI
         }
      }
   }
   return ""
}

type Stream struct {
   VideoProfiles []struct {
      Name string
      Streams []struct {
         URI string
      }
   }
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
