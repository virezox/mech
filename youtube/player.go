package youtube

import (
   "bytes"
   "encoding/json"
   "fmt"
   "io"
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

func GetVideoInfo(id string, detailPage bool) (*Player, error) {
   req, err := http.NewRequest("GET", origin + "/get_video_info", nil)
   if err != nil {
      return nil, err
   }
   q := req.URL.Query()
   q.Set("c", "ANDROID")
   q.Set("cver", "16.05")
   q.Set("eurl", origin)
   q.Set("html5", "1")
   q.Set("video_id", id)
   if detailPage {
      q.Set("el", "detailpage")
   }
   req.URL.RawQuery = q.Encode()
   fmt.Println(invert, req.Method, reset, req.URL)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      return nil, fmt.Errorf("status %v", res.Status)
   }
   body, err := io.ReadAll(res.Body)
   if err != nil {
      return nil, err
   }
   req.URL.RawQuery = string(body)
   play := req.URL.Query().Get("player_response")
   p := new(Player)
   if err := json.Unmarshal([]byte(play), p); err != nil {
      return nil, err
   }
   return p, nil
}

func IPlayer(id string) (*Player, error) {
   i := newYouTubeI()
   i.VideoID = id
   res, err := i.post("/youtubei/v1/player")
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

type youTubeI struct {
   Context struct {
      Client struct {
         ClientName string `json:"clientName"`
         ClientVersion string `json:"clientVersion"`
      } `json:"client"`
   } `json:"context"`
   Query string `json:"query"`
   VideoID string `json:"videoId"`
}

func newYouTubeI() youTubeI {
   var i youTubeI
   i.Context.Client.ClientName = "MWEB"
   i.Context.Client.ClientVersion = "2.19700101"
   return i
}

func (i youTubeI) post(path string) (*http.Response, error) {
   buf := new(bytes.Buffer)
   if err := json.NewEncoder(buf).Encode(i); err != nil {
      return nil, err
   }
   req, err := http.NewRequest("POST", origin + path, buf)
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
