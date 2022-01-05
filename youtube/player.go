package youtube

import (
   "bytes"
   "encoding/json"
   "github.com/89z/format"
   "net/http"
   "strconv"
   "time"
)

const origin = "https://www.youtube.com"

var (
   Android = Client{Name: "ANDROID", Version: "16.43.34"}
   Embed = Client{Name: "ANDROID", Screen: "EMBED", Version: "16.43.34"}
   Mweb = Client{Name: "MWEB", Version: "2.20211109.01.00"}
)

var Key = Auth{"X-Goog-Api-Key", "AIzaSyAO_FJ2SlqU8Q4STEHLGCilw_Y9_11qcW8"}

func post(addr string, head Auth, body youTubeI) (*http.Response, error) {
   buf := new(bytes.Buffer)
   if err := json.NewEncoder(buf).Encode(body); err != nil {
      return nil, err
   }
   req, err := http.NewRequest("POST", addr, buf)
   if err != nil {
      return nil, err
   }
   req.Header.Set(head.Key, head.Value)
   format.Log.Dump(req)
   return new(http.Transport).RoundTrip(req)
}

type Auth struct {
   Key string
   Value string
}

type Client struct {
   Name string `json:"clientName"`
   Screen string `json:"clientScreen,omitempty"`
   Version string `json:"clientVersion"`
}

type PlayabilityStatus struct {
   Status string // LOGIN_REQUIRED
   Reason string // Sign in to confirm your age
}

func (p PlayabilityStatus) Error() string {
   return p.Status + " " + p.Reason
}

type Player struct {
   Microformat struct {
      PlayerMicroformatRenderer struct {
         AvailableCountries []string
         PublishDate string // 2013-06-11
      }
   }
   PlayabilityStatus PlayabilityStatus
   StreamingData struct {
      AdaptiveFormats []Format
      // just including this so I can bail if video is Dash Manifest
      DashManifestURL string
      Formats []Format
   }
   VideoDetails struct {
      Author string
      ShortDescription string
      Title string
      VideoID string
      ViewCount float64 `json:"viewCount,string"`
   }
}

// youtube.com/watch?v=hi8ryzFqrAE
func NewPlayer(id string, head Auth, body Client) (*Player, error) {
   if len(id) != 11 {
      return nil, invalidVideo{id}
   }
   var i youTubeI
   i.Context.Client = body
   i.VideoID = id
   if head != Key {
      i.RacyCheckOK = true
   }
   if body.Screen != "" {
      i.Context.ThirdParty = &thirdParty{origin}
   }
   res, err := post(origin + "/youtubei/v1/player", head, i)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   play := new(Player)
   if err := json.NewDecoder(res.Body).Decode(play); err != nil {
      return nil, err
   }
   return play, nil
}

func (p Player) Date() (time.Time, error) {
   date := p.Microformat.PlayerMicroformatRenderer.PublishDate
   return time.Parse("2006-01-02", date)
}

type invalidVideo struct {
   id string
}

func (i invalidVideo) Error() string {
   return "invalid video ID " + strconv.Quote(i.id)
}

type thirdParty struct {
   EmbedURL string `json:"embedUrl"`
}

type youTubeI struct {
   Context struct {
      Client Client `json:"client"`
      ThirdParty *thirdParty `json:"thirdParty,omitempty"`
   } `json:"context"`
   Continuation string `json:"continuation,omitempty"`
   Params string `json:"params,omitempty"`
   Query string `json:"query,omitempty"`
   RacyCheckOK bool `json:"racyCheckOk,omitempty"`
   VideoID string `json:"videoId,omitempty"`
}
