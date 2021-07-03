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

type Android struct {
   PlayabilityStatus struct {
      Reason string
   }
   StreamingData `json:"streamingData"`
   VideoDetails `json:"videoDetails"`
}

func NewAndroid(id string) (*Android, error) {
   res, err := ClientAndroid.newPlayer(id).post()
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   a := new(Android)
   if err := json.NewDecoder(res.Body).Decode(a); err != nil {
      return nil, err
   }
   return a, nil
}

type Client struct {
   ClientName string `json:"clientName"`
   ClientVersion string `json:"clientVersion"`
}

func (c Client) newPlayer(id string) player {
   var p player
   p.Client = c
   p.VideoID = id
   return p
}

type Context struct {
   Client `json:"client"`
}

type MWeb struct {
   Microformat `json:"microformat"`
   VideoDetails `json:"videoDetails"`
}

func NewMWeb(id string) (*MWeb, error) {
   res, err := ClientMWeb.newPlayer(id).post()
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   mw := new(MWeb)
   if err := json.NewDecoder(res.Body).Decode(mw); err != nil {
      return nil, err
   }
   return mw, nil
}

type Microformat struct {
   PlayerMicroformatRenderer `json:"playerMicroformatRenderer"`
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

type player struct {
   Context `json:"context"`
   VideoID string `json:"videoId"`
}

func (p player) post() (*http.Response, error) {
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
