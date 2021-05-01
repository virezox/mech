package youtube

import (
   "encoding/json"
   "errors"
   "fmt"
   "io/ioutil"
   "log"
   "net/http"
   "net/url"
   "regexp"
   "strconv"
   "time"
)

type Video struct {
   Author          string
   Description     string
   Duration        time.Duration
   Formats         []Format
   ID              string
   Title           string
}

func (v *Video) parseVideoInfo(body []byte) error {
   answer, err := url.ParseQuery(string(body))
   if err != nil { return err }
   status := answer.Get("status")
   if status != "ok" {
      return fmt.Errorf("status: %q reason: %q", status, answer.Get("reason"))
   }
   // read the streams map
   playerResponse := answer.Get("player_response")
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
   if seconds, _ := strconv.Atoi(prData.Microformat.PlayerMicroformatRenderer.LengthSeconds); seconds > 0 {
      v.Duration = time.Duration(seconds) * time.Second
   }
   // Assign Streams
   v.Formats = append(prData.StreamingData.Formats, prData.StreamingData.AdaptiveFormats...)
   if len(v.Formats) == 0 {
      return errors.New("no formats found in the server's answer")
   }
   return nil
}

// Client offers methods to download video metadata and video streams.
type Client struct {
	// Debug enables debugging output through log package
	Debug bool

	// HTTPClient can be used to set a custom HTTP client.
	// If not set, http.DefaultClient will be used
	HTTPClient *http.Client

	// decipherOpsCache cache decipher operations
	decipherOpsCache DecipherOperationsCache
}

// GetVideo fetches video metadata
func (c *Client) GetVideo(id string) (*Video, error) {
   // Circumvent age restriction to pretend access through googleapis.com
   eurl := "https://youtube.googleapis.com/v/" + id
   body, err := c.httpGetBodyBytes("https://youtube.com/get_video_info?video_id="+id+"&eurl="+eurl)
   if err != nil { return nil, err }
   v := &Video{ID: id}
   err = v.parseVideoInfo(body)
   // If the uploader has disabled embedding the video on other sites, parse video page
   if err == ErrNotPlayableInEmbed {
      html, err := c.httpGetBodyBytes("https://www.youtube.com/watch?v="+id)
      if err != nil { return nil, err }
      return v, v.parseVideoPage(html)
   }
   return v, err
}

// GetStream returns the HTTP response for a specific format
func (c *Client) GetStream(video *Video, format *Format) (*http.Response, error) {
   url, err := c.GetStreamURL(video, format)
   if err != nil { return nil, err }
   return c.httpGet(url)
}

// GetStreamURL returns the url for a specific format
func (c *Client) GetStreamURL(video *Video, format *Format) (string, error) {
   if format.URL != "" { return format.URL, nil }
   cipher := format.Cipher
   if cipher == "" { return "", ErrCipherNotFound }
   return c.decipherURL(video.ID, cipher)
}

// httpGet does a HTTP GET request, checks the response to be a 200 OK and returns it
func (c *Client) httpGet(url string) (resp *http.Response, err error) {
   client := c.HTTPClient
   if client == nil {
      client = http.DefaultClient
   }
   log.Println("GET", url)
   req, err := http.NewRequest(http.MethodGet, url, nil)
   if err != nil { return nil, err }
   // Add range header to disable throttling
   // see https://github.com/kkdai/youtube/pull/170
   req.Header.Set("Range", "bytes=0-")
   resp, err = client.Do(req)
   if err != nil { return nil, err }
   switch resp.StatusCode {
   case http.StatusOK, http.StatusPartialContent:
   default:
      resp.Body.Close()
      return nil, fmt.Errorf("StatusCode %v", resp.StatusCode)
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

	return ioutil.ReadAll(resp.Body)
}
