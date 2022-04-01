package youtube

import (
   "encoding/json"
   "github.com/89z/mech"
   "net/http"
)

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

type Client struct {
   Name string `json:"clientName"`
   Version string `json:"clientVersion"`
}

var Mweb = Client{"MWEB", "2.20220322.05.00"}

var Clients = []Client{
   Mweb,
   {"ANDROID", "17.12.34"},
   {"ANDROID_EMBEDDED_PLAYER", "17.12.34"}, // HtVdAasjOgU
   {"TVHTML5", "7.20220323.10.00"},
   {"WEB", "2.20220325.00.00"},
   {"WEB_CREATOR", "1.20220324.00.00"},
   {"WEB_EMBEDDED_PLAYER", "1.20220323.01.00"},
   {"WEB_KIDS", "2.20220323.08.00"},
   {"WEB_REMIX", "1.20220321.01.00"},
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
   if res.StatusCode != http.StatusOK {
      return nil, errorString(res.Status)
   }
   play := new(Player)
   if err := json.NewDecoder(res.Body).Decode(play); err != nil {
      return nil, err
   }
   return play, nil
}

type errorString string

func (e errorString) Error() string {
   return string(e)
}
