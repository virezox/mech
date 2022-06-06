// github.com/89z
package amc

import (
   "github.com/89z/format"
   "github.com/89z/mech"
   "net/http"
   "strconv"
   "strings"
)

var LogLevel format.LogLevel

func GetNID(input string) (int64, error) {
   _, nid, found := strings.Cut(input, "--")
   if found {
      input = nid
   }
   return strconv.ParseInt(input, 10, 64)
}

type Playback struct {
   PlaybackJsonData PlaybackJsonData
   BcJWT string
}

func (p Playback) Base() string {
   var buf strings.Builder
   buf.WriteString(p.PlaybackJsonData.Custom_Fields.Show)
   buf.WriteByte('-')
   buf.WriteString(p.PlaybackJsonData.Custom_Fields.Season)
   buf.WriteByte('-')
   buf.WriteString(p.PlaybackJsonData.Custom_Fields.Episode)
   buf.WriteByte('-')
   buf.WriteString(mech.Clean(p.PlaybackJsonData.Name))
   return buf.String()
}

func (p Playback) DASH() *Source {
   for _, source := range p.PlaybackJsonData.Sources {
      if source.Type == "application/dash+xml" {
         return &source
      }
   }
   return nil
}

func (p Playback) Header() http.Header {
   head := make(http.Header)
   head.Set("bcov-auth", p.BcJWT)
   return head
}

type PlaybackJsonData struct {
   Custom_Fields struct {
      Show string // 1
      Season string // 2
      Episode string // 3
   }
   Name string // 4
   Sources []Source
}

type Source struct {
   Key_Systems *struct {
      Widevine struct {
         License_URL string
      } `json:"com.widevine.alpha"`
   }
   Src string
   Type string
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
