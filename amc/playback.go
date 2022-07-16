package amc

import (
   "bytes"
   "encoding/json"
   "net/http"
   "strconv"
   "strings"
)

type playback_request struct {
   Ad_Tags struct {
      Lat int `json:"lat"`
      Mode string `json:"mode"`
      PPID int `json:"ppid"`
      Player_Height int `json:"playerHeight"`
      Player_Width int `json:"playerWidth"`
      URL string `json:"url"`
   } `json:"adtags"`
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

func (Playback) Request_Body(b []byte) ([]byte, error) {
   return b, nil
}

func (Playback) Response_Body(b []byte) ([]byte, error) {
   return b, nil
}

func (self Data) Source() *Source {
   for _, item := range self.Sources {
      if item.Type == "application/dash+xml" {
         return &item
      }
   }
   return nil
}

func (self Playback) Request_Header() http.Header {
   token := self.head.Get("X-AMCN-BC-JWT")
   head := make(http.Header)
   head.Set("bcov-auth", token)
   return head
}

type Data struct {
   Custom_Fields struct {
      Show string
      Season string
      Episode string
   }
   Name string
   Sources []Source
}

func (self Data) Base() string {
   var b strings.Builder
   b.WriteString(self.Custom_Fields.Show)
   b.WriteByte('-')
   b.WriteString(self.Custom_Fields.Season)
   b.WriteByte('-')
   b.WriteString(self.Custom_Fields.Episode)
   b.WriteByte('-')
   b.WriteString(self.Name)
   return b.String()
}

func (self Auth) Playback(nID int64) (*Playback, error) {
   var b []byte
   b = append(b, "https://gw.cds.amcn.com/playback-id/api/v1/playback/"...)
   b = strconv.AppendInt(b, nID, 10)
   var p playback_request
   p.Ad_Tags.Mode = "on-demand"
   p.Ad_Tags.URL = "!"
   body, err := json.MarshalIndent(p, "", " ")
   if err != nil {
      return nil, err
   }
   req, err := http.NewRequest("POST", string(b), bytes.NewReader(body))
   if err != nil {
      return nil, err
   }
   req.Header = http.Header{
      "Authorization": {"Bearer " + self.Data.Access_Token},
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
   var play Playback
   if err := json.NewDecoder(res.Body).Decode(&play.body); err != nil {
      return nil, err
   }
   play.head = res.Header
   return &play, nil
}

type Playback struct {
   head http.Header
   body struct {
      Data struct {
         PlaybackJsonData Data
      }
   }
}

func (self Playback) Data() Data {
   return self.body.Data.PlaybackJsonData
}

func (self Playback) Request_URL() string {
   return self.Data().Source().Key_Systems.Widevine.License_URL
}
