package youtube

import (
   "encoding/json"
   "github.com/89z/mech"
   "net/http"
)

type Client struct {
   Name string
   MinVersion string
   MaxVersion string
}

var Android = Client{Name: "ANDROID", MaxVersion: "17.11.34"}

var Mweb = Client{Name: "MWEB", MaxVersion: "2.20220322.05.00"}

// HtVdAasjOgU
var Embed = Client{Name: "ANDROID_EMBEDDED_PLAYER", MaxVersion: "17.11.34"}

func (c Client) Player(id string) (*Player, error) {
   return c.PlayerHeader(googAPI, id)
}

type errorString string

func (e errorString) Error() string {
   return string(e)
}

type context struct {
   Client struct {
      ClientName string `json:"clientName"`
      ClientVersion string `json:"clientVersion"`
   } `json:"client"`
}

func (c Client) PlayerHeader(head http.Header, id string) (*Player, error) {
   var body struct {
      Context context `json:"context"`
      RacyCheckOK bool `json:"racyCheckOk,omitempty"`
      VideoID string `json:"videoId"`
   }
   body.Context.Client.ClientName = c.Name
   body.Context.Client.ClientVersion = c.MaxVersion
   if head.Get("Authorization") != "" {
      body.RacyCheckOK = true
   }
   body.VideoID = id
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

func (c Client) Search(query string) (*Search, error) {
   var body struct {
      Context context `json:"context"`
      Params string `json:"params"`
      Query string `json:"query"`
   }
   body.Context.Client.ClientName = c.Name
   body.Context.Client.ClientVersion = c.MaxVersion
   filter := NewFilter()
   filter.Type(Type["Video"])
   param := NewParams()
   param.Filter(filter)
   body.Params = param.Encode()
   body.Query = query
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
