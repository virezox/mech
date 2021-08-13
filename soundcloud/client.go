package soundcloud

import (
   "bytes"
   "encoding/json"
   "fmt"
   "io"
   "net/http"
   "net/url"
   "regexp"
   "strconv"
   "strings"
)

const (
   resolveURL = "https://api-v2.soundcloud.com/resolve"
   trackURL = "https://api-v2.soundcloud.com/tracks"
)

type client struct {
   httpClient *http.Client
   clientID   string
}

func newClient(clientID string, httpClient *http.Client) *client {
   if httpClient == nil {
      httpClient = http.DefaultClient
   }
   return &client{clientID:   clientID, httpClient: httpClient}
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
      return nil, fmt.Errorf("status %v", res.Status)
   }
   return io.ReadAll(res.Body)
}

func (c *client) buildURL(base string, clientID bool, query ...string) (string, error) {
   if len(query)%2 != 0 {
      return "", fmt.Errorf("invalid query %v", query)
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
// PlaylistID and PlaylistSecretToken are necessary to retrieve private tracks
// in private playlists.
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

// resolve is a handy API endpoint that returns info from the given resource
// URL
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

// Kind is a string
type Kind string

// KindTrack is the kind for a Track
const KindTrack Kind = "tracks"


var (
   urlRegex = regexp.MustCompile(`(?m)^https?:\/\/(soundcloud\.com)\/(.*)$`)
)

// IsURL returns true if the provided url is a valid SoundCloud URL
func IsURL(url string, testMobile, testFirebase bool) bool {
   success := false
   if !success {
      success = len(urlRegex.FindAllString(url, -1)) > 0
   }
   return success
}

// StripMobilePrefix removes the prefix for mobile urls. Returns the same string if an error parsing the URL occurs
func StripMobilePrefix(u string) string {
	if !strings.Contains(u, "m.soundcloud.com") {
		return u
	}
	_url, err := url.Parse(u)
	if err != nil {
		return u
	}
	_url.Host = "soundcloud.com"
	return _url.String()
}

func sliceContains(slice []int64, x int64) bool {
   for _, i := range slice {
      if i == x {
         return true
      }
   }
   return false
}

// FetchClientID fetches a SoundCloud client ID. This algorithm is adapted from
// https://www.npmjs.com/package/soundcloud-key-fetch. The basic notion of how
// this function works is that SoundCloud provides a client ID so its web app
// can make API requests. This client ID (along with other intialization data
// for the web app) is provided in a JavaScript file imported through a
// <script> tag in the HTML. This function scrapes the HTML and tries to find
// the URL to that JS file, and then scrapes the JS file to find the client ID.								//
func FetchClientID() (string, error) {
   resp, err := http.Get("https://soundcloud.com")
   if err != nil {
   return "", err
   }
   body, err := io.ReadAll(resp.Body)
   if err != nil {
   return "", err
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
   return "", err
   }
   body, err = io.ReadAll(resp.Body)
   if err != nil {
   return "", err
   }
   bodyString = string(body)
   // Extract the client ID
   if strings.Contains(bodyString, `,client_id:"`) {
   clientID := strings.Split(bodyString, `,client_id:"`)[1]
   clientID = strings.Split(clientID, `"`)[0]
   return clientID, nil
   }
   return "", fmt.Errorf("%v fail", bodyString)
}
