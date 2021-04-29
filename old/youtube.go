// YouTube
package youtube

import (
   "bytes"
   "encoding/json"
   "fmt"
   "github.com/89z/x"
   "math"
   "net/http"
   "strconv"
   "time"
)

const API = "https://www.youtube.com/get_video_info"

func numberFormat(d float64) string {
   e := int(math.Log10(d)) / 3
   f := d / math.Pow10(e * 3)
   return fmt.Sprintf("%.3f", f) + []string{"", " k", " M", " G"}[e]
}

func sinceHours(left string) (float64, error) {
   right := "1970-01-01T00:00:00Z"[len(left):]
   t, err := time.Parse(time.RFC3339, left + right)
   if err != nil { return 0, err }
   return time.Since(t).Hours(), nil
}

func Info(id string) (Player, error) {
   req, err := http.NewRequest("GET", API, nil)
   if err != nil {
      return Player{}, err
   }
   val := req.URL.Query()
   val.Set("video_id", id)
   req.URL.RawQuery = val.Encode()
   x.LogInfo("GET", req.URL)
   res, err := new(http.Client).Do(req)
   if err != nil {
      return Player{}, err
   }
   buf := new(bytes.Buffer)
   buf.ReadFrom(res.Body)
   req.URL.RawQuery = buf.String()
   play := req.URL.Query().Get("player_response")
   buf = bytes.NewBufferString(play)
   var video Player
   json.NewDecoder(buf).Decode(&video)
   return video, nil
}

type Player struct {
   Microformat struct {
      PlayerMicroformatRenderer struct {
         PublishDate string
         ViewCount string
         Description struct {
            SimpleText string
         }
         Title struct {
            SimpleText string
         }
      }
   }
}
