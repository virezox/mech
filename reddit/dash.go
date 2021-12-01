package reddit

import (
   "encoding/json"
   "encoding/xml"
   "github.com/89z/mech"
   "html"
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

type DASH struct {
   Period struct {
      AdaptationSet []struct {
         // sometimes MimeType is here, for example `pqoz44`
         MimeType string `xml:"mimeType,attr"`
         Representation []struct {
            ID string `xml:"id,attr"`
            // sometimes MimeType is here, for example `fffrnw`
            MimeType string `xml:"mimeType,attr"`
            BaseURL string
         }
      }
   }
}

type Link struct {
   Media struct {
      Reddit_Video struct {
         DASH_URL Text // v.redd.it/16cqbkev2ci51/DASHPlaylist.mpd
         HLS_URL string // v.redd.it/16cqbkev2ci51/HLSPlaylist.m3u8
      }
   }
   Subreddit string
   Title string
   URL string // v.redd.it/pjn0j2z4v6o71
}

func (l Link) DASH() (*DASH, error) {
   addr := l.Media.Reddit_Video.DASH_URL.String()
   req, err := http.NewRequest("GET", addr, nil)
   if err != nil {
      return nil, err
   }
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   dash := new(DASH)
   if err := xml.NewDecoder(res.Body).Decode(dash); err != nil {
      return nil, err
   }
   prefix, _ := path.Split(addr)
   for aKey, aVal := range dash.Period.AdaptationSet {
      for rKey, rVal := range aVal.Representation {
         rVal.BaseURL = prefix + rVal.BaseURL
         if rVal.MimeType == "" {
            rVal.MimeType = aVal.MimeType
         }
         dash.Period.AdaptationSet[aKey].Representation[rKey] = rVal
      }
   }
   return dash, nil
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
   res, err := new(http.Transport).RoundTrip(req)
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
   return nil, mech.NotFound{"t3"}
}

type Text string

func (t Text) String() string {
   str := string(t)
   return html.UnescapeString(str)
}
