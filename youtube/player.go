package youtube

import (
   "bytes"
   "encoding/json"
   "fmt"
   "net/http"
)

var (
   clientAndroid = client{"ANDROID", "15.01"}
   clientMweb = client{"MWEB", "2.19700101"}
)

type Microformat struct {
   PlayerMicroformatRenderer `json:"playerMicroformatRenderer"`
}

type Player struct {
   Microformat `json:"microformat"`
   PlayabilityStatus struct {
      Reason string
   }
   StreamingData `json:"streamingData"`
   VideoDetails `json:"videoDetails"`
}

func PlayerAndroid(id string) (*Player, error) {
   res, err := clientAndroid.video(id).post("/youtubei/v1/player")
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

func PlayerMweb(id string) (*Player, error) {
   res, err := clientMweb.video(id).post("/youtubei/v1/player")
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

type PlayerMicroformatRenderer struct {
   AvailableCountries []string
   PublishDate string
}

type StreamingData struct {
   AdaptiveFormats Formats
}

type VideoDetails struct {
   Author string
   ShortDescription string
   Title string
   ViewCount int `json:"viewCount,string"`
}

type client struct {
   ClientName string `json:"clientName"`
   ClientVersion string `json:"clientVersion"`
}

func (c client) query(s string) request {
   var r request
   r.Context.Client = c
   r.Query = s
   return r
}

func (c client) video(id string) request {
   var r request
   r.Context.Client = c
   r.VideoID = id
   return r
}

type request struct {
   Context struct {
      Client client `json:"client"`
   } `json:"context"`
   Query string `json:"query"`
   VideoID string `json:"videoId"`
}

func (r request) post(path string) (*http.Response, error) {
   buf := new(bytes.Buffer)
   err := json.NewEncoder(buf).Encode(r)
   if err != nil {
      return nil, err
   }
   req, err := http.NewRequest("POST", "https://www.youtube.com" + path, buf)
   if err != nil {
      return nil, err
   }
   val := req.URL.Query()
   val.Set("key", "AIzaSyAO_FJ2SlqU8Q4STEHLGCilw_Y9_11qcW8")
   req.URL.RawQuery = val.Encode()
   fmt.Println(invert, req.Method, reset, req.URL)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   if res.StatusCode != http.StatusOK {
      return nil, fmt.Errorf("status %v", res.Status)
   }
   return res, nil
}
