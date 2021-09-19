package reddit

import (
   "encoding/json"
   "encoding/xml"
   "fmt"
   "net/http"
   "net/url"
)

const Origin = "https://api.reddit.com"

var Verbose bool

// redd.it/ppbsh
// redd.it/pql06n
func ValidID(id string) error {
   switch len(id) {
   case 5, 6:
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

type Link struct {
   Media struct {
      Reddit_Video struct {
         DASH_URL string
      }
   }
   Subreddit string
   Title string
   URL string // https://v.redd.it/pjn0j2z4v6o71
}

func (l Link) MPD() (*MPD, error) {
   addr, err := url.Parse(l.Media.Reddit_Video.DASH_URL)
   if err != nil {
      return nil, err
   }
   addr.RawQuery = ""
   if Verbose {
      fmt.Println("GET", addr)
   }
   r, err := http.Get(addr.String())
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

type MPD struct {
   Period struct {
      AdaptationSet []Adaptation
   }
}

type Post struct {
   Data struct {
      Children []struct {
         Kind string
         Data json.RawMessage
      }
   }
}

func NewPost(id string) (*Post, error) {
   req, err := http.NewRequest("GET", Origin + "/by_id/t3_" + id, nil)
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

func (p Post) Link() (*Link, error) {
   for _, child := range p.Data.Children {
      if child.Kind == "t3" {
         lin := new(Link)
         if err := json.Unmarshal(child.Data, lin); err != nil {
            return nil, err
         }
         return lin, nil
      }
   }
   return nil, fmt.Errorf("children %v", p.Data.Children)
}
