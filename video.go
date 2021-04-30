package youtube

import (
   "bytes"
   "encoding/json"
   "errors"
   "fmt"
   "io"
   "log"
   "net/http"
   "net/url"
   "regexp"
   "strconv"
   "strings"
   "time"
)

type Video struct {
   ID              string
   Title           string
   Description     string
   Author          string
   Duration        time.Duration
   Formats         FormatList
   Thumbnails      []Thumbnail
   DASHManifestURL string // URI of the DASH manifest file
   HLSManifestURL  string // URI of the HLS manifest file
}

// NewVideo fetches video metadata
func NewVideo(url string) (*Video, error) {
   id, err := extractVideoID(url)
   if err != nil {
      return nil, fmt.Errorf("extractVideoID failed: %w", err)
   }
   return videoFromID(id)
}

func videoFromID(id string) (*Video, error) {
   eurl := "https://youtube.googleapis.com/v/" + id
   body, err := httpGetBodyBytes("https://youtube.com/get_video_info?video_id="+id+"&eurl="+eurl)
   if err != nil { return nil, err }
   v := &Video{ID: id}
   err = v.parseVideoInfo(body)
   if err == ErrNotPlayableInEmbed {
      html, err := httpGetBodyBytes("https://www.youtube.com/watch?v="+id)
      if err != nil { return nil, err }
      return v, v.parseVideoPage(html)
   }
   return v, err
}

// httpGet does a HTTP GET request, checks the response to be a 200 OK and returns it
func httpGet(url string) (resp *http.Response, err error) {
   client := http.DefaultClient
   log.Println("GET", url)
   req, err := http.NewRequest(http.MethodGet, url, nil)
   if err != nil { return nil, err }
   req.Header.Set("Range", "bytes=0-")
   resp, err = client.Do(req)
   if err != nil { return nil, err }
   switch resp.StatusCode {
   case http.StatusOK, http.StatusPartialContent:
   default:
      resp.Body.Close()
      return nil, ErrUnexpectedStatusCode(resp.StatusCode)
   }
   return
}

