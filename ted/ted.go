package ted

import (
   "encoding/json"
   "github.com/89z/format"
   "net/http"
   "net/url"
   "strconv"
   "strings"
)

var LogLevel format.LogLevel

type TalkResponse struct {
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
   talk := new(TalkResponse)
   if err := json.NewDecoder(res.Body).Decode(talk); err != nil {
      return nil, err
   }
   return talk, nil
}

type Video struct {
   Bitrate int64
   Size int64
   URL string
}

func (v Video) String() string {
   buf := []byte("Bitrate:")
   buf = strconv.AppendInt(buf, v.Bitrate, 10)
   buf = append(buf, " Size:"...)
   buf = append(buf, format.Size.GetInt64(v.Size)...)
   buf = append(buf, " URL:"...)
   buf = append(buf, v.GetURL()...)
   return string(buf)
}

func (v Video) GetURL() string {
   addr, err := url.Parse(v.URL)
   if err != nil {
      return v.URL
   }
   addr.RawQuery = ""
   return addr.String()
}
