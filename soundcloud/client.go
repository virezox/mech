package soundcloud

import (
   "bytes"
   "encoding/json"
   "fmt"
   "io"
   "io/ioutil"
   "net/http"
   "net/url"
   "strconv"
   "strings"
)

const (
   resolveURL = "https://api-v2.soundcloud.com/resolve"
   searchURL = "https://api-v2.soundcloud.com/search"
   trackURL = "https://api-v2.soundcloud.com/tracks"
   usersURL = "https://api-v2.soundcloud.com/users/"
)

type client struct {
	httpClient *http.Client
	clientID   string
}

// FailedRequestError is an error response from the SoundCloud API
type FailedRequestError struct {
	Status int
	ErrMsg string
}

func (f *FailedRequestError) Error() string {
	if f.ErrMsg == "" {
		return fmt.Sprintf("Request returned non 2xx Status: %d", f.Status)
	}

	return fmt.Sprintf("Request failed with Status %d: %s", f.Status, f.ErrMsg)
}

func newClient(clientID string, httpClient *http.Client) *client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	return &client{
		httpClient: httpClient,
		clientID:   clientID,
	}
}

func (c *client) makeRequest(method, url string, jsonBody interface{}) ([]byte, error) {
   var jsonBytes []byte
   var err error
   if jsonBody != nil {
      jsonBytes, err = json.Marshal(jsonBody)
      if err != nil {
         return nil, err
      }
   }
   req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonBytes))
   if err != nil {
      return nil, err
   }
   res, err := c.httpClient.Do(req)
   if err != nil {
   return nil, err
   }
   if res.StatusCode < 200 || res.StatusCode > 299 {
   if data, err := ioutil.ReadAll(res.Body); err == nil {
   return nil, &FailedRequestError{Status: res.StatusCode, ErrMsg: string(data)}
   }
   return nil, &FailedRequestError{Status: res.StatusCode}
   }
   data, err := ioutil.ReadAll(res.Body)
   if err != nil {
   return data, nil
   }
   return data, nil
}

func (c *client) buildURL(base string, clientID bool, query ...string) (string, error) {
	if len(query)%2 != 0 {
		return "", fmt.Errorf("invalid query: URL %q Query: %q", base, strings.Join(query, ","))
	}

	u, err := url.Parse(string(base))
	if err != nil {
		return "", err
	}
	q := u.Query()

	for i := 0; i < len(query); i += 2 {
		q.Add(query[i], query[i+1])
	}

	if clientID {
		q.Add("client_id", c.clientID)
	}

	u.RawQuery = q.Encode()
	return u.String(), nil
}

// GetTrackInfoOptions can contain the URL of the track or the ID of the track.
// PlaylistID and PlaylistSecretToken are necessary to retrieve private tracks in private playlists.
type GetTrackInfoOptions struct {
	URL                 string
	ID                  []int64
	PlaylistID          int64
	PlaylistSecretToken string
}

func (c *client) getTrackInfo(options GetTrackInfoOptions) ([]Track, error) {
   var u string
   var data []byte
   var err error
   var trackInfo []Track
   if options.ID != nil && len(options.ID) > 0 {
      ids := []string{}
      for _, id := range options.ID {
         ids = append(ids, strconv.FormatInt(id, 10))
      }
      if options.PlaylistID == 0 && options.PlaylistSecretToken == "" {
         u, err = c.buildURL(trackURL, true, "ids", strings.Join(ids, ","))
      } else {
         u, err = c.buildURL(
            trackURL, true, "ids", strings.Join(ids, ","), "playlistId",
            fmt.Sprintf("%d", options.PlaylistID), "playlistSecretToken",
            options.PlaylistSecretToken,
         )
      }
      if err != nil {
         return nil, err
      }
      data, err = c.makeRequest("GET", u, nil)
      if err != nil {
         return nil, err
      }
      err = json.Unmarshal(data, &trackInfo)
      if err != nil {
         return nil, err
      }
   } else if options.URL != "" {
      data, err = c.resolve(options.URL)
      if err != nil {
         return nil, err
      }
      trackSingle := Track{}
      err = json.Unmarshal(data, &trackSingle)
      if err != nil {
         return nil, err
      }
      trackInfo = []Track{trackSingle}
   } else {
      return nil, fmt.Errorf("%v invalid", options)
   }
   if options.ID != nil && len(options.ID) > 0 {
      trimmedIDs := []int64{}
      trackInfoIDs := []int64{}
      for _, track := range trackInfo {
         trackInfoIDs = append(trackInfoIDs, track.ID)
      }
      for _, id := range options.ID {
         if sliceContains(trackInfoIDs, id) {
            trimmedIDs = append(trimmedIDs, id)
         }
      }
      c.sortTrackInfo(trimmedIDs, trackInfo)
   }
   return trackInfo, nil
}

