package youtube

import (
   "bytes"
   "encoding/json"
   "fmt"
   "net/http"
)

var (
   ClientAndroid = Client{"ANDROID", "15.01"}
   ClientMWeb = Client{"MWEB", "2.19700101"}
)

func Android(id string) (*Player, error) {
   res, err := ClientAndroid.PlayerRequest(id).post()
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   a := new(Player)
   if err := json.NewDecoder(res.Body).Decode(a); err != nil {
      return nil, err
   }
   return a, nil
}

func (c Client) PlayerRequest(id string) PlayerRequest {
   var p PlayerRequest
   p.Context.Client = c
   p.VideoID = id
   return p
}

func MWeb(id string) (*Player, error) {
   res, err := ClientMWeb.PlayerRequest(id).post()
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   mw := new(Player)
   if err := json.NewDecoder(res.Body).Decode(mw); err != nil {
      return nil, err
   }
   return mw, nil
}

func (p PlayerRequest) post() (*http.Response, error) {
   buf := new(bytes.Buffer)
   err := json.NewEncoder(buf).Encode(p)
   if err != nil {
      return nil, err
   }
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
