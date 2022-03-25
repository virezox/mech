package youtube

import (
   "encoding/json"
   "fmt"
   "github.com/89z/mech"
   "mime"
   "net/http"
   "net/url"
   "path"
   "strings"
   "time"
)

var Android = Client{Name: "ANDROID", Version: "17.11.34"}

var Mweb = Client{Name: "MWEB", Version: "2.20211109.01.00"}

// HtVdAasjOgU
var Embed = Client{Name: "ANDROID_EMBEDDED_PLAYER", Version: "17.09.33"}

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

type Client struct {
   Name string `json:"clientName"`
   Screen string `json:"clientScreen,omitempty"`
   Version string `json:"clientVersion"`
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

func (p Player) Base() string {
   return mech.Clean(p.VideoDetails.Author + "-" + p.VideoDetails.Title)
}

func (p Player) Date() (time.Time, error) {
   value := p.Microformat.PlayerMicroformatRenderer.PublishDate
   return time.Parse("2006-01-02", value)
}

func (p Player) Format(f fmt.State, verb rune) {
   fmt.Fprintln(f, p.Status())
   fmt.Fprintln(f, "VideoID:", p.VideoDetails.VideoID)
   fmt.Fprintln(f, "Length:", p.VideoDetails.LengthSeconds)
   fmt.Fprintln(f, "ViewCount:", p.VideoDetails.ViewCount)
   fmt.Fprintln(f, "Author:", p.VideoDetails.Author)
   fmt.Fprintln(f, "Title:", p.VideoDetails.Title)
   date := p.Microformat.PlayerMicroformatRenderer.PublishDate
   if date != "" {
      fmt.Fprintln(f, "Date:", date)
   }
   for _, form := range p.StreamingData.AdaptiveFormats {
      fmt.Fprintln(f)
      form.Format(f, verb)
   }
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

func (f Formats) Len() int {
   return len(f)
}

type Player struct {
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
   Microformat struct {
      PlayerMicroformatRenderer struct {
         PublishDate string // 2013-06-11
      }
   }
   StreamingData struct {
      AdaptiveFormats Formats
      Formats Formats
   }
}

type Formats []Format

func (f Formats) MediaType() error {
   for i, form := range f {
      typ, param, err := mime.ParseMediaType(form.MimeType)
      if err != nil {
         return err
      }
      param["codecs"], _, _ = strings.Cut(param["codecs"], ".")
      f[i].MimeType = mime.FormatMediaType(typ, param)
   }
   return nil
}

func (f Formats) Swap(i, j int) {
   f[i], f[j] = f[j], f[i]
}

func (c Client) Player(id string) (*Player, error) {
   return c.PlayerHeader(googAPI, id)
}

func (c Client) PlayerHeader(head http.Header, id string) (*Player, error) {
   var body struct {
      RacyCheckOK bool `json:"racyCheckOk,omitempty"`
      VideoID string `json:"videoId"`
      Context struct {
         Client Client `json:"client"`
      } `json:"context"`
   }
   body.VideoID = id
   if head.Get("Authorization") != "" {
      body.RacyCheckOK = true // Cr381pDsSsA
   }
   body.Context.Client = c
   buf, err := mech.Encode(body)
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

func (c Client) Search(query string) (*Search, error) {
   var body struct {
      Params string `json:"params"`
      Query string `json:"query"`
      Context struct {
         Client Client `json:"client"`
      } `json:"context"`
   }
   filter := NewFilter()
   filter.Type(Type["Video"])
   param := NewParams()
   param.Filter(filter)
   body.Params = param.Encode()
   body.Query = query
   body.Context.Client = c
   buf, err := mech.Encode(body)
   if err != nil {
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
