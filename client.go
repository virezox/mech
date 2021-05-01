package youtube

import (
   "encoding/json"
   "fmt"
   "io/ioutil"
   "log"
   "net/http"
)

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

// GetPlaylist fetches playlist metadata
func (c *Client) GetPlaylist(url string) (*Playlist, error) {
	return c.GetPlaylistContext(url)
}

// GetPlaylistContext fetches playlist metadata, with a context, along with a list of Videos, and some basic information
// for these videos. Playlist entries cannot be downloaded, as they lack all the required metadata, but
// can be used to enumerate all IDs, Authors, Titles, etc.
func (c *Client) GetPlaylistContext(url string) (*Playlist, error) {
	id, err := extractPlaylistID(url)
	if err != nil {
		return nil, fmt.Errorf("extractPlaylistID failed: %w", err)
	}
	requestURL := fmt.Sprintf(playlistFetchURL, id)
	resp, err := c.httpGet(requestURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	data, err := extractPlaylistJSON(resp.Body)
	if err != nil {
		return nil, err
	}
	p := &Playlist{ID: id}
	return p, json.Unmarshal(data, p)
}

func (c *Client) VideoFromPlaylistEntry(entry *PlaylistEntry) (*Video, error) {
	return c.videoFromID(entry.ID)
}

func (c *Client) VideoFromPlaylistEntryContext(entry *PlaylistEntry) (*Video, error) {
	return c.videoFromID(entry.ID)
}

// GetStream returns the HTTP response for a specific format
func (c *Client) GetStream(video *Video, format *Format) (*http.Response, error) {
	return c.GetStreamContext(video, format)
}

// GetStreamContext returns the HTTP response for a specific format with a context
func (c *Client) GetStreamContext(video *Video, format *Format) (*http.Response, error) {
	url, err := c.GetStreamURLContext(video, format)
	if err != nil {
		return nil, err
	}

	return c.httpGet(url)
}

// GetStreamURL returns the url for a specific format
func (c *Client) GetStreamURL(video *Video, format *Format) (string, error) {
	return c.GetStreamURLContext(video, format)
}

// GetStreamURLContext returns the url for a specific format with a context
func (c *Client) GetStreamURLContext(video *Video, format *Format) (string, error) {
	if format.URL != "" {
		return format.URL, nil
	}

	cipher := format.Cipher
	if cipher == "" {
		return "", ErrCipherNotFound
	}

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

	return ioutil.ReadAll(resp.Body)
}
