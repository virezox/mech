package soundcloud

import (
   "fmt"
   "net/http"
   "strings"
)

// API is a wrapper for the SoundCloud private API used internally for soundcloud.com
type API struct {
	client              *client
	StripMobilePrefix   bool
	ConvertFirebaseURLs bool
}

// APIOptions are the options for creating an API struct
type APIOptions struct {
	ClientID            string       // optional and a new one will be fetched if not provided
	HTTPClient          *http.Client // the HTTP client to make requests with
	StripMobilePrefix   bool         // whether or not to convert mobile URLs to regular URLs
	ConvertFirebaseURLs bool         // whether or not to convert SoundCloud firebase URLs to regular URLs
}

// New returns a pointer to a new SoundCloud API struct.
func New(options APIOptions) (*API, error) {

	if options.ClientID == "" {
		var err error
		options.ClientID, err = FetchClientID()
		if err != nil {
                        return nil, err
		}
	}

	if options.HTTPClient == nil {
		options.HTTPClient = http.DefaultClient
	}

	return &API{
		client:              newClient(options.ClientID, options.HTTPClient),
		StripMobilePrefix:   options.StripMobilePrefix,
		ConvertFirebaseURLs: options.ConvertFirebaseURLs,
	}, nil
}

// GetDownloadURL retuns the URL to download a track. This is useful if you
// want to implement your own downloading algorithm. If the track has a
// publicly available download link, that link will be preferred and the
// streamType parameter will be ignored. streamType can be either "hls" or
// "progressive", defaults to "progressive"
func (sc *API) GetDownloadURL(url string, streamType string) (string, error) {
   url, err := sc.prepareURL(url)
   if err != nil {
      return "", err
   }
   streamType = strings.ToLower(streamType)
   if streamType == "" {
      streamType = "progressive"
   }
   if IsURL(url, false, false) {
      info, err := sc.client.getTrackInfo(GetTrackInfoOptions{
      URL: url,
      })
      if err != nil {
         return "", err
      }
      if len(info) == 0 {
         return "", fmt.Errorf("%v fail", url)
      }
      if info[0].Downloadable && info[0].HasDownloadsLeft {
      downloadURL, err := sc.client.getDownloadURL(info[0].ID)
      if err != nil {
         return "", err
      }
      return downloadURL, nil
      }
      for _, transcoding := range info[0].Media.Transcodings {
      if strings.ToLower(transcoding.Format.Protocol) == streamType {
      mediaURL, err := sc.client.getMediaURL(transcoding.URL)
      if err != nil {
      return "", err
      }
      return mediaURL, nil
      }
      }
      mediaURL, err := sc.client.getMediaURL(info[0].Media.Transcodings[0].URL)
      if err != nil {
      return "", err
      }
      return mediaURL, nil
   }
   return "", fmt.Errorf("%v is not a track URL", url)
}

func (sc *API) prepareURL(url string) (string, error) {
   if sc.StripMobilePrefix {
      if IsMobileURL(url) {
         url = StripMobilePrefix(url)
      }
   }
   return url, nil
}
