package youtube

import (
   "bytes"
   "encoding/json"
   "fmt"
   "net/http"
)

const origin = "https://www.youtube.com"

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

var (
   Android = Client{"ANDROID", "16.05"}
   Mweb = Client{"MWEB", "2.19700101"}
)

type PlayerMicroformatRenderer struct {
   PublishDate string
}

type StreamingData struct {
   AdaptiveFormats FormatSlice
}

type VideoDetails struct {
   Author string
   ShortDescription string
   Title string
   VideoID string
   ViewCount int `json:"viewCount,string"`
}

type Client struct {
   ClientName string `json:"clientName"`
   ClientVersion string `json:"clientVersion"`
}

type youTubeI struct {
   Context struct {
      Client Client `json:"client"`
   } `json:"context"`
   Query string `json:"query"`
   VideoID string `json:"videoId"`
}

func post(url string, body youTubeI) (*http.Response, error) {
   buf := new(bytes.Buffer)
   if err := json.NewEncoder(buf).Encode(body); err != nil {
      return nil, err
   }
   req, err := http.NewRequest("POST", url, buf)
   if err != nil {
      return nil, err
   }
   req.Header.Set("X-Goog-Api-Key", "AIzaSyAO_FJ2SlqU8Q4STEHLGCilw_Y9_11qcW8")
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

func NewPlayer(id string, c Client) (*Player, error) {
   var body youTubeI
   body.Context.Client = c
   body.VideoID = id
   res, err := post(origin + "/youtubei/v1/player", body)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   p := new(Player)
   if err := json.NewDecoder(res.Body).Decode(p); err != nil {
      return nil, err
   }
   return p, nil
}
