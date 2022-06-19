package youtube

import (
   "bytes"
   "encoding/json"
   "errors"
   "github.com/89z/mech"
   "net/http"
)

const goog_api = "AIzaSyAO_FJ2SlqU8Q4STEHLGCilw_Y9_11qcW8"

func (y YouTubeI) Exchange(id string, ex *Exchange) (*Player, error) {
   y.VideoId = id
   buf, err := mech.Encode(y)
   if err != nil {
      return nil, err
   }
   req, err := http.NewRequest("POST", origin + "/youtubei/v1/player", buf)
   if err != nil {
      return nil, err
   }
   if ex != nil {
      req.Header.Set("Authorization", "Bearer " + ex.Access_Token)
   } else {
      req.Header.Set("X-Goog-Api-Key", goog_api)
   }
   LogLevel.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      return nil, errors.New(res.Status)
   }
   play := new(Player)
   if err := json.NewDecoder(res.Body).Decode(play); err != nil {
      return nil, err
   }
   return play, nil
}

func (y YouTubeI) Player(id string) (*Player, error) {
   return y.Exchange(id, nil)
}
type YouTubeI struct {
   ContentCheckOk bool `json:"contentCheckOk,omitempty"`
   Context Context `json:"context"`
   Query string `json:"query,omitempty"`
   RacyCheckOk bool `json:"racyCheckOk,omitempty"`
   VideoId string `json:"videoId,omitempty"`
   Params []byte `json:"params,omitempty"`
}

func (y YouTubeI) Search(query string) (*Search, error) {
   y.Query = query
   filter := NewFilter()
   filter.Type(Type["Video"])
   param := NewParams()
   param.Filter(filter)
   var err error
   y.Params, err = param.MarshalBinary()
   if err != nil {
      return nil, err
   }
   buf := new(bytes.Buffer)
   if err := json.NewEncoder(buf).Encode(y); err != nil {
      return nil, err
   }
   req, err := http.NewRequest("POST", origin + "/youtubei/v1/search", buf)
   if err != nil {
      return nil, err
   }
   req.Header.Set("X-Goog-Api-Key", goog_api)
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
   ClientName string `json:"clientName"`
   ClientVersion string `json:"clientVersion"`
}

type Context struct {
   Client Client `json:"client"`
}

const androidVersion = "17.23.35"

// 1
var Android = YouTubeI{
   Context: Context{
      Client: Client{"ANDROID", androidVersion},
   },
}

// 2
var AndroidEmbed = YouTubeI{
   Context: Context{
      Client: Client{"ANDROID_EMBEDDED_PLAYER", androidVersion},
   },
}

// 3
var AndroidRacy = YouTubeI{
   Context: Context{
      Client: Client{"ANDROID", androidVersion},
   },
   RacyCheckOk: true,
}

// 4
var AndroidContent = YouTubeI{
   Context: Context{
      Client: Client{"ANDROID", androidVersion},
   },
   RacyCheckOk: true,
   ContentCheckOk: true,
}

var Mweb = YouTubeI{
   Context: Context{
      Client: Client{"MWEB", "2.20220322.05.00"},
   },
}

