package reddit

import (
   "encoding/json"
   "encoding/xml"
   "fmt"
   "github.com/89z/mech"
   "net/http"
   "path"
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

type DASH struct {
   Period struct {
      AdaptationSet []struct {
         // sometimes MimeType is here, for example `pqoz44`
         MimeType string `xml:"mimeType,attr"`
         Representation []struct {
            BaseURL string
            Height int `xml:"height,attr"`
            // sometimes MimeType is here, for example `fffrnw`
            MimeType string `xml:"mimeType,attr"`
         }
      }
   }
}

func (l Link) DASH() (*DASH, error) {
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
   dash := new(DASH)
   if err := xml.NewDecoder(res.Body).Decode(dash); err != nil {
      return nil, err
   }
   prefix, _ := path.Split(l.Media.Reddit_Video.DASH_URL)
   for aKey, aVal := range dash.Period.AdaptationSet {
      for rKey, rVal := range aVal.Representation {
         rVal.BaseURL = prefix + rVal.BaseURL
         dash.Period.AdaptationSet[aKey].Representation[rKey] = rVal
      }
   }
   return dash, nil
}
