package amc

import (
   "github.com/89z/format"
   "net/http"
)

func (p Playback) Header() http.Header {
   head := make(http.Header)
   head.Set("bcov-auth", p.BcJWT)
   return head
}

func (p Playback) DASH() *Source {
   for _, source := range p.Body.Data.PlaybackJsonData.Sources {
      if source.Type == "application/dash+xml" {
         return &source
      }
   }
   return nil
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

var LogLevel format.LogLevel

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
