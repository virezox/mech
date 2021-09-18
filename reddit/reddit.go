package reddit

import (
   "encoding/json"
   "encoding/xml"
   "fmt"
   "html"
   "net/http"
)

var Verbose bool

type MPD struct {
   Period struct {
      AdaptationSet []struct {
         Representation []struct {
            BaseURL string
         }
      }
   }
}

type Post []struct {
   Data struct {
      Children []struct {
         Data struct {
            Media *struct {
               Reddit_Video struct {
                  DASH_URL string
               }
            }
            URL string
         }
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

func (p Post) MPD() (*MPD, error) {
   addr := p.dashURL()
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

func (p Post) dashURL() string {
   for _, list := range p {
      for _, c := range list.Data.Children {
         return html.UnescapeString(c.Data.Media.Reddit_Video.DASH_URL)
      }
   }
   return ""
}
