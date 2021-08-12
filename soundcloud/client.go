package soundcloud

import (
   "bytes"
   "encoding/json"
   "fmt"
   "github.com/pkg/errors"
   "io"
   "io/ioutil"
   "net/http"
   "net/url"
   "strconv"
   "strings"
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

const trackURL = "https://api-v2.soundcloud.com/tracks"
const resolveURL = "https://api-v2.soundcloud.com/resolve"
const usersURL = "https://api-v2.soundcloud.com/users/"
const searchURL = "https://api-v2.soundcloud.com/search"

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
			return nil, errors.Wrap(err, "Failed to marshal json body")
		}
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonBytes))
	if err != nil {
		return nil, errors.Wrap(err, "Failed to make http request")
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
		return "", fmt.Errorf("Invalid query: URL (%s) Query: (%s)", base, strings.Join(query, ","))
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
			u, err = c.buildURL(trackURL, true, "ids", strings.Join(ids, ","), "playlistId", fmt.Sprintf("%d", options.PlaylistID), "playlistSecretToken", options.PlaylistSecretToken)
		}
		if err != nil {
			return nil, errors.Wrap(err, "Failed to build URL for getTrackInfo()")
		}

		data, err = c.makeRequest("GET", u, nil)
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(data, &trackInfo)

		if err != nil {
			return nil, errors.Wrap(err, "JSON is not valid track info")
		}
	} else if options.URL != "" {
		// TO-DO: Validate the URL
		data, err = c.resolve(options.URL)
		if err != nil {
			return nil, err
		}

		trackSingle := Track{}
		err = json.Unmarshal(data, &trackSingle)
		if err != nil {
			return nil, errors.Wrap(err, "Failed to unmarshal track JSON data")
		}
		trackInfo = []Track{trackSingle}
	} else {
		return nil, errors.New("Invalid options. URL or ID must be provided")
	}

	if options.ID != nil && len(options.ID) > 0 {
		// For some reason the track URL returns the tracks out of order,
		// so we need to sort the response to maintain consistency

		// Private tracks will not be fetched if options.PlaylistID and options.PlaylistSecretToken
		// are not provided, so we need to update the IDs slice before we sort

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
		return "", errors.Wrap(err, "Failed to build URL for getMediaURL")
	}

	media := &MediaURLResponse{}
	data, err := c.makeRequest("GET", u, nil)
	if err != nil {
		return "", err
	}

	err = json.Unmarshal(data, media)
	if err != nil {
		return "", errors.Wrap(err, "Failed to unmarshal JSON response in getMediaURL")
	}

	return media.URL, nil
}

// getDownloadURL gets the download URL of a publicly downloadable track
func (c *client) getDownloadURL(id int64) (string, error) {
	u, err := c.buildURL(fmt.Sprintf("https://api-v2.soundcloud.com/tracks/%d/download", id), true)
	if err != nil {
		return "", errors.Wrap(err, "Failed to build URL for getDownloadURL")
	}

	res := &DownloadURLResponse{}
	data, err := c.makeRequest("GET", u, nil)
	if err != nil {
		return "", err
	}

	err = json.Unmarshal(data, res)
	if err != nil {
		return "", errors.Wrap(err, "Failed to unmarshal JSON response in getDownloadURL")
	}

	return res.URL, nil
}

func (c *client) downloadProgressive(url string, dst io.Writer) error {
	// The track audio file is just a regular audio file that can be downloaded
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return errors.Wrap(err, "Failed to make request")
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
		return errors.Wrap(err, "downloadProgressive() failed")
	}

	return nil
}

