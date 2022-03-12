package youtube

import (
   "bytes"
   "encoding/json"
   "github.com/89z/format"
   "net/http"
   "net/url"
   "path"
   "strconv"
   "strings"
   "time"
)

const origin = "https://www.youtube.com"

var googAPI = http.Header{
   "X-Goog-Api-Key": {"AIzaSyAO_FJ2SlqU8Q4STEHLGCilw_Y9_11qcW8"},
}

// https://youtube.com/shorts/9Vsdft81Q6w
// https://youtube.com/watch?v=XY-hOqcPGCY
func VideoID(address string) (string, error) {
   parse, err := url.Parse(address)
   if err != nil {
      return "", err
   }
   v := parse.Query().Get("v")
   if v != "" {
      return v, nil
   }
   return path.Base(parse.Path), nil
}

func encode(val interface{}) (*bytes.Buffer, error) {
   buf := new(bytes.Buffer)
   enc := json.NewEncoder(buf)
   enc.SetIndent("", " ")
   err := enc.Encode(val)
   if err != nil {
      return nil, err
   }
   return buf, nil
}

type Client struct {
   Name string `json:"clientName"`
   Screen string `json:"clientScreen,omitempty"`
   Version string `json:"clientVersion"`
}

type Context struct {
   Client Client `json:"client"`
   ThirdParty *ThirdParty `json:"thirdParty,omitempty"`
}

var Android = Context{
   Client: Client{Name: "ANDROID", Version: "17.09.33"},
}

// HsUATh_Nc2U
var Embed = Context{
   Client: Client{Name: "ANDROID", Screen: "EMBED", Version: "17.09.33"},
   ThirdParty: &ThirdParty{EmbedURL: origin},
}

var Mweb = Context{
   Client: Client{Name: "MWEB", Version: "2.20211109.01.00"},
}

func (c Context) Player(id string) (*Player, error) {
   return c.PlayerHeader(googAPI, id)
}

func (c Context) PlayerHeader(head http.Header, id string) (*Player, error) {
   var body struct {
      Context Context `json:"context"`
      RacyCheckOK bool `json:"racyCheckOk,omitempty"`
      VideoID string `json:"videoId"`
   }
   body.Context = c
   body.VideoID = id
   if head.Get("Authorization") != "" {
      body.RacyCheckOK = true // Cr381pDsSsA
   }
   buf, err := encode(body)
   if err != nil {
      return nil, err
   }
   req, err := http.NewRequest("POST", origin + "/youtubei/v1/player", buf)
   if err != nil {
      return nil, err
   }
   req.Header = head
   LogLevel.Dump(req)
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

func (c Context) Search(query string) (*Search, error) {
   var body struct {
      Context Context `json:"context"`
      Params string `json:"params"`
      Query string `json:"query"`
   }
   body.Query = query
   filter := NewFilter().Type(TypeVideo)
   body.Params = NewParams().Filter(filter).Encode()
   body.Context = c
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

type Height struct {
   StreamingData
   Target int
}

func (h Height) Less(i, j int) bool {
   return h.distance(i) < h.distance(j)
}

func (h Height) distance(i int) int {
   diff := h.AdaptiveFormats[i].Height - h.Target
   if diff >= 0 {
      return diff
   }
   return -diff
}

type Item struct {
   CompactVideoRenderer *struct {
      Title struct {
         Runs []struct {
            Text string
         }
      }
      VideoID string
   }
}

type Player struct {
   StreamingData StreamingData
   Microformat struct {
      PlayerMicroformatRenderer struct {
         PublishDate string // 2013-06-11
      }
   }
   PlayabilityStatus struct {
      Status string // "OK", "LOGIN_REQUIRED"
      Reason string // "", "Sign in to confirm your age"
   }
   VideoDetails struct {
      VideoID string
      LengthSeconds int64 `json:"lengthSeconds,string"`
      ViewCount int64 `json:"viewCount,string"`
      Author string
      Title string
      ShortDescription string
   }
}

func (p Player) Base() string {
   return format.Clean(p.VideoDetails.Author + "-" + p.VideoDetails.Title)
}

func (p Player) Date() (time.Time, error) {
   value := p.Microformat.PlayerMicroformatRenderer.PublishDate
   return time.Parse("2006-01-02", value)
}

func (p Player) Details() string {
   buf := []byte("VideoID: ")
   buf = append(buf, p.VideoDetails.VideoID...)
   buf = append(buf, "\nLength: "...)
   buf = strconv.AppendInt(buf, p.VideoDetails.LengthSeconds, 10)
   buf = append(buf, "\nViewCount: "...)
   buf = strconv.AppendInt(buf, p.VideoDetails.ViewCount, 10)
   buf = append(buf, "\nAuthor: "...)
   buf = append(buf, p.VideoDetails.Author...)
   buf = append(buf, "\nTitle: "...)
   buf = append(buf, p.VideoDetails.Title...)
   return string(buf)
}

func (p Player) Status() string {
   var buf strings.Builder
   buf.WriteString("Status: ")
   buf.WriteString(p.PlayabilityStatus.Status)
   if p.PlayabilityStatus.Reason != "" {
      buf.WriteString("\nReason: ")
      buf.WriteString(p.PlayabilityStatus.Reason)
   }
   return buf.String()
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

type StreamingData struct {
   AdaptiveFormats []Format
}

func (s StreamingData) Len() int {
   return len(s.AdaptiveFormats)
}

func (s StreamingData) Swap(i, j int) {
   swap := s.AdaptiveFormats[i]
   s.AdaptiveFormats[i] = s.AdaptiveFormats[j]
   s.AdaptiveFormats[j] = swap
}

type ThirdParty struct {
   EmbedURL string `json:"embedUrl"`
}
