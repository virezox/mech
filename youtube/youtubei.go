package youtube

import (
   "bytes"
   "encoding/json"
   "net/http"
   "net/url"
   "strconv"
   "strings"
   "time"
)

const origin = "https://www.youtube.com"

var (
   Mweb = Client{Name: "MWEB", Version: "2.20211109.01.00"}
   // com.google.android.youtube
   Android = Client{Name: "ANDROID", Version: "17.06.32"}
   Embed = Client{Name: "ANDROID", Screen: "EMBED", Version: "17.06.32"}
)

var googAPI = http.Header{
   "X-Goog-Api-Key": {"AIzaSyAO_FJ2SlqU8Q4STEHLGCilw_Y9_11qcW8"},
}

// youtube.com/watch?v=XY-hOqcPGCY
func VideoID(address string) (string, error) {
   addr, err := url.Parse(address)
   if err != nil {
      return "", err
   }
   return addr.Query().Get("v"), nil
}

type Client struct {
   Name string `json:"clientName"`
   Screen string `json:"clientScreen,omitempty"`
   Version string `json:"clientVersion"`
}

func (c Client) Player(id string) (*Player, error) {
   return c.PlayerHeader(googAPI, id)
}

func (c Client) PlayerHeader(head http.Header, id string) (*Player, error) {
   res, err := c.player(head, id)
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

func (c Client) PlayerResponse(id string) (*http.Response, error) {
   return c.player(googAPI, id)
}

func (c Client) Search(query string) (*Search, error) {
   var body searchRequest
   body.Context.Client = c
   body.Query = query
   filter := NewFilter().Type(TypeVideo)
   body.Params = NewParams().Filter(filter).Encode()
   buf := new(bytes.Buffer)
   if err := json.NewEncoder(buf).Encode(body); err != nil {
      return nil, err
   }
   req, err := http.NewRequest("POST", origin + "/youtubei/v1/search", buf)
   if err != nil {
      return nil, err
   }
   req.Header = googAPI
   LogLevel.Dump(req)
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

func (c Client) player(head http.Header, id string) (*http.Response, error) {
   // youtube.com/watch?v=hi8ryzFqrAE
   if len(id) != 11 {
      return nil, invalidVideo{id}
   }
   var body playerRequest
   body.VideoID = id
   body.Context.Client = c
   body.RacyCheckOK = true
   if c.Screen != "" {
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
   req.Header = head
   LogLevel.Dump(req)
   return new(http.Transport).RoundTrip(req)
}

type Item struct {
   CompactVideoRenderer *struct {
      LengthText text
      Title text
      VideoID string
   }
}

type PlayabilityStatus struct {
   Status string // "OK", "LOGIN_REQUIRED"
   Reason string // "", "Sign in to confirm your age"
}

func (p PlayabilityStatus) String() string {
   var buf strings.Builder
   buf.WriteString("Status: ")
   buf.WriteString(p.Status)
   if p.Reason != "" {
      buf.WriteString("\nReason: ")
      buf.WriteString(p.Reason)
   }
   return buf.String()
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
   VideoDetails VideoDetails
}

func (p Player) Date() (time.Time, error) {
   date := p.Microformat.PlayerMicroformatRenderer.PublishDate
   return time.Parse("2006-01-02", date)
}

func (p Player) base() string {
   return p.VideoDetails.Author + "-" + p.VideoDetails.Title
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

type VideoDetails struct {
   VideoID string
   LengthSeconds int64 `json:"lengthSeconds,string"`
   ViewCount int64 `json:"viewCount,string"`
   Author string
   Title string
   ShortDescription string
}

func (v VideoDetails) Duration() time.Duration {
   return time.Duration(v.LengthSeconds) * time.Second
}

func (v VideoDetails) String() string {
   buf := []byte("VideoID: ")
   buf = append(buf, v.VideoID...)
   buf = append(buf, "\nLength: "...)
   buf = append(buf, v.Duration().String()...)
   buf = append(buf, "\nViewCount: "...)
   buf = strconv.AppendInt(buf, v.ViewCount, 10)
   buf = append(buf, "\nAuthor: "...)
   buf = append(buf, v.Author...)
   buf = append(buf, "\nTitle: "...)
   buf = append(buf, v.Title...)
   return string(buf)
}

type invalidVideo struct {
   id string
}

func (i invalidVideo) Error() string {
   return "invalid video ID " + strconv.Quote(i.id)
}

type playerRequest struct {
   Context struct {
      Client Client `json:"client"`
      ThirdParty *thirdParty `json:"thirdParty,omitempty"`
   } `json:"context"`
   RacyCheckOK bool `json:"racyCheckOk,omitempty"`
   VideoID string `json:"videoId,omitempty"`
}

type searchRequest struct {
   Context struct {
      Client Client `json:"client"`
   } `json:"context"`
   Params string `json:"params,omitempty"`
   Query string `json:"query,omitempty"`
}

type text struct {
   Runs []struct {
      Text string
   }
}

type thirdParty struct {
   EmbedURL string `json:"embedUrl"`
}
