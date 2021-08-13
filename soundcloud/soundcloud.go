package soundcloud

import (
   "fmt"
   "net/http"
   "strings"
)

// APIOptions are the options for creating an API struct
type APIOptions struct {
   // optional and a new one will be fetched if not provided
   ClientID string
   // the HTTP client to make requests with
   HTTPClient *http.Client
}

// New returns a pointer to a new SoundCloud API struct.
func New() (*client, error) {
   clientID, err := FetchClientID()
   if err != nil {
      return nil, err
   }
   return newClient(
      clientID, http.DefaultClient,
   ), nil
}

// GetDownloadURL retuns the URL to download a track. This is useful if you
// want to implement your own downloading algorithm. If the track has a
// publicly available download link, that link will be preferred and the
// streamType parameter will be ignored. streamType can be either "hls" or
// "progressive", defaults to "progressive"
func (sc client) GetDownloadURL(url string, streamType string) (string, error) {
   streamType = strings.ToLower(streamType)
   if streamType == "" {
      streamType = "progressive"
   }
   if IsURL(url, false, false) {
      info, err := sc.getTrackInfo(GetTrackInfoOptions{
      URL: url,
      })
      if err != nil {
         return "", err
      }
      if len(info) == 0 {
         return "", fmt.Errorf("%v fail", url)
      }
      if info[0].Downloadable && info[0].HasDownloadsLeft {
      downloadURL, err := sc.getDownloadURL(info[0].ID)
      if err != nil {
         return "", err
      }
      return downloadURL, nil
      }
      for _, transcoding := range info[0].Media.Transcodings {
      if strings.ToLower(transcoding.Format.Protocol) == streamType {
      mediaURL, err := sc.getMediaURL(transcoding.URL)
      if err != nil {
      return "", err
      }
      return mediaURL, nil
      }
      }
      mediaURL, err := sc.getMediaURL(info[0].Media.Transcodings[0].URL)
      if err != nil {
      return "", err
      }
      return mediaURL, nil
   }
   return "", fmt.Errorf("%v is not a track URL", url)
}

// Track represents the JSON response of a track's info
type Track struct {
   CommentCount      int64  `json:"comment_count"`
   CreatedAt         string `json:"created_at"`
   Description       string
   DisplayDate       string `json:"display_date"`
   DownloadCount     int64  `json:"download_count"`
   Downloadable      bool
   DurationMS        int64  `json:"duration"`
   FullDurationMS    int64  `json:"full_duration"`
   Genre             string
   HasDownloadsLeft  bool   `json:"has_downloads_left"`
   ID                int64
   Kind              string
   LabelName         string `json:"label_name"`
   LastModified      string `json:"last_modified"`
   LikesCount        int64  `json:"likes_count"`
   Media             Media
   MonetizationModel string `json:"monetization_model"`
   Permalink         string `json:"permalink"`
   PermalinkURL      string `json:"permalink_url"`
   PlaybackCount     int64  `json:"playback_count"`
   Policy            string `json:"polic"`
   Public            bool   `json:"public"`
   RepostsCount      int64  `json:"reposts_count"`
   SecretToken       string `json:"secret_token"`
   Streamable        bool   `json:"streamable"`
   TagList           string `json:"tag_list"`
   Title             string `json:"title"`
   URI               string `json:"uri"`
   UserID            int64  `json:"user_id"`
   WaveformURL       string `json:"waveform_url"`
}

// Media contains an array of transcoding for a track
type Media struct {
	Transcodings []Transcoding `json:"transcodings"`
}

// Transcoding contains information about the transcoding of a track
type Transcoding struct {
   Format  TranscodingFormat `json:"format"`
   Preset  string            `json:"preset"`
   Snipped bool              `json:"snipped"`
   URL     string            `json:"url"`
}

// TranscodingFormat contains the protocol by which the track is delivered ("progressive" or "HLS"), and
// the mime type of the track
type TranscodingFormat struct {
	Protocol string `json:"protocol"`
	MimeType string `json:"mime_type"`
}

// MediaURLResponse is the JSON response of retrieving media information of a track
type MediaURLResponse struct {
	URL string `json:"url"`
}

// DownloadURLResponse is the JSON respose of retrieving media information of a publicly downloadable track
type DownloadURLResponse struct {
	URL string `json:"redirectUri"`
}
