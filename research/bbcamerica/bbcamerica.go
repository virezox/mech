package bbcamerica

import (
   "bytes"
   "encoding/json"
   "github.com/89z/format"
   "net/http"
   "strconv"
)

func (u Unauth) Header() http.Header {
   head := make(http.Header)
   head.Set("bcov-auth", u.Data.Access_Token)
   return head
}

type Unauth struct {
   Data struct {
      Access_Token string
   }
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

func (p Playback) DASH() *Source {
   for _, source := range p.Data.PlaybackJsonData.Sources {
      if source.Type == "application/dash+xml" {
         return &source
      }
   }
   return nil
}

type Playback struct {
   Data struct {
      PlaybackJsonData struct {
         Sources []Source
      }
   }
}

func (u Unauth) Playback(nid int64) (*Playback, error) {
   var (
      addr []byte
      body playbackRequest
   )
   addr = append(addr, "https://gw.cds.amcn.com/playback-id/api/v1/playback/"...)
   addr = strconv.AppendInt(addr, nid, 10)
   body.AdTags.Mode = "on-demand"
   body.AdTags.URL = "!"
   buf := new(bytes.Buffer)
   if err := json.NewEncoder(buf).Encode(body); err != nil {
      return nil, err
   }
   req, err := http.NewRequest("POST", string(addr), buf)
   if err != nil {
      return nil, err
   }
   req.Header = http.Header{
      "Authorization": {"Bearer " + u.Data.Access_Token},
      "Content-Type": {"application/json"},
      "X-Amcn-Device-Ad-Id": {"!"},
      "X-Amcn-Language": {"en"},
      "X-Amcn-Network": {"bbca"},
      "X-Amcn-Platform": {"web"},
      "X-Amcn-Service-Id": {"bbca"},
      "X-Amcn-Tenant": {"amcn"},
      "X-Ccpa-Do-Not-Sell": {"passData"},
   }
   LogLevel.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   play := new(Playback)
   if err := json.NewDecoder(res.Body).Decode(play); err != nil {
      return nil, err
   }
   return play, nil
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

var LogLevel format.LogLevel

func NewUnauth() (*Unauth, error) {
   req, err := http.NewRequest(
      "POST", "https://gw.cds.amcn.com/auth-orchestration-id/api/v1/unauth", nil,
   )
   if err != nil {
      return nil, err
   }
   req.Header = http.Header{
      "X-Amcn-Device-Id": {"!"},
      "X-Amcn-Language": {"en"},
      "X-Amcn-Network": {"bbca"},
      "X-Amcn-Platform": {"web"},
      "X-Amcn-Tenant": {"amcn"},
   }
   LogLevel.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   auth := new(Unauth)
   if err := json.NewDecoder(res.Body).Decode(auth); err != nil {
      return nil, err
   }
   return auth, nil
}
