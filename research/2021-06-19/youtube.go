package youtube

import (
   "encoding/json"
   "fmt"
   "net/http"
   "strings"
)

const PlayerAPI = "https://www.youtube.com/youtubei/v1/player"

func NewPlayer(id string) (Player, error) {
   body := fmt.Sprintf(`
   {
      "videoId": "%v", "context": {
         "client": {"clientName": "WEB", "clientVersion": "1.19700101"}
      }
   }
   `, id)
   req, err := http.NewRequest("POST", PlayerAPI, strings.NewReader(body))
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