// httpGetBodyBytes reads the whole HTTP body and returns it
func httpGetBodyBytes(url string) ([]byte, error) {
	resp, err := httpGet(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}

type FormatList []Format

var videoRegexpList = []*regexp.Regexp{
	regexp.MustCompile(`(?:v|embed|watch\?v)(?:=|/)([^"&?/=%]{11})`),
	regexp.MustCompile(`(?:=|/)([^"&?/=%]{11})`),
	regexp.MustCompile(`([^"&?/=%]{11})`),
}

// extractVideoID extracts the videoID from the given string
func extractVideoID(videoID string) (string, error) {
	if strings.Contains(videoID, "youtu") || strings.ContainsAny(videoID, "\"?&/<%=") {
		for _, re := range videoRegexpList {
			if isMatch := re.MatchString(videoID); isMatch {
				subs := re.FindStringSubmatch(videoID)
				videoID = subs[1]
			}
		}
	}

	if strings.ContainsAny(videoID, "?&/<%=") {
		return "", ErrInvalidCharactersInVideoID
	}
	if len(videoID) < 10 {
		return "", ErrVideoIDMinLength
	}

	return videoID, nil
}

type DecipherOperation func([]byte) []byte

const (
	ErrCipherNotFound             = constError("cipher not found")
	ErrInvalidCharactersInVideoID = constError("invalid characters in video id")
	ErrVideoIDMinLength           = constError("the video id must be at least 10 characters long")
	ErrReadOnClosedResBody        = constError("http: read on closed response body")
	ErrNotPlayableInEmbed         = constError("embedding of this video has been disabled")
	ErrInvalidPlaylist            = constError("no playlist detected or invalid playlist ID")
)

type constError string

func (e constError) Error() string {
	return string(e)
}


type ErrPlayabiltyStatus struct {
	Status string
	Reason string
}

func (err ErrPlayabiltyStatus) Error() string {
	return fmt.Sprintf("cannot playback and download, status: %s, reason: %s", err.Status, err.Reason)
}

// ErrUnexpectedStatusCode is returned on unexpected HTTP status codes
type ErrUnexpectedStatusCode int

func (err ErrUnexpectedStatusCode) Error() string {
	return fmt.Sprintf("unexpected status code: %d", err)
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

var playerResponsePattern = regexp.MustCompile(`var ytInitialPlayerResponse\s*=\s*(\{.+?\});`)

func (v *Video) parseVideoPage(body []byte) error {
   initialPlayerResponse := playerResponsePattern.FindSubmatch(body)
   if initialPlayerResponse == nil || len(initialPlayerResponse) < 2 {
      return errors.New("no ytInitialPlayerResponse found in the server's answer")
   }
   var prData playerResponseData
   if err := json.Unmarshal(initialPlayerResponse[1], &prData); err != nil {
      return fmt.Errorf("unable to parse player response JSON: %w", err)
   }
   return v.extractDataFromPlayerResponse(prData)
}

func (v *Video) extractDataFromPlayerResponse(prData playerResponseData) error {
	v.Title = prData.VideoDetails.Title
	v.Description = prData.VideoDetails.ShortDescription
	v.Author = prData.VideoDetails.Author
	v.Thumbnails = prData.VideoDetails.Thumbnail.Thumbnails

	if seconds, _ := strconv.Atoi(prData.Microformat.PlayerMicroformatRenderer.LengthSeconds); seconds > 0 {
		v.Duration = time.Duration(seconds) * time.Second
	}

	// Assign Streams
	v.Formats = append(prData.StreamingData.Formats, prData.StreamingData.AdaptiveFormats...)
	if len(v.Formats) == 0 {
		return errors.New("no formats found in the server's answer")
	}

	v.HLSManifestURL = prData.StreamingData.HlsManifestURL
	v.DASHManifestURL = prData.StreamingData.DashManifestURL

	return nil
}

type playerResponseData struct {
   Microformat struct {
      PlayerMicroformatRenderer struct {
         Thumbnail struct {
            Thumbnails []struct {
               URL    string `json:"url"`
               Width  int    `json:"width"`
               Height int    `json:"height"`
            }
         } `json:"thumbnail"`
         Embed struct {
            IframeURL      string `json:"iframeUrl"`
            FlashURL       string `json:"flashUrl"`
            Width          int    `json:"width"`
            Height         int    `json:"height"`
            FlashSecureURL string `json:"flashSecureUrl"`
         } `json:"embed"`
         Title struct {
            SimpleText string `json:"simpleText"`
         } `json:"title"`
         Description struct {
            SimpleText string `json:"simpleText"`
         } `json:"description"`
         LengthSeconds      string   `json:"lengthSeconds"`
         OwnerProfileURL    string   `json:"ownerProfileUrl"`
         ExternalChannelID  string   `json:"externalChannelId"`
         AvailableCountries []string `json:"availableCountries"`
         IsUnlisted         bool     `json:"isUnlisted"`
         HasYpcMetadata     bool     `json:"hasYpcMetadata"`
         ViewCount          string   `json:"viewCount"`
         Category           string   `json:"category"`
         PublishDate        string   `json:"publishDate"`
         OwnerChannelName   string   `json:"ownerChannelName"`
         UploadDate         string   `json:"uploadDate"`
      } `json:"playerMicroformatRenderer"`
   } `json:"microformat"`
   StreamingData struct {
      ExpiresInSeconds string   `json:"expiresInSeconds"`
      Formats          []Format `json:"formats"`
      AdaptiveFormats  []Format `json:"adaptiveFormats"`
      DashManifestURL  string   `json:"dashManifestUrl"`
      HlsManifestURL   string   `json:"hlsManifestUrl"`
   } `json:"streamingData"`
   VideoDetails struct {
      VideoID          string `json:"videoId"`
      Title            string `json:"title"`
      LengthSeconds    string `json:"lengthSeconds"`
      ChannelID        string `json:"channelId"`
      IsOwnerViewing   bool   `json:"isOwnerViewing"`
      ShortDescription string `json:"shortDescription"`
      IsCrawlable      bool   `json:"isCrawlable"`
      Thumbnail        struct {
         Thumbnails []Thumbnail
      }
      AverageRating     float64 `json:"averageRating"`
      AllowRatings      bool    `json:"allowRatings"`
      ViewCount         string  `json:"viewCount"`
      Author            string  `json:"author"`
      IsPrivate         bool    `json:"isPrivate"`
      IsUnpluggedCorpus bool    `json:"isUnpluggedCorpus"`
      IsLiveContent     bool    `json:"isLiveContent"`
   } `json:"videoDetails"`
}

type Thumbnail struct {
	URL    string
	Width  uint
	Height uint
}

type Format struct {
   ItagNo           int    `json:"itag"`
   URL              string `json:"url"`
   MimeType         string `json:"mimeType"`
   Quality          string `json:"quality"`
   Cipher           string `json:"signatureCipher"`
   Bitrate          int    `json:"bitrate"`
   FPS              int    `json:"fps"`
   Width            int    `json:"width"`
   Height           int    `json:"height"`
   LastModified     string `json:"lastModified"`
   ContentLength    string `json:"contentLength"`
   QualityLabel     string `json:"qualityLabel"`
   ProjectionType   string `json:"projectionType"`
   AverageBitrate   int    `json:"averageBitrate"`
   AudioQuality     string `json:"audioQuality"`
   ApproxDurationMs string `json:"approxDurationMs"`
   AudioSampleRate  string `json:"audioSampleRate"`
   AudioChannels    int    `json:"audioChannels"`
   IndexRange *struct {
      Start string `json:"start"`
      End   string `json:"end"`
   } `json:"indexRange"`
}


const API = "https://www.youtube.com/get_video_info"

type player struct {
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

func oldPlayer(id string) (player, error) {
   req, err := http.NewRequest("GET", API, nil)
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
