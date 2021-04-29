package youtube

import (
   "fmt"
   "io"
   "log"
   "net/http"
   "regexp"
   "strings"
   "time"
)

// Client offers methods to download video metadata and video streams.
type Client struct {
	// Debug enables debugging output through log package
	Debug bool

	// HTTPClient can be used to set a custom HTTP client.
	// If not set, http.DefaultClient will be used
	HTTPClient *http.Client
}

// GetVideo fetches video metadata
func (c *Client) GetVideo(url string) (*Video, error) {
   id, err := ExtractVideoID(url)
   if err != nil {
      return nil, fmt.Errorf("extractVideoID failed: %w", err)
   }
   return c.videoFromID(id)
}

func (c *Client) videoFromID(id string) (*Video, error) {
	// Circumvent age restriction to pretend access through googleapis.com
	eurl := "https://youtube.googleapis.com/v/" + id
	body, err := c.httpGetBodyBytes("https://youtube.com/get_video_info?video_id="+id+"&eurl="+eurl)
	if err != nil {
		return nil, err
	}

	v := &Video{
		ID: id,
	}

	err = v.parseVideoInfo(body)

	// If the uploader has disabled embedding the video on other sites, parse video page
	if err == ErrNotPlayableInEmbed {
		html, err := c.httpGetBodyBytes("https://www.youtube.com/watch?v="+id)
		if err != nil {
			return nil, err
		}

		return v, v.parseVideoPage(html)
	}

	return v, err
}

// httpGet does a HTTP GET request, checks the response to be a 200 OK and returns it
func (c *Client) httpGet(url string) (resp *http.Response, err error) {
   client := c.HTTPClient
   if client == nil { client = http.DefaultClient }
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
func (c *Client) httpGetBodyBytes(url string) ([]byte, error) {
	resp, err := c.httpGet(url)
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

// ExtractVideoID extracts the videoID from the given string
func ExtractVideoID(videoID string) (string, error) {
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

type ErrResponseStatus struct {
	Status string
	Reason string
}

func (err ErrResponseStatus) Error() string {
	if err.Status == "" {
		return "no response status found in the server's answer"
	}

	if err.Reason == "" {
		return fmt.Sprintf("response status: '%s', no reason given", err.Status)
	}

	return fmt.Sprintf("response status: '%s', reason: '%s'", err.Status, err.Reason)
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


var (
	_ DecipherOperationsCache = NewSimpleCache()
)

const defaultCacheExpiration = time.Minute * time.Duration(5)

type DecipherOperationsCache interface {
	Get(videoID string) []DecipherOperation
	Set(video string, operations []DecipherOperation)
}

type SimpleCache struct {
	videoID    string
	expiredAt  time.Time
	operations []DecipherOperation
}

func NewSimpleCache() *SimpleCache {
	return &SimpleCache{}
}

// Get : get cache  when it has same video id and not expired
func (s SimpleCache) Get(videoID string) []DecipherOperation {
	return s.GetCacheBefore(videoID, time.Now())
}

// GetCacheBefore : can pass time for testing
func (s SimpleCache) GetCacheBefore(videoID string, time time.Time) []DecipherOperation {
	if videoID == s.videoID && s.expiredAt.After(time) {
		operations := make([]DecipherOperation, len(s.operations))
		copy(operations, s.operations)
		return operations
	}
	return nil
}

// Set : set cache with default expiration
func (s *SimpleCache) Set(videoID string, operations []DecipherOperation) {
	s.setWithExpiredTime(videoID, operations, time.Now().Add(defaultCacheExpiration))
}

func (s *SimpleCache) setWithExpiredTime(videoID string, operations []DecipherOperation, time time.Time) {
	s.videoID = videoID
	s.operations = make([]DecipherOperation, len(operations))
	copy(s.operations, operations)
	s.expiredAt = time
}
