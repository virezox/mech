package youtube

import (
   "bytes"
   "encoding/json"
   "errors"
   "github.com/89z/mech"
   "net/http"
)

const goog_API = "AIzaSyAO_FJ2SlqU8Q4STEHLGCilw_Y9_11qcW8"

func (c Config) Exchange(id string, ex *Exchange) (*Player, error) {
   c.Video_ID = id
   buf, err := mech.Encode(c)
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
      req.Header.Set("X-Goog-Api-Key", goog_API)
   }
   Log.Dump(req)
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

func (c Config) Player(id string) (*Player, error) {
   return c.Exchange(id, nil)
}

type Config struct {
   Content_Check_OK bool `json:"contentCheckOk,omitempty"`
   Context Context `json:"context"`
   Query string `json:"query,omitempty"`
   Racy_Check_OK bool `json:"racyCheckOk,omitempty"`
   Video_ID string `json:"videoId,omitempty"`
   Params []byte `json:"params,omitempty"`
}

func (c Config) Search(query string) (*Search, error) {
   c.Query = query
   filter := New_Filter()
   filter.Type(Type["Video"])
   param := New_Params()
   param.Filter(filter)
   c.Params = param.Marshal()
   buf := new(bytes.Buffer)
   if err := json.NewEncoder(buf).Encode(c); err != nil {
      return nil, err
   }
   req, err := http.NewRequest("POST", origin + "/youtubei/v1/search", buf)
   if err != nil {
      return nil, err
   }
   req.Header.Set("X-Goog-Api-Key", goog_API)
   Log.Dump(req)
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
   Client_Name string `json:"clientName"`
   Client_Version string `json:"clientVersion"`
}

type Context struct {
   Client Client `json:"client"`
}

const android_version = "17.23.35"

// 1
var Android = Config{
   Context: Context{
      Client: Client{"ANDROID", android_version},
   },
}

// 2
var Android_Embed = Config{
   Context: Context{
      Client: Client{"ANDROID_EMBEDDED_PLAYER", android_version},
   },
}

// 3
var Android_Racy = Config{
   Context: Context{
      Client: Client{"ANDROID", android_version},
   },
   Racy_Check_OK: true,
}

// 4
var Android_Content = Config{
   Context: Context{
      Client: Client{"ANDROID", android_version},
   },
   Racy_Check_OK: true,
   Content_Check_OK: true,
}

var Mweb = Config{
   Context: Context{
      Client: Client{"MWEB", "2.20220322.05.00"},
   },
}

