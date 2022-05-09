package ted

import (
   "encoding/json"
   "fmt"
   "github.com/89z/format"
   "net/http"
   "strings"
   "time"
)

var LogLevel format.LogLevel

type TalkResponse struct {
   SpeakerName string
   Title string
   FilmedTimestamp int64
   Downloads struct {
      Video []Video
   }
}

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

func (t TalkResponse) Format(f fmt.State, verb rune) {
   fmt.Fprintln(f, "Speaker:", t.SpeakerName)
   fmt.Fprintln(f, "Title:", t.Title)
   fmt.Fprint(f, "Time: ", t.Time())
   for _, v := range t.Downloads.Video {
      fmt.Fprint(f, "\nBitrate:", v.Bitrate)
      fmt.Fprint(f, " Size:", v.Size)
      fmt.Fprint(f, " Format:", v.Format)
      if verb == 'a' {
         fmt.Fprint(f, " URL:", v.URL)
      }
   }
}

func (t TalkResponse) Time() time.Time {
   return time.UnixMilli(t.FilmedTimestamp)
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

type Video struct {
   Bitrate int64
   Size int64
   Format string
   URL string
}

type errorString string

func (e errorString) Error() string {
   return string(e)
}
