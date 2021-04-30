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
func httpGetBodyBytes(url string) ([]byte, error) {
   req, err := http.NewRequest(http.MethodGet, url, nil)
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
func NewVideo(url string) (*Video, error) {
   id, err := extractVideoID(url)
   if err != nil {
      return nil, fmt.Errorf("extractVideoID failed: %w", err)
   }
   eurl := "https://youtube.googleapis.com/v/" + id
   body, err := httpGetBodyBytes("https://youtube.com/get_video_info?video_id="+id+"&eurl="+eurl)
   if err != nil { return nil, err }
   v := new(Video)
   err = v.parseVideoInfo(body)
   if err == errors.New("embedding of this video has been disabled") {
      html, err := httpGetBodyBytes("https://www.youtube.com/watch?v="+id)
      if err != nil { return nil, err }
      return v, v.parseVideoPage(html)
   }
   return v, err
}

func (v *Video) extractDataFromPlayerResponse(prData playerResponseData) error {
   v.Title = prData.VideoDetails.Title
   v.Description = prData.VideoDetails.ShortDescription
   v.DASHManifestURL = prData.StreamingData.DashManifestURL
   return nil
}

func (v *Video) parseVideoInfo(body []byte) error {
   query, err := url.ParseQuery(string(body))
   if err != nil { return err }
   status := query.Get("status")
   if status != "ok" {
      return fmt.Errorf(
         "response status: %q, reason: %q", status, query.Get("reason"),
      )
   }
   playerResponse := query.Get("player_response")
   if playerResponse == "" {
      return errors.New("no player_response found in the server's answer")
   }
   var prData playerResponseData
   if err := json.Unmarshal([]byte(playerResponse), &prData); err != nil {
      return fmt.Errorf("unable to parse player response JSON: %w", err)
   }
   return v.extractDataFromPlayerResponse(prData)
}

func (v *Video) parseVideoPage(body []byte) error {
   re := `var ytInitialPlayerResponse\s*=\s*(\{.+?\});`
   playerResponse := regexp.MustCompile(re).FindSubmatch(body)
   if playerResponse == nil || len(playerResponse) < 2 {
      return errors.New("no ytInitialPlayerResponse found in the server's answer")
   }
   var prData playerResponseData
   if err := json.Unmarshal(playerResponse[1], &prData); err != nil {
      return fmt.Errorf("unable to parse player response JSON: %w", err)
   }
   return v.extractDataFromPlayerResponse(prData)
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

type playerResponseData struct {
   Microformat struct {
      PlayerMicroformatRenderer struct {
         Embed struct {
            IframeURL      string
            FlashURL       string
            Width          int
            Height         int
            FlashSecureURL string
         }
         Title struct {
            SimpleText string
         }
         Description struct {
            SimpleText string
         }
         LengthSeconds      string
         OwnerProfileURL    string
         ExternalChannelID  string
         AvailableCountries []string
         IsUnlisted         bool
         HasYpcMetadata     bool
         ViewCount          string
         Category           string
         PublishDate        string
         OwnerChannelName   string
         UploadDate         string
      }
   }
   StreamingData struct {
      ExpiresInSeconds string
      DashManifestURL  string
   }
   VideoDetails struct {
      VideoID          string
      Title            string
      LengthSeconds    string
      ChannelID        string
      IsOwnerViewing   bool
      ShortDescription string
      IsCrawlable      bool
      AverageRating     float64
      AllowRatings      bool
      ViewCount         string
      Author            string
      IsPrivate         bool
      IsUnpluggedCorpus bool
      IsLiveContent     bool
   }
}
