package youtube

import (
   "context"
   "fmt"
   "io"
   "log"
   "net/http"
   "strconv"
   "strings"
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
	return c.GetVideoContext(context.Background(), url)
}

// GetVideoContext fetches video metadata with a context
func (c *Client) GetVideoContext(ctx context.Context, url string) (*Video, error) {
	id, err := ExtractVideoID(url)
	if err != nil {
		return nil, fmt.Errorf("extractVideoID failed: %w", err)
	}
	return c.videoFromID(ctx, id)
}

func (c *Client) videoFromID(ctx context.Context, id string) (*Video, error) {
	// Circumvent age restriction to pretend access through googleapis.com
	eurl := "https://youtube.googleapis.com/v/" + id
	body, err := c.httpGetBodyBytes(ctx, "https://youtube.com/get_video_info?video_id="+id+"&eurl="+eurl)
	if err != nil {
		return nil, err
	}

	v := &Video{
		ID: id,
	}

	err = v.parseVideoInfo(body)

	// If the uploader has disabled embedding the video on other sites, parse video page
	if err == ErrNotPlayableInEmbed {
		html, err := c.httpGetBodyBytes(ctx, "https://www.youtube.com/watch?v="+id)
		if err != nil {
			return nil, err
		}

		return v, v.parseVideoPage(html)
	}

	return v, err
}

// httpGet does a HTTP GET request, checks the response to be a 200 OK and returns it
func (c *Client) httpGet(ctx context.Context, url string) (resp *http.Response, err error) {
   client := c.HTTPClient
   if client == nil { client = http.DefaultClient }
   log.Println("GET", url)
   req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
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
func (c *Client) httpGetBodyBytes(ctx context.Context, url string) ([]byte, error) {
	resp, err := c.httpGet(ctx, url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}

type FormatList []Format

// Quality returns a new FormatList filtered by quality, quality label or itag,
// but not audio quality
func (list FormatList) Quality(quality string) (result FormatList) {
	for _, f := range list {
		itag, _ := strconv.Atoi(quality)
		if itag == f.ItagNo || strings.Contains(f.Quality, quality) || strings.Contains(f.QualityLabel, quality) {
			result = append(result, f)
		}
	}
	return result
}
