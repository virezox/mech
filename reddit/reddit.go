package reddit

import (
   "encoding/json"
   "encoding/xml"
   "fmt"
   "github.com/89z/mech"
   "net/http"
   "sort"
)

const Origin = "https://api.reddit.com"

// redd.it/ppbsh
// redd.it/pql06n
func Valid(id string) bool {
   switch len(id) {
   case 5, 6:
      return true
   }
   return false
}

type Adaptation struct {
   MimeType string `xml:"mimeType,attr"`
   Representation Representation
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
   req, err := http.NewRequest("GET", l.Media.Reddit_Video.DASH_URL, nil)
   if err != nil {
      return nil, err
   }
   req.URL.RawQuery = ""
   res, err := mech.RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   media := new(MPD)
   if err := xml.NewDecoder(res.Body).Decode(media); err != nil {
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
   res, err := mech.RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
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

type Representation []struct {
   BaseURL string
   Height int `xml:"height,attr"`
   MimeType string `xml:"mimeType,attr"`
}

func (r Representation) Sort() {
   sort.Slice(r, func(a, b int) bool {
      return r[b].Height < r[a].Height
   })
}
