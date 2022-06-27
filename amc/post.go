package amc

import (
   "net/http"
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

func (p Playback) DASH() *Source {
   for _, source := range p.PlaybackJsonData.Sources {
      if source.Type == "application/dash+xml" {
         return &source
      }
   }
   return nil
}

type Playback_JSON_Data struct {
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

type Playback struct {
   BC_JWT string
   PlaybackJsonData Playback_JSON_Data
}

func (p Playback) Request_URL() string {
   return p.DASH().Key_Systems.Widevine.License_URL
}

func (p Playback) Request_Header() http.Header {
   head := make(http.Header)
   head.Set("bcov-auth", p.BC_JWT)
   return head
}

func (p Playback) Request_Body(buf []byte) ([]byte, error) {
   return buf, nil
}

func (p Playback) Response_Body(buf []byte) ([]byte, error) {
   return buf, nil
}