func (c *client) getPlaylistInfo(url string) (Playlist, error) {
	playlist := Playlist{}
	u, err := c.buildURL(resolveURL, true, "url", url)
	if err != nil {
		return playlist, errors.Wrap(err, "Failed to build URL for getPlaylistInfo")
	}

	data, err := c.makeRequest("GET", u, nil)
	if err != nil {
		return playlist, err
	}

	err = json.Unmarshal(data, &playlist)

	if err != nil {
		return playlist, errors.Wrap(err, "Returned JSON is not valid track info")
	}

	playlistID := playlist.ID
	playlistSecretToken := playlist.SecretToken
	if playlist.TrackCount > 5 {
		// SoundCloud provides info for the first 5 tracks,
		// the rest must be retrieved.
		ids := make([]int64, playlist.TrackCount-5)

		count := 0
		for _, track := range playlist.Tracks[5:] {
			ids[count] = track.ID
			count++
		}

		playlist.Tracks = playlist.Tracks[:5]

		if len(ids) > 50 {
			// The SoundCloud API limits querying tracks to 50 at a time,
			// so we have to split the requests.

			temp := make([]Track, len(ids))
			playlist.Tracks = append(playlist.Tracks, temp...)

			workers := len(ids) / 50

			type result struct {
				startIndex int
				trackInfo  []Track
			}

			errChan := make(chan error)
			resultsChan := make(chan result)
			for i := 0; i <= workers; i++ {
				start := i * 50
				end := start + 50
				if i == workers {
					end = start + (len(ids) % 50)
				}
				go func() {
					trackInfo, err := c.getTrackInfo(GetTrackInfoOptions{
						ID:                  ids[start:end],
						PlaylistID:          playlistID,
						PlaylistSecretToken: playlistSecretToken,
					})

					if err != nil {
						errChan <- err
						return
					}

					resultsChan <- result{
						startIndex: start,
						trackInfo:  trackInfo,
					}
				}()
			}

			completeCount := -1

			for {
				select {
				case err = <-errChan:
					if err != nil {
						return playlist, err
					}
				case r := <-resultsChan:
					completeCount++

					for i, track := range r.trackInfo {
						playlist.Tracks[r.startIndex+i+5] = track
					}

					if completeCount == workers {
						break
					}
				}

				if completeCount == workers {
					break
				}
			}

		} else {
			trackInfo, err := c.getTrackInfo(GetTrackInfoOptions{
				ID:                  ids,
				PlaylistID:          playlistID,
				PlaylistSecretToken: playlistSecretToken,
			})

			if err != nil {
				return playlist, err
			}

			for _, track := range trackInfo {
				playlist.Tracks = append(playlist.Tracks, track)
			}
		}
	}

	data, err = json.Marshal(&playlist)
	if err != nil {
		return playlist, errors.Wrap(err, "Failed to get playlist data")
	}

	return playlist, nil
}

// resolve is a handy API endpoint that returns info from the given resource URL
func (c *client) resolve(url string) ([]byte, error) {
	u, err := c.buildURL(resolveURL, true, "url", strings.TrimRight(url, "/"))
	if err != nil {
		return nil, errors.Wrap(err, "Failed to build URL for resolve()")
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

func (c *client) getUser(options GetUserOptions) (User, error) {
	var user User
	var u string
	var err error

	if options.ProfileURL != "" {
		u, err = c.buildURL(resolveURL, true, "url", options.ProfileURL)
	} else if options.ID != 0 {
		u, err = c.buildURL(usersURL+strconv.FormatInt(options.ID, 10), true)
	} else {
		return user, errors.New("One of options.ProfileURL or options.ID is required")
	}

	if err != nil {
		return user, errors.Wrap(err, "Failed to build URL for getUser()")
	}

	data, err := c.makeRequest("GET", u, nil)
	if err != nil {
		return user, err
	}

	err = json.Unmarshal(data, &user)
	if err != nil {
		return user, errors.Wrap(err, "Failed to get user")
	}

	return user, nil
}

// GetLikesOptions are the options for getting a user's likes.
type GetLikesOptions struct {
	ProfileURL string // URL to the user's profile (will use this or ID to choose user)
	ID         int64  //  User's ID if you have it
	Limit      int    // How many tracks to return (defaults to 10)
	Offset     int    // How many tracks to offset by (used for pagination; defaults to 0)
	Type       string // What type of resource to return. One of ["track", "playlist", "all"]. Defaults to "all"
}

func (c *client) getLikes(options GetLikesOptions) (*PaginatedQuery, error) {
	var query PaginatedQuery
	var u string // URL takes the form: https://api-v2.soundcloud.com/users/<id>/likes
	var err error

	if options.ProfileURL != "" {
		user, err := c.getUser(GetUserOptions{ProfileURL: options.ProfileURL})
		if err != nil {
			return nil, err
		}

		options.ID = user.ID
	} else if options.ID == 0 {
		return nil, errors.New("One of options.ProfileURL or options.ID is required")
	}

	if options.Limit == 0 {
		options.Limit = 10
	}

	if options.Offset == 0 {
		options.Offset = 0
	}

	if options.Type == "" {
		options.Type = "all"
	}

	if options.Type == "track" {
		options.Type = "track_likes"
	} else if options.Type == "playlist" {
		options.Type = "playlist_likes"
	} else {
		options.Type = "likes"
	}

	u, err = c.buildURL(usersURL+strconv.FormatInt(options.ID, 10)+"/"+options.Type, true, "limit", strconv.Itoa(options.Limit), "offset", strconv.Itoa(options.Offset))
	if err != nil {
		return nil, errors.Wrap(err, "Failed to build URL for getLikes()")
	}
	data, err := c.makeRequest("GET", u, nil)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &query)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to unmarshal received likes data")
	}

	return &query, nil
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

		u, err = c.buildURL(searchURL+string(kind), true, "q", options.Query, "limit", strconv.Itoa(options.Limit), "offset", strconv.Itoa(options.Offset))

		if err != nil {
			return nil, errors.Wrap(err, "Failed to build URL for search()")
		}
	}

	data, err := c.makeRequest("GET", u, nil)

	if err != nil {
		return nil, err
	}

	response := &PaginatedQuery{}
	err = json.Unmarshal(data, response)

	if err != nil {
		return nil, errors.Wrap(err, "Failed to unmarshal response")
	}

	return response, nil
}