func (c *client) sortTrackInfo(ids []int64, tracks []Track) {
	// Bubble Sort for now. Maybe switch to a more efficient sorting algorithm later??
	//
	// Because the API request in getTrackInfo is limited to 50 tracks at once
	// time complexity will always be <= O(50^2)

	for j, id := range ids {

		if tracks[j].ID != id {
			for k := 0; k < len(tracks); k++ {
				if tracks[k].ID == id {
					temp := tracks[j]
					tracks[j] = tracks[k]
					tracks[k] = temp
				}
			}
		}
	}
}

func (c *client) getMediaURL(url string) (string, error) {
	// The media URL is the actual link to the audio file for the track
	u, err := c.buildURL(url, true)
	if err != nil {
               return "", err
	}

	media := &MediaURLResponse{}
	data, err := c.makeRequest("GET", u, nil)
	if err != nil {
		return "", err
	}

	err = json.Unmarshal(data, media)
	if err != nil {
               return "", err
	}

	return media.URL, nil
}

// getDownloadURL gets the download URL of a publicly downloadable track
func (c *client) getDownloadURL(id int64) (string, error) {
	u, err := c.buildURL(fmt.Sprintf("https://api-v2.soundcloud.com/tracks/%d/download", id), true)
	if err != nil {
               return "", err
	}

	res := &DownloadURLResponse{}
	data, err := c.makeRequest("GET", u, nil)
	if err != nil {
		return "", err
	}

	err = json.Unmarshal(data, res)
	if err != nil {
               return "", err
	}

	return res.URL, nil
}

func (c *client) downloadProgressive(url string, dst io.Writer) error {
	// The track audio file is just a regular audio file that can be downloaded
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
               return err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}

	if res.StatusCode < 200 || res.StatusCode > 299 {
		if data, err := ioutil.ReadAll(res.Body); err == nil {
			return &FailedRequestError{Status: res.StatusCode, ErrMsg: string(data)}
		}
		return &FailedRequestError{Status: res.StatusCode}
	}

	_, err = io.Copy(dst, res.Body)
	if err != nil {
               return err
	}

	return nil
}

// resolve is a handy API endpoint that returns info from the given resource URL
func (c *client) resolve(url string) ([]byte, error) {
	u, err := c.buildURL(resolveURL, true, "url", strings.TrimRight(url, "/"))
	if err != nil {
               return nil, err
	}

	data, err := c.makeRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// GetUserOptions contains either the profile url of the user or the ID of the user
type GetUserOptions struct {
	ProfileURL string
	ID         int64
}

// GetLikesOptions are the options for getting a user's likes.
type GetLikesOptions struct {
	ProfileURL string // URL to the user's profile (will use this or ID to choose user)
	ID         int64  //  User's ID if you have it
	Limit      int    // How many tracks to return (defaults to 10)
	Offset     int    // How many tracks to offset by (used for pagination; defaults to 0)
	Type       string // What type of resource to return. One of ["track", "playlist", "all"]. Defaults to "all"
}

// SearchOptions are the parameters for executing a search
type SearchOptions struct {
	// This is the NextHref property of PaginatedQuery structs
	QueryURL string
	Query    string
	// Number of items to return
	Limit int
	// Number of items to offset by (for pagination)
	Offset int
	// The type of item to return
	Kind Kind
}

// Kind is a string
type Kind string

// KindTrack is the kind for a Track
const KindTrack Kind = "tracks"

// KindAlbum is the kind for an album
const KindAlbum Kind = "albums"

// KindPlaylist is the kind for a playlist
const KindPlaylist Kind = "playlist"

// KindUser is the kind for a user
const KindUser Kind = "users"

func (c *client) search(options SearchOptions) (*PaginatedQuery, error) {
	var u string
	var err error

	if options.Limit == 0 {
		options.Limit = 10
	}

	if options.Kind == KindPlaylist {
		options.Kind = "playlist_without_albums"
	}

	if options.QueryURL != "" {
		u = options.QueryURL
	} else {
		kind := "/" + options.Kind
		if kind == "/" {
			kind = ""
		}

		u, err = c.buildURL(
                  searchURL+string(kind), true, "q", options.Query, "limit",
                  strconv.Itoa(options.Limit), "offset",
                  strconv.Itoa(options.Offset),
               )

		if err != nil {
                        return nil, err
		}
	}

	data, err := c.makeRequest("GET", u, nil)

	if err != nil {
		return nil, err
	}

	response := &PaginatedQuery{}
	err = json.Unmarshal(data, response)

	if err != nil {
               return nil, err
	}

	return response, nil
}