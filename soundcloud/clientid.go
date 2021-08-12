package soundcloud

import (
   "encoding/json"
   "github.com/pkg/errors"
   "io/ioutil"
   "net/http"
   "strings"
)

// FetchClientID fetches a SoundCloud client ID.
// This algorithm is adapted from:
//     https://www.npmjs.com/package/soundcloud-key-fetch
func FetchClientID() (string, error) {
	// // // // // // // // // // // // // // // // // // // // // // // // // // // // //
	// 																					//
	// The basic notion of how this function works is that SoundCloud provides          //
	// a client ID so its web app can make API requests.								//
	//																					//
	// This client ID (along with other intialization data for the web app) is provided //
	// in a JavaScript file imported through a <script> tag in the HTML.				//
	//																					//
	// This function scrapes the HTML and tries to find the URL to that JS file,		//
	// and then scrapes the JS file to find the client ID.								//
	//																					//
	// // // // // // // // // // // // // // // // // // // // // // // // // // // // //

	resp, err := http.Get("https://soundcloud.com")
	if err != nil {
		return "", errors.Wrap(err, "Failed to fetch SoundCloud Client ID")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", errors.Wrap(err, "Failed to read body while fetching SoundCloud Client ID")
	}

	bodyString := string(body)
	// The link to the JS file with the client ID looks like this:
	// <script crossorigin src="https://a-v2.sndcdn.com/assets/sdfhkjhsdkf.js"></script
	split := strings.Split(bodyString, `<script crossorigin src="`)
	urls := []string{}

	// Extract all the URLS that match our pattern
	for _, raw := range split {
		u := strings.Replace(raw, `"></script>`, "", 1)
		u = strings.Split(u, "\n")[0]
		if string([]rune(u)[0:31]) == "https://a-v2.sndcdn.com/assets/" {
			urls = append(urls, u)
		}
	}

	// It seems like our desired URL is always imported last,
	// so we use urls[len(urls) - 1]
	resp, err = http.Get(urls[len(urls)-1])
	if err != nil {
		return "", errors.Wrap(err, "Failed to fetch SoundCloud Client ID")
	}

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", errors.Wrap(err, "Failed to read body while fetching SoundCloud Client ID")
	}

	bodyString = string(body)

	// Extract the client ID
	if strings.Contains(bodyString, `,client_id:"`) {
		clientID := strings.Split(bodyString, `,client_id:"`)[1]
		clientID = strings.Split(clientID, `"`)[0]
		return clientID, nil
	}

	return "", errors.New("Could not find a SoundCloud client ID")
}


// GetTracks returns any of the items in the PaginatedQuery's collection that match the Track struct type
func (pq *PaginatedQuery) GetTracks() ([]Track, error) {

	tracks := make([]Track, 0)

	for _, item := range pq.Collection {
		track := Track{}
		b, err := json.Marshal(item)
		if err != nil {
			return nil, errors.Wrap(err, "Failed to marshal PaginatedQuery collection item")
		}

		err = json.Unmarshal(b, &track)
		if err != nil {
			continue
		}

		if track.Kind != "track" {
			continue
		}

		tracks = append(tracks, track)
	}

	return tracks, nil
}

// GetPlaylists returns any of the items in the PaginatedQuery's collection that match the Playlist struct type
func (pq *PaginatedQuery) GetPlaylists() ([]Playlist, error) {
	playlists := make([]Playlist, 0)

	for _, item := range pq.Collection {
		playlist := Playlist{}
		b, err := json.Marshal(item)
		if err != nil {
			return nil, errors.Wrap(err, "Failed to marshal PaginatedQuery collection item")
		}

		err = json.Unmarshal(b, &playlist)
		if err != nil {
			continue
		}

		if playlist.Kind != "playlist" {
			continue
		}

		playlists = append(playlists, playlist)
	}

	return playlists, nil
}

// GetLikes returns any of the items in the PaginatedQuery's collection that match the Like struct type
func (pq *PaginatedQuery) GetLikes() ([]Like, error) {
	likes := make([]Like, 0)

	for _, item := range pq.Collection {
		like := Like{}
		b, err := json.Marshal(item)
		if err != nil {
			return nil, errors.Wrap(err, "Failed to marshal PaginatedQuery collection item")
		}

		err = json.Unmarshal(b, &like)
		if err != nil {
			continue
		}

		if like.Kind != "like" {
			continue
		}

		likes = append(likes, like)
	}

	return likes, nil
}
