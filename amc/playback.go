package amc
// github.com/89z

import (
   "bytes"
   "errors"
   "github.com/89z/mech"
   "github.com/89z/mech/widevine"
   "io"
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
   buf.WriteString(mech.Clean(p.PlaybackJsonData.Name))
   return buf.String()
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
func (p Playback) DASH() *Source {
   for _, source := range p.PlaybackJsonData.Sources {
      if source.Type == "application/dash+xml" {
         return &source
      }
   }
   return nil
}

func (p Playback) Content(c widevine.Client) (*widevine.Content, error) {
   source := p.DASH()
   mod, err := c.Module()
   if err != nil {
      return nil, err
   }
   in, err := mod.Marshal()
   if err != nil {
      return nil, err
   }
   req, err := http.NewRequest(
      "POST", source.Key_Systems.Widevine.License_URL, bytes.NewReader(in),
   )
   if err != nil {
      return nil, err
   }
   req.Header.Set("bcov-auth", p.BC_JWT)
   LogLevel.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      return nil, errors.New(res.Status)
   }
   out, err := io.ReadAll(res.Body)
   if err != nil {
      return nil, err
   }
   keys, err := mod.Unmarshal(out)
   if err != nil {
      return nil, err
   }
   return keys.Content(), nil
}

type Playback struct {
   PlaybackJsonData PlaybackJsonData
   BC_JWT string
}
