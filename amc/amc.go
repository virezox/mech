// github.com/89z
package amc

import (
   "github.com/89z/format"
   "net/http"
   "strconv"
   "strings"
)

func (p Playback) Base() string {
   var buf strings.Builder
   buf.WriteString(p.PlaybackJsonData.Custom_Fields.Show)
   buf.WriteByte('-')
   buf.WriteString(p.PlaybackJsonData.Custom_Fields.Season)
   buf.WriteByte('-')
   buf.WriteString(p.PlaybackJsonData.Custom_Fields.Episode)
   buf.WriteByte('-')
   buf.WriteString(p.PlaybackJsonData.Name)
   return buf.String()
}

type PlaybackJsonData struct {
   Sources []Source
   Custom_Fields struct {
      Show string // 1
      Season string // 2
      Episode string // 3
   }
   Name string // 4
}

type Playback struct {
   PlaybackJsonData PlaybackJsonData
   BcJWT string
}

var LogLevel format.LogLevel

func GetNID(addr string) (int64, error) {
   _, nid, ok := strings.Cut(addr, "--")
   if !ok {
      return 0, notFound{"--"}
   }
   return strconv.ParseInt(nid, 10, 64)
}

func (p Playback) Header() http.Header {
   head := make(http.Header)
   head.Set("bcov-auth", p.BcJWT)
   return head
}

type Source struct {
   Key_Systems struct {
      Widevine struct {
         License_URL string
      } `json:"com.widevine.alpha"`
   }
   Src string
   Type string
}

type notFound struct {
   value string
}

func (n notFound) Error() string {
   var buf []byte
   buf = strconv.AppendQuote(buf, n.value)
   buf = append(buf, " is not found"...)
   return string(buf)
}

type playbackRequest struct {
   AdTags struct {
      Lat int `json:"lat"`
      Mode string `json:"mode"`
      PPID int `json:"ppid"`
      PlayerHeight int `json:"playerHeight"`
      PlayerWidth int `json:"playerWidth"`
      URL string `json:"url"`
   } `json:"adtags"`
}

type playbackResponse struct {
   Data struct {
      PlaybackJsonData PlaybackJsonData
   }
}

func (p Playback) DASH() *Source {
   for _, source := range p.PlaybackJsonData.Sources {
      if source.Type == "application/dash+xml" {
         return &source
      }
   }
   return nil
}
