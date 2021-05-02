package youtube

import (
   "bytes"
   "encoding/json"
   "fmt"
   "net/http"
)

func newVideo(id string) (Video, error) {
   req, err := http.NewRequest(http.MethodGet, API, nil)
   if err != nil {
      return Video{}, err
   }
   val := req.URL.Query()
   val.Set("video_id", id)
   req.URL.RawQuery = val.Encode()
   req.Header.Set("Range", "bytes=0-")
   res, err := new(http.Client).Do(req)
   if err != nil {
      return Video{}, err
   }
   defer res.Body.Close()
   switch res.StatusCode {
   case http.StatusOK, http.StatusPartialContent:
   default:
      return Video{}, fmt.Errorf("StatusCode %v", res.StatusCode)
   }
   buf := new(bytes.Buffer)
   buf.ReadFrom(res.Body)
   req.URL.RawQuery = buf.String()
   play := req.URL.Query().Get("player_response")
   buf = bytes.NewBufferString(play)
   var vid Video
   err = json.NewDecoder(buf).Decode(&vid)
   if err != nil {
      return Video{}, err
   }
   return vid, nil
}
