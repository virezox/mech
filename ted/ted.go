package ted

import (
   "encoding/json"
   "fmt"
   "github.com/89z/format"
   "net/http"
   "strings"
)

func (v Video) Format(f fmt.State, verb rune) {
   fmt.Fprint(f, "Bitrate:", v.Bitrate)
   fmt.Fprint(f, " Size:", v.Size)
   if verb == 'a' {
      fmt.Fprint(f, " URL:", v.URL)
   }
}

var LogLevel format.LogLevel

func NewTalkResponse(slug string) (*TalkResponse, error) {
   var buf strings.Builder
   buf.WriteString("https://devices.ted.com/api/v2/videos/")
   buf.WriteString(slug)
   buf.WriteString("/react_native_v2.json")
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
   if res.StatusCode != http.StatusOK {
      return nil, errorString(res.Status)
   }
   talk := new(TalkResponse)
   if err := json.NewDecoder(res.Body).Decode(talk); err != nil {
      return nil, err
   }
   return talk, nil
}

type errorString string

func (e errorString) Error() string {
   return string(e)
}

type Video struct {
   Bitrate int64
   Size int64
   URL string
}

type TalkResponse struct {
   Downloads struct {
      Video []Video
   }
}

func (t TalkResponse) Video(bitrate int64) *Video {
   distance := func(v *Video) int64 {
      if v.Bitrate > bitrate {
         return v.Bitrate - bitrate
      }
      return bitrate - v.Bitrate
   }
   var dst *Video
   for i, src := range t.Downloads.Video {
      if dst == nil || distance(&src) < distance(dst) {
         dst = &t.Downloads.Video[i]
      }
   }
   return dst
}
