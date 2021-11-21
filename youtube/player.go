package youtube

import (
   "bytes"
   "encoding/json"
   "github.com/89z/mech"
   "net/http"
)

const origin = "https://www.youtube.com"

var (
   Android = Client{Name: "ANDROID", Version: "16.43.34"}
   Embed = Client{Name: "ANDROID", Screen: "EMBED", Version: "16.43.34"}
   Mweb = Client{Name: "MWEB", Version: "2.20211109.01.00"}
)

// youtube.com/watch?v=hi8ryzFqrAE
func Valid(id string) bool {
   return len(id) == 11
}

type Client struct {
   Name string `json:"clientName"`
   Screen string `json:"clientScreen,omitempty"`
   Version string `json:"clientVersion"`
}

type Player struct {
   Microformat struct {
      PlayerMicroformatRenderer struct {
         AvailableCountries []string
         PublishDate string
      }
   }
   PlayabilityStatus struct {
      Reason string
      Status string
   }
   StreamingData struct {
      AdaptiveFormats []Format
      // just including this so I can bail if video is Dash Manifest
      DashManifestURL string
   }
   VideoDetails struct {
      Author string
      ShortDescription string
      Title string
      ViewCount float64 `json:"viewCount,string"`
   }
}

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
   return mech.RoundTrip(req)
}

type Auth struct {
   Key string
   Value string
}

func NewPlayer(id string, head Auth, body Client) (*Player, error) {
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
   p := new(Player)
   if err := json.NewDecoder(res.Body).Decode(p); err != nil {
      return nil, err
   }
   return p, nil
}

func (p Player) Author() string {
   return p.VideoDetails.Author
}

func (p Player) Countries() []string {
   return p.Microformat.PlayerMicroformatRenderer.AvailableCountries
}

func (p Player) Date() string {
   return p.Microformat.PlayerMicroformatRenderer.PublishDate
}

func (p Player) Description() string {
   return p.VideoDetails.ShortDescription
}

func (p Player) Title() string {
   return p.VideoDetails.Title
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
