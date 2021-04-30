package youtube

import (
   "bytes"
   "encoding/json"
   "errors"
   "fmt"
   "io"
   "net/http"
   "net/url"
   "regexp"
   "strings"
)

// extractVideoID extracts the videoID from the given string
func extractVideoID(videoID string) (string, error) {
   videoRegexpList := []string{
      0: `(?:v|embed|watch\?v)(?:=|/)([^"&?/=%]{11})`,
      1: `(?:=|/)([^"&?/=%]{11})`,
      2: `([^"&?/=%]{11})`,
   }
   if strings.Contains(videoID, "youtu") || strings.ContainsAny(videoID, "\"?&/<%=") {
      for _, pat := range videoRegexpList {
         re := regexp.MustCompile(pat)
         if isMatch := re.MatchString(videoID); isMatch {
            subs := re.FindStringSubmatch(videoID)
            videoID = subs[1]
         }
      }
   }
   if strings.ContainsAny(videoID, "?&/<%=") {
      return "", errors.New("invalid characters in video id")
   }
   if len(videoID) < 10 {
      return "", errors.New("the video id must be at least 10 characters long")
   }
   return videoID, nil
}

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

// NewVideo fetches video metadata
func NewVideo(addr string) (*Video, error) {
   id, err := extractVideoID(addr)
   if err != nil {
      return nil, fmt.Errorf("extractVideoID failed: %w", err)
   }
   eurl := "https://youtube.googleapis.com/v/" + id
   body, err := httpGetBodyBytes("https://youtube.com/get_video_info?video_id="+id+"&eurl="+eurl)
   if err != nil { return nil, err }
   query, err := url.ParseQuery(string(body))
   if err != nil { return nil, err }
   status := query.Get("status")
   if status != "ok" {
      return nil, fmt.Errorf(
         "response status: %q, reason: %q", status, query.Get("reason"),
      )
   }
   playerResponse := query.Get("player_response")
   if playerResponse == "" {
      return nil, errors.New("no player_response found in the server's answer")
   }
   var prData Video
   err = json.Unmarshal([]byte(playerResponse), &prData)
   if err != nil {
      return nil, fmt.Errorf("unable to parse player response JSON %v", err)
   }
   return &prData, nil
}

func oldPlayer(id string) (player, error) {
   api := "https://www.youtube.com/get_video_info"
   req, err := http.NewRequest("GET", api, nil)
   if err != nil {
      return player{}, err
   }
   val := req.URL.Query()
   val.Set("video_id", id)
   req.URL.RawQuery = val.Encode()
   res, err := new(http.Client).Do(req)
   if err != nil {
      return player{}, err
   }
   buf := new(bytes.Buffer)
   buf.ReadFrom(res.Body)
   req.URL.RawQuery = buf.String()
   play := req.URL.Query().Get("player_response")
   buf = bytes.NewBufferString(play)
   var video player
   json.NewDecoder(buf).Decode(&video)
   return video, nil
}
