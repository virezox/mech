package amc

import (
   "bytes"
   "encoding/json"
   "net/http"
   "strconv"
   "strings"
)

func (p Playback) Source() *Source {
   for _, source := range p.body.Data.PlaybackJsonData.Sources {
      if source.Type == "application/dash+xml" {
         return &source
      }
   }
   return nil
}

type Playback struct {
   header http.Header
   body struct {
      Data struct {
         PlaybackJsonData struct {
            Custom_Fields struct {
               Show string // 1
               Season string // 2
               Episode string // 3
            }
            Name string // 4
            Sources []Source
         }
      }
   }
}

func (p Playback) Base() string {
   d := p.body.Data.PlaybackJsonData
   var b strings.Builder
   b.WriteString(d.Custom_Fields.Show)
   b.WriteByte('-')
   b.WriteString(d.Custom_Fields.Season)
   b.WriteByte('-')
   b.WriteString(d.Custom_Fields.Episode)
   b.WriteByte('-')
   b.WriteString(d.Name)
   return b.String()
}

func (a Auth) Playback(nID int64) (*Playback, error) {
   addr := []byte("https://gw.cds.amcn.com/playback-id/api/v1/playback/")
   addr = strconv.AppendInt(addr, nID, 10)
   var in playback_request
   in.Ad_Tags.Mode = "on-demand"
   in.Ad_Tags.URL = "!"
   body, err := json.MarshalIndent(in, "", " ")
   if err != nil {
      return nil, err
   }
   req, err := http.NewRequest("POST", string(addr), bytes.NewReader(body))
   if err != nil {
      return nil, err
   }
   req.Header = http.Header{
      "Authorization": {"Bearer " + a.Data.Access_Token},
      "Content-Type": {"application/json"},
      "X-Amcn-Device-Ad-Id": {"!"},
      "X-Amcn-Language": {"en"},
      "X-Amcn-Network": {"amcplus"},
      "X-Amcn-Platform": {"web"},
      "X-Amcn-Service-Id": {"amcplus"},
      "X-Amcn-Tenant": {"amcn"},
      "X-Ccpa-Do-Not-Sell": {"doNotPassData"},
   }
   res, err := Client.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   var out Playback
   out.header = res.Header
   if err := json.NewDecoder(res.Body).Decode(&out.body); err != nil {
      return nil, err
   }
   return &out, nil
}

func (p Playback) Request_URL() string {
   return p.Source().Key_Systems.Widevine.License_URL
}

func (p Playback) Request_Header() http.Header {
   head := make(http.Header)
   jwt := p.header.Get("X-AMCN-BC-JWT")
   head.Set("bcov-auth", jwt)
   return head
}

func (Playback) Request_Body(buf []byte) ([]byte, error) {
   return buf, nil
}

func (Playback) Response_Body(buf []byte) ([]byte, error) {
   return buf, nil
}
