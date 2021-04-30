package youtube

import (
   "encoding/json"
   "errors"
   "fmt"
   "io"
   "net/http"
   "net/url"
)

// httpGetBodyBytes reads the whole HTTP body and returns it
func httpGetBodyBytes(addr string) ([]byte, error) {
   req, err := http.NewRequest(http.MethodGet, addr, nil)
   if err != nil { return nil, err }
   req.Header.Set("Range", "bytes=0-")
   res, err := new(http.Client).Do(req)
   if err != nil { return nil, err }
   defer res.Body.Close()
   switch res.StatusCode {
   case http.StatusOK, http.StatusPartialContent:
   default:
      return nil, fmt.Errorf("unexpected status code: %v", res.StatusCode)
   }
   return io.ReadAll(res.Body)
}

type Video struct {
   Microformat struct {
      PlayerMicroformatRenderer struct {
         Description struct {
            SimpleText string
         }
         PublishDate        string
         Title struct {
            SimpleText string
         }
         ViewCount          string
      }
   }
   StreamingData struct {
      DashManifestURL  string
   }
}

// NewVideo fetches video metadata
func NewVideo(id string) (Video, error) {
   eurl := "https://youtube.googleapis.com/v/" + id
   body, err := httpGetBodyBytes("https://youtube.com/get_video_info?video_id="+id+"&eurl="+eurl)
   if err != nil {
      return Video{}, err
   }
   query, err := url.ParseQuery(string(body))
   if err != nil {
      return Video{}, err
   }
   status := query.Get("status")
   if status != "ok" {
      return Video{}, fmt.Errorf(
         "response status: %q, reason: %q", status, query.Get("reason"),
      )
   }
   playerResponse := query.Get("player_response")
   if playerResponse == "" {
      return Video{}, errors.New(
         "no player_response found in the server's answer",
      )
   }
   var prData Video
   err = json.Unmarshal([]byte(playerResponse), &prData)
   if err != nil {
      return Video{}, fmt.Errorf("unable to parse player response JSON %v", err)
   }
   return prData, nil
}
