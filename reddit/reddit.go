package reddit

import (
   "encoding/json"
   "encoding/xml"
   "fmt"
   "html"
   "mime"
   "net/http"
)

var Verbose bool

// redd.it/pql06n
func ValidID(id string) error {
   if len(id) == 6 {
      return nil
   }
   return fmt.Errorf("%q invalid as ID", id)
}

type Adaptation struct {
   MimeType string `xml:"mimeType,attr"`
   Representation []struct {
      BaseURL string
      Height int `xml:"height,attr"`
   }
}

func (a Adaptation) Ext() string {
   exts, err := mime.ExtensionsByType(a.MimeType)
   if err != nil {
      return ""
   }
   return exts[0]
}

type MPD struct {
   Period struct {
      AdaptationSet []Adaptation
   }
}

type Post []struct {
   Data struct {
      Children []struct {
         Kind string
         Data json.RawMessage
      }
   }
}

func NewPost(id string) (*Post, error) {
   req, err := http.NewRequest(
      "GET", "https://www.reddit.com/comments/" + id + ".json", nil,
   )
   if err != nil {
      return nil, err
   }
   req.Header.Set("User-Agent", "Mozilla")
   if Verbose {
      fmt.Println(req.Method, req.URL)
   }
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      return nil, fmt.Errorf("status %q", res.Status)
   }
   pos := new(Post)
   if err := json.NewDecoder(res.Body).Decode(pos); err != nil {
      return nil, err
   }
   return pos, nil
}

func (p Post) T3() (*T3, error) {
   var kinds []string
   for _, list := range p {
      for _, child := range list.Data.Children {
         if child.Kind == "t3" {
            t3 := new(T3)
            if err := json.Unmarshal(child.Data, t3); err != nil {
               return nil, err
            }
            return t3, nil
         }
         kinds = append(kinds, child.Kind)
      }
   }
   return nil, fmt.Errorf("kinds %v", kinds)
}

type T3 struct {
   Media struct {
      Reddit_Video struct {
         DASH_URL string
      }
   }
   Subreddit string
   Title string
   URL string // https://v.redd.it/pjn0j2z4v6o71
}

func (t T3) MPD() (*MPD, error) {
   addr := html.UnescapeString(t.Media.Reddit_Video.DASH_URL)
   if Verbose {
      fmt.Println("GET", addr)
   }
   r, err := http.Get(addr)
   if err != nil {
      return nil, err
   }
   defer r.Body.Close()
   media := new(MPD)
   if err := xml.NewDecoder(r.Body).Decode(media); err != nil {
      return nil, err
   }
   return media, nil
}
