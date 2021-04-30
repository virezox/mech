package youtube

import (
   "bytes"
   "encoding/json"
   "net/http"
)

func oldVideo(id string) (Video, error) {
   api := "https://www.youtube.com/get_video_info"
   req, err := http.NewRequest("GET", api, nil)
   if err != nil {
      return Video{}, err
   }
   val := req.URL.Query()
   val.Set("video_id", id)
   req.URL.RawQuery = val.Encode()
   res, err := new(http.Client).Do(req)
   if err != nil {
      return Video{}, err
   }
   buf := new(bytes.Buffer)
   buf.ReadFrom(res.Body)
   req.URL.RawQuery = buf.String()
   play := req.URL.Query().Get("player_response")
   buf = bytes.NewBufferString(play)
   var video Video
   json.NewDecoder(buf).Decode(&video)
   return video, nil
}
