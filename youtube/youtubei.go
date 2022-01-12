package youtube

import (
   "bytes"
   "encoding/json"
   "github.com/89z/format"
   "net/http"
)

const origin = "https://www.youtube.com"

var (
   Android = Client{Name: "ANDROID", Version: "16.43.34"}
   Embed = Client{Name: "ANDROID", Screen: "EMBED", Version: "16.43.34"}
   Mweb = Client{Name: "MWEB", Version: "2.20211109.01.00"}
)

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
      ViewCount int `json:"viewCount,string"`
   }
}

func (s Search) Items() []Item {
   var items []Item
   for _, sect := range s.Contents.SectionListRenderer.Contents {
      if sect.ItemSectionRenderer != nil {
         for _, item := range sect.ItemSectionRenderer.Contents {
            if item.CompactVideoRenderer != nil {
               items = append(items, item)
            }
         }
      }
   }
   return items
}

type text struct {
   Runs []struct {
      Text string
   }
}

type Item struct {
   CompactVideoRenderer *struct {
      LengthText text
      Title text
      VideoID string
   }
}

type Search struct {
   Contents struct {
      SectionListRenderer struct {
         Contents []struct {
            ItemSectionRenderer *struct {
               Contents []Item
            }
         }
      }
   }
}

type thirdParty struct {
   EmbedURL string `json:"embedUrl"`
}

type playerRequest struct {
   Context struct {
      Client Client `json:"client"`
      ThirdParty *thirdParty `json:"thirdParty,omitempty"`
   } `json:"context"`
   RacyCheckOK bool `json:"racyCheckOk,omitempty"`
   VideoID string `json:"videoId,omitempty"`
}

func NewPlayer(id string, auth Auth, client Client) (*Player, error) {
   var body playerRequest
   body.VideoID = id
   body.Context.Client = client
   if auth != Key {
      body.RacyCheckOK = true
   }
   if client.Screen != "" {
      body.Context.ThirdParty = &thirdParty{origin}
   }
   buf := new(bytes.Buffer)
   if err := json.NewEncoder(buf).Encode(body); err != nil {
      return nil, err
   }
   req, err := http.NewRequest("POST", origin + "/youtubei/v1/player", buf)
   if err != nil {
      return nil, err
   }
   req.Header.Set(auth.Key, auth.Value)
   format.Log.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
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

type searchRequest struct {
   Context struct {
      Client Client `json:"client"`
   } `json:"context"`
   Params string `json:"params,omitempty"`
   Query string `json:"query,omitempty"`
}

var Key = Auth{"X-Goog-Api-Key", "AIzaSyAO_FJ2SlqU8Q4STEHLGCilw_Y9_11qcW8"}

type Auth struct {
   Key string
   Value string
}

func NewSearch(query string) (*Search, error) {
   var body searchRequest
   body.Context.Client = Mweb
   body.Query = query
   param := NewParams()
   filter := NewFilter()
   filter.Type(TypeVideo)
   param.Filter(filter)
   body.Params = param.Encode()
   buf := new(bytes.Buffer)
   if err := json.NewEncoder(buf).Encode(body); err != nil {
      return nil, err
   }
   req, err := http.NewRequest("POST", origin + "/youtubei/v1/search", buf)
   if err != nil {
      return nil, err
   }
   req.Header.Set(Key.Key, Key.Value)
   format.Log.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   search := new(Search)
   if err := json.NewDecoder(res.Body).Decode(search); err != nil {
      return nil, err
   }
   return search, nil
}
