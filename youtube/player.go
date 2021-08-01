package youtube

import (
   "bytes"
   "encoding/json"
   "fmt"
   "net/http"
   "net/http/httputil"
   "os"
)

const origin = "https://www.youtube.com"

var (
   Key = Auth{"X-Goog-Api-Key", "AIzaSyAO_FJ2SlqU8Q4STEHLGCilw_Y9_11qcW8"}
   Verbose = false
)

var (
   Android = Client{Name: "ANDROID", Version: "16.05"}
   Embed = Client{Name: "ANDROID", Screen: "EMBED", Version: "16.05"}
   Mweb = Client{Name: "MWEB", Version: "2.19700101"}
   TV = Client{Name: "TVHTML5", Version: "7.20200101"}
)

func ValidID(id string) error {
   if len(id) == 11 {
      return nil
   }
   return fmt.Errorf("%q invalid as ID", id)
}

func post(url string, head Auth, body youTubeI) (*http.Response, error) {
   buf := new(bytes.Buffer)
   if err := json.NewEncoder(buf).Encode(body); err != nil {
      return nil, err
   }
   req, err := http.NewRequest("POST", url, buf)
   if err != nil {
      return nil, err
   }
   req.Header.Set(head.Key, head.Value)
   if Verbose {
      d, err := httputil.DumpRequest(req, true)
      if err != nil {
         return nil, err
      }
      os.Stdout.Write(d)
   }
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   if res.StatusCode != http.StatusOK {
      return nil, fmt.Errorf("status %v", res.Status)
   }
   return res, nil
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
      AdaptiveFormats FormatSlice
      // just including this so I can bail if video is Dash Manifest
      DashManifestURL string
   }
   VideoDetails struct {
      Author string
      ShortDescription string
      Title string
      ViewCount int `json:"viewCount,string"`
   }
}

func NewPlayer(id string, head Auth, body Client) (*Player, error) {
   var i youTubeI
   i.Context.Client = body
   i.VideoID = id
   // OAuth
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

func (p Player) Views() int {
   return p.VideoDetails.ViewCount
}

type thirdParty struct {
   EmbedURL string `json:"embedUrl"`
}

type youTubeI struct {
   Context struct {
      Client Client `json:"client"`
      ThirdParty *thirdParty `json:"thirdParty,omitempty"`
   } `json:"context"`
   Params string `json:"params,omitempty"`
   Query string `json:"query,omitempty"`
   RacyCheckOK bool `json:"racyCheckOk,omitempty"`
   VideoID string `json:"videoId,omitempty"`
}
