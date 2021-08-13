package soundcloud

import (
   "fmt"
   "io"
   "net/http"
   "net/url"
   "regexp"
   "strconv"
   "strings"
)

var (
   firebaseURLRegex = regexp.MustCompile(`(?m)^https?:\/\/(soundcloud\.app\.goo\.gl)\/(.*)$`)
   mobileURLRegex = regexp.MustCompile(`(?m)^https?:\/\/(m\.soundcloud\.com)\/(.*)$`)
   urlRegex = regexp.MustCompile(`(?m)^https?:\/\/(soundcloud\.com)\/(.*)$`)
)

// IsURL returns true if the provided url is a valid SoundCloud URL
func IsURL(url string, testMobile, testFirebase bool) bool {
	success := false
	if testMobile {
		success = IsMobileURL(url)
	}

	if testFirebase && !success {
		success = IsFirebaseURL(url)
	}

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

// IsFirebaseURL returns true if the url is a SoundCloud Firebase url (has the following form: https://soundcloud.app.goo.gl/xxxxxxxx)
func IsFirebaseURL(u string) bool {
	return len(firebaseURLRegex.FindAllString(u, -1)) > 0
}

// IsMobileURL returns true if the url is a SoundCloud Firebase url (has the following form: https://m.soundcloud.com/xxxxxx)
func IsMobileURL(u string) bool {
	return len(mobileURLRegex.FindAllString(u, -1)) > 0
}

// IsSearchURL returns true  if the provided url is a valid search url
func IsSearchURL(url string) bool {
	return strings.Index(url, "https://soundcloud.com/search?") == 0
}

// IsPersonalizedTrackURL returns true if the provided url is a valid personalized track url. Ex/
// https://soundcloud.com/discover/sets/personalized-tracks::sam:335899198
func IsPersonalizedTrackURL(url string) bool {
	return strings.Contains(url, "https://soundcloud.com/discover/sets/personalized-tracks::")
}

// ExtractIDFromPersonalizedTrackURL extracts the track ID from a personalized track URL, returns -1
// if no track ID can be extracted
func ExtractIDFromPersonalizedTrackURL(url string) int64 {
	if !IsPersonalizedTrackURL(url) {
		return -1
	}

	split := strings.Split(url, ":")
	if len(split) < 5 {
		return -1
	}

	id, err := strconv.ParseInt(split[4], 10, 64)
	if err != nil {
		return -1
	}

	return id
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
