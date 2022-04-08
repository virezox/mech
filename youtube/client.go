package youtube

import (
   "encoding/json"
   "github.com/89z/mech"
   "net/http"
)

type Client struct {
   ClientName string `json:"clientName"`
   ClientVersion string `json:"clientVersion"`
}

type Context struct {
   Client Client `json:"client"`
}

type YouTubeI struct {
   ContentCheckOK bool `json:"contentCheckOk,omitempty"`
   Context Context `json:"context"`
   Params string `json:"params,omitempty"`
   Query string `json:"query,omitempty"`
   RacyCheckOK bool `json:"racyCheckOk,omitempty"`
   VideoID string `json:"videoId,omitempty"`
}

var Android = YouTubeI{
   Context: Context{
      Client: Client{"ANDROID", "17.11.34"},
   },
}

var AndroidContent = YouTubeI{
   Context: Context{
      Client: Client{"ANDROID", "17.11.34"},
   },
   RacyCheckOK: true,
   ContentCheckOK: true,
}

var AndroidEmbed = YouTubeI{
   Context: Context{
      Client: Client{"ANDROID_EMBEDDED_PLAYER", "17.11.34"},
   },
}

var AndroidRacy = YouTubeI{
   Context: Context{
      Client: Client{"ANDROID", "17.11.34"},
   },
   RacyCheckOK: true,
}

var Mweb = YouTubeI{
   Context: Context{
      Client: Client{"MWEB", "2.20220322.05.00"},
   },
}

var googAPI = http.Header{
   "X-Goog-Api-Key": {"AIzaSyAO_FJ2SlqU8Q4STEHLGCilw_Y9_11qcW8"},
}

func (y YouTubeI) Search(query string) (*Search, error) {
   y.Query = query
   filter := NewFilter()
   filter.Type(Type["Video"])
   param := NewParams()
   param.Filter(filter)
   y.Params = param.Encode()
   buf, err := mech.Encode(y)
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

type errorString string

func (e errorString) Error() string {
   return string(e)
}

func (y YouTubeI) Player(id string) (*Player, error) {
   return y.Header(googAPI, id)
}

////////////////////////////////////////////////////////////////////////////////

func (y YouTubeI) Header(head http.Header, id string) (*Player, error) {
   y.VideoID = id
   buf, err := mech.Encode(y)
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
