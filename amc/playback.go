package amc

import (
   "bytes"
   "github.com/89z/mech"
   "github.com/89z/mech/widevine"
   "io"
   "net/http"
   "strings"
)

func (p Playback) Content(c widevine.Client) (*widevine.Content, error) {
   source := p.DASH()
   module, err := c.Key_ID()
   if err != nil {
      return nil, err
   }
   buf, err := module.Marshal()
   if err != nil {
      return nil, err
   }
   req, err := http.NewRequest(
      "POST", source.Key_Systems.Widevine.License_URL, bytes.NewReader(buf),
   )
   if err != nil {
      return nil, err
   }
   req.Header.Set("bcov-auth", p.BC_JWT)
   res, err := Client.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   buf, err = io.ReadAll(res.Body)
   if err != nil {
      return nil, err
   }
   keys, err := module.Unmarshal(buf)
   if err != nil {
      return nil, err
   }
   return keys.Content(), nil
}

type Playback struct {
   PlaybackJsonData Playback_JSON_Data
   BC_JWT string
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

func (p Playback) DASH() *Source {
   for _, source := range p.PlaybackJsonData.Sources {
      if source.Type == "application/dash+xml" {
         return &source
      }
   }
   return nil
}
