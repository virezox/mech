// YouTube
package youtube

import (
   "bytes"
   "encoding/json"
   "fmt"
   "net/http"
)

const (
   chunk = 10_000_000
   invert = "\x1b[7m"
   reset = "\x1b[m"
)

type Context struct {
   Client struct {
      ClientName string `json:"clientName"`
      ClientVersion string `json:"clientVersion"`
   } `json:"client"`
}


type Player struct {
   Context `json:"context"`
   VideoID string `json:"videoId"`
}

func NewPlayer(id, name, version string) Player {
   var p Player
   p.Context.Client.ClientName = name
   p.Context.Client.ClientVersion = version
   p.VideoID = id
   return p
}

func (p Player) Post() (*http.Response, error) {
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

func (s Search) Post() (*Result, error) {
   buf := new(bytes.Buffer)
   json.NewEncoder(buf).Encode(s)
   req, err := http.NewRequest(
      "POST", "https://www.youtube.com/youtubei/v1/search", buf,
   )
   if err != nil {
      return nil, err
   }
   val := req.URL.Query()
   val.Set("key", "AIzaSyDCU8hByM-4DrUqRUYnGn-3llEO78bcxq8")
   req.URL.RawQuery = val.Encode()
   fmt.Println(invert, "POST", reset, req.URL)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      return nil, fmt.Errorf("status %v", res.Status)
   }
   r := new(Result)
   json.NewDecoder(res.Body).Decode(r)
   return r, nil
}
