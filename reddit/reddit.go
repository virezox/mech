package reddit

import (
   "encoding/json"
   "encoding/xml"
   "fmt"
   "html"
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

type MPD struct {
   Period struct {
      AdaptationSet []struct {
         Representation []struct {
            BaseURL string
            Height int `xml:"height,attr"`
         }
      }
   }
}

type Post []struct {
   Data struct {
      Children []json.RawMessage
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
   for _, list := range p {
      for _, child := range list.Data.Children {
         t3 := new(T3)
         if err := json.Unmarshal(child, t3); err != nil {
            return nil, err
         }
         return t3, nil
      }
   }
   return nil, fmt.Errorf("post length %v", len(p))
}

type T3 struct {
   Data struct {
      Media struct {
         Reddit_Video struct {
            DASH_URL string
         }
      }
      Subreddit string
      Title string
      URL string
   }
}

func (t T3) MPD() (*MPD, error) {
   addr := html.UnescapeString(t.Data.Media.Reddit_Video.DASH_URL)
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
