package youtube

import (
   "bytes"
   "encoding/json"
   "fmt"
   "net/http"
)

const PlayerAPI = "https://www.youtube.com/youtubei/v1/player"

type Format struct {
   Bitrate int64
   ContentLength int64 `json:"contentLength,string"`
   Height int
   Itag int
   MimeType string
   SignatureCipher string
   URL string
}

type Player struct {
   StreamingData struct {
      AdaptiveFormats []Format
   }
   Microformat struct {
      PlayerMicroformatRenderer struct {
         AvailableCountries []string
         PublishDate string
      }
   }
   VideoDetails struct {
      Author string
      ShortDescription string
      Title string
      ViewCount int `json:"viewCount,string"`
   }
}

func NewPlayer(id string) (Player, error) {
   body := m{
      "videoId": id, "context": m{
         "client": m{"clientName": "WEB", "clientVersion": "1.19700101"},
      },
   }
   buf := new(bytes.Buffer)
   json.NewEncoder(buf).Encode(body)
   req, err := http.NewRequest("POST", PlayerAPI, buf)
   if err != nil {
      return Player{}, err
   }
   val := req.URL.Query()
   val.Set("key", "AIzaSyAO_FJ2SlqU8Q4STEHLGCilw_Y9_11qcW8")
   req.URL.RawQuery = val.Encode()
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return Player{}, err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      return Player{}, fmt.Errorf("status %v", res.Status)
   }
   var play Player
   json.NewDecoder(res.Body).Decode(&play)
   return play, nil
}

type m map[string]interface{}
