package bbcamerica

import (
   "encoding/json"
   "github.com/89z/format"
   "net/http"
   "strings"
)

type Playback struct {
   Data struct {
      PlaybackJsonData struct {
         Sources []struct {
            Type string
            Src string
         }
      }
   }
}

func (u Unauth) Playback() (*Playback, error) {
   body := strings.NewReader(`
   {
      "adtags": {
         "lat": 0,
         "url": "https://www.bbcamerica.com/shows/killing-eve/episodes/season-4-just-dunk-me--1052529",
         "playerWidth": 1920,
         "playerHeight": 1080,
         "ppid": 1,
         "mode": "on-demand"
      }
   }
   `)
   req, err := http.NewRequest(
      "POST",
      "https://gw.cds.amcn.com/playback-id/api/v1/playback/1052529",
      body,
   )
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

type Unauth struct {
   Data struct {
      Access_Token string
   }
}
