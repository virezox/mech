package soundcloud

import (
   "bytes"
   "encoding/json"
   "fmt"
   "io"
   "net/http"
   "net/http/httputil"
   "os"
   "strings"
)

const (
   resolveURL = "https://api-v2.soundcloud.com/resolve"
   trackURL = "https://api-v2.soundcloud.com/tracks"
)

type Client struct {
   ID string
}

// NewClient returns a pointer to a new SoundCloud API struct. First fetch a
// SoundCloud client ID. This algorithm is adapted from
// https://www.npmjs.com/package/soundcloud-key-fetch. The basic notion of how
// this function works is that SoundCloud provides a client ID so its web app
// can make API requests. This client ID (along with other intialization data
// for the web app) is provided in a JavaScript file imported through a
// <script> tag in the HTML. This function scrapes the HTML and tries to find
// the URL to that JS file, and then scrapes the JS file to find the client ID.
func NewClient() (*Client, error) {
   addr := "https://soundcloud.com"
   fmt.Println("GET", addr)
   resp, err := http.Get(addr)
   if err != nil {
      return nil, err
   }
   body, err := io.ReadAll(resp.Body)
   if err != nil {
      return nil, err
   }
   // The link to the JS file with the client ID looks like this:
   // <script crossorigin
   // src="https://a-v2.sndcdn.com/assets/sdfhkjhsdkf.js"></script
   split := bytes.Split(body, []byte(`<script crossorigin src="`))
   var urls []string
   // Extract all the URLS that match our pattern
   for _, raw := range split {
      u := bytes.Replace(raw, []byte(`"></script>`), nil, 1)
      u = bytes.Split(u, []byte{'\n'})[0]
      if bytes.HasPrefix(u, []byte("https://a-v2.sndcdn.com/assets/")) {
         urls = append(urls, string(u))
      }
   }
   // It seems like our desired URL is always imported last
   addr = urls[len(urls)-1]
   fmt.Println("GET", addr)
   resp, err = http.Get(addr)
   if err != nil {
      return nil, err
   }
   body, err = io.ReadAll(resp.Body)
   if err != nil {
      return nil, err
   }
   // Extract the client ID
   if !bytes.Contains(body, []byte(`,client_id:"`)) {
      return nil, fmt.Errorf("%q fail", addr)
   }
   clientID := bytes.Split(body, []byte(`,client_id:"`))[1]
   clientID = bytes.Split(clientID, []byte{'"'})[0]
   return &Client{
      string(clientID),
   }, nil
}

// GetDownloadURL retuns the URL to download a track. This is useful if you
// want to implement your own downloading algorithm. If the track has a
// publicly available download link, that link will be preferred and the
// streamType parameter will be ignored. streamType can be either "hls" or
// "progressive", defaults to "progressive"
func (sc Client) GetDownloadURL(addr string) (string, error) {
   if !strings.HasPrefix(addr, "https://soundcloud.com/") {
      return "", fmt.Errorf("%q is not a track URL", addr)
   }
   tracks, err := sc.getTrackInfo(addr, nil)
   if err != nil {
      return "", err
   }
   if len(tracks) == 0 {
      return "", fmt.Errorf("%v fail", addr)
   }
   for _, transcoding := range tracks[0].Media.Transcodings {
      if strings.ToLower(transcoding.Format.Protocol) == "progressive" {
         mediaURL, err := sc.getMediaURL(transcoding.URL)
         if err != nil {
            return "", err
         }
         return mediaURL, nil
      }
   }
   return "", fmt.Errorf("%q fail", addr)
}

// The media URL is the actual link to the audio file for the track
func (c Client) getMediaURL(addr string) (string, error) {
   req, err := http.NewRequest("GET", addr, nil)
   if err != nil {
      return "", err
   }
   q := req.URL.Query()
   q.Set("client_id", c.ID)
   req.URL.RawQuery = q.Encode()
   d, err := httputil.DumpRequest(req, false)
   if err != nil {
      return "", err
   }
   os.Stdout.Write(d)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return "", err
   }
   defer res.Body.Close()
   // MediaURLResponse is the JSON response of retrieving media information of
   // a track
   var media struct {
      URL string
   }
   if err := json.NewDecoder(res.Body).Decode(&media); err != nil {
      return "", err
   }
   return media.URL, nil
}

func (c Client) getTrackInfo(addr string, ids []int64) ([]Track, error) {
   if addr == "" {
      return nil, fmt.Errorf("%q invalid", addr)
   }
   req, err := http.NewRequest("GET", resolveURL, nil)
   if err != nil {
      return nil, err
   }
   q := req.URL.Query()
   q.Set("client_id", c.ID)
   q.Set("url", strings.TrimRight(addr, "/"))
   req.URL.RawQuery = q.Encode()
   d, err := httputil.DumpRequest(req, false)
   if err != nil {
      return nil, err
   }
   os.Stdout.Write(d)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   data, err := io.ReadAll(res.Body)
   if err != nil {
      return nil, err
   }
   var trackSingle Track
   if err := json.Unmarshal(data, &trackSingle); err != nil {
      return nil, err
   }
   return []Track{trackSingle}, nil
}

// Track represents the JSON response of a track's info
type Track struct {
   Downloadable      bool
   HasDownloadsLeft  bool   `json:"has_downloads_left"`
   CreatedAt         string `json:"created_at"`
   Description       string
   DurationMS        int64  `json:"duration"`
   FullDurationMS    int64  `json:"full_duration"`
   ID                int64
   Kind string
   // Media contains an array of transcoding for a track
   Media struct {
      // Transcoding contains information about the transcoding of a track
      Transcodings []struct {
         // TranscodingFormat contains the protocol by which the track is
         // delivered ("progressive" or "HLS"), and the mime type of the track
         Format struct {
            MimeType string `json:"mime_type"`
            Protocol string
         }
         Preset  string
         Snipped bool
         URL     string
      }
   }
   Permalink string
   PermalinkURL string `json:"permalink_url"`
   PlaybackCount int64  `json:"playback_count"`
   SecretToken string `json:"secret_token"`
   Streamable bool
   Title string
   URI string
   WaveformURL string `json:"waveform_url"`
}
