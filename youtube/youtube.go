// YouTube
package youtube

import (
   "bytes"
   "encoding/json"
   "fmt"
   "net/http"
)

const chunk = 10_000_000

type Context struct {
   Client struct {
      ClientName string `json:"clientName"`
      ClientVersion string `json:"clientVersion"`
   } `json:"client"`
}

type Search struct {
   Context `json:"context"`
   Query string `json:"query"`
}

func NewSearch(query string) Search {
   var s Search
   s.Context.Client.ClientName = "WEB"
   s.Context.Client.ClientVersion = VersionWeb
   s.Query = query
   return s
}

type player struct {
   Context `json:"context"`
   VideoID string `json:"videoId"`
}

func newPlayer(id, name, version string) player {
   var p player
   p.Context.Client.ClientName = name
   p.Context.Client.ClientVersion = version
   p.VideoID = id
   return p
}


func (p player) post() (*http.Response, error) {
   buf := new(bytes.Buffer)
   json.NewEncoder(buf).Encode(p)
   req, err := http.NewRequest(
      "POST", "https://www.youtube.com/youtubei/v1/player", buf,
   )
   if err != nil {
      return nil, err
   }
   val := req.URL.Query()
   val.Set("key", "AIzaSyAO_FJ2SlqU8Q4STEHLGCilw_Y9_11qcW8")
   req.URL.RawQuery = val.Encode()
   fmt.Println(invert, "POST", reset, req.URL)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   if res.StatusCode != http.StatusOK {
      return nil, fmt.Errorf("status %v", res.Status)
   }
   return res, nil
}
