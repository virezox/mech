package youtube

import (
   "encoding/json"
   "github.com/89z/mech"
   "net/http"
)

type Client struct {
   Name string `json:"clientName"`
   Screen string `json:"clientScreen,omitempty"`
   Version string `json:"clientVersion"`
}

var Android = Client{Name: "ANDROID", Version: "17.11.34"}

var Mweb = Client{Name: "MWEB", Version: "2.20211109.01.00"}

// HtVdAasjOgU
var Embed = Client{Name: "ANDROID_EMBEDDED_PLAYER", Version: "17.11.34"}

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
