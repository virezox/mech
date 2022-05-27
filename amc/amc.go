package amc

import (
   "github.com/89z/format"
   "net/http"
   "strconv"
   "strings"
)

var LogLevel format.LogLevel

func GetNID(addr string) (int64, error) {
   _, nid, ok := strings.Cut(addr, "--")
   if !ok {
      return 0, notFound{"--"}
   }
   return strconv.ParseInt(nid, 10, 64)
}

type Playback struct {
   BcJWT string
   Body struct {
      Data struct {
         PlaybackJsonData struct {
            Name string
            Sources []Source
         }
      }
   }
}

func (p Playback) DASH() *Source {
   for _, source := range p.Body.Data.PlaybackJsonData.Sources {
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
