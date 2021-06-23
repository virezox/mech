package youtube

import (
   "bytes"
   "encoding/json"
   "fmt"
   "net/http"
)

type Request struct {
   Context struct {
      Client struct {
         ClientName string `json:"clientName"`
         ClientVersion string `json:"clientVersion"`
      } `json:"client"`
   } `json:"context"`
   Query string `json:"query"`
   VideoID string `json:"videoId"`
}

func QueryRequest(query string) Request {
   var r Request
   r.Context.Client.ClientName = "WEB"
   r.Context.Client.ClientVersion = VersionWeb
   r.Query = query
   return r
}

func VideoRequest(id, name, version string) Request {
   var r Request
   r.Context.Client.ClientName = name
   r.Context.Client.ClientVersion = version
   r.VideoID = id
   return r
}

func (r Request) post() (*http.Response, error) {
   buf := new(bytes.Buffer)
   json.NewEncoder(buf).Encode(r)
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
