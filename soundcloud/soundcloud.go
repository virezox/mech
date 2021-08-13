package soundcloud

import (
   "bytes"
   "encoding/json"
   "fmt"
   "github.com/89z/mech"
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
   res, err := http.Get(addr)
   if err != nil {
      return nil, err
   }
   scan := mech.NewScanner(res.Body)
   // The link to the JS file with the client ID looks like this:
   // <script crossorigin src="https://a-v2.sndcdn.com/assets/sdfhkjhsdkf.js">
   // Extract all the URLS that match our pattern. It seems like our desired
   // URL is imported last
   for scan.ScanAttr("crossorigin", "") {
      addr = scan.Attr("src")
   }
   fmt.Println("GET", addr)
   if res, err := http.Get(addr); err != nil {
      return nil, err
   } else {
      defer res.Body.Close()
      body, err := io.ReadAll(res.Body)
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
}

// GetDownloadURL retuns the URL to download a track. This is useful if you
// want to implement your own downloading algorithm. If the track has a
// publicly available download link, that link will be preferred and the
// streamType parameter will be ignored. streamType can be either "hls" or
// "progressive", defaults to "progressive"
func (c Client) GetDownloadURL(addr string) (string, error) {
   if !strings.HasPrefix(addr, "https://soundcloud.com/") {
      return "", fmt.Errorf("%q is not a track URL", addr)
   }
   req, err := http.NewRequest("GET", resolveURL, nil)
   if err != nil {
      return "", err
   }
   q := req.URL.Query()
   q.Set("client_id", c.ID)
   q.Set("url", strings.TrimRight(addr, "/"))
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
   var trackSingle Track
   if err := json.NewDecoder(res.Body).Decode(&trackSingle); err != nil {
      return "", err
   }
   for _, transcoding := range trackSingle.Media.Transcodings {
      if strings.ToLower(transcoding.Format.Protocol) == "progressive" {
         mediaURL, err := c.getMediaURL(transcoding.URL)
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

// Track represents the JSON response of a track's info
type Track struct {
   // Media contains an array of transcoding for a track
   Media struct {
      // Transcoding contains information about the transcoding of a track
      Transcodings []struct {
         // TranscodingFormat contains the protocol by which the track is
         // delivered ("progressive" or "HLS"), and the mime type of the track
         Format struct {
            Protocol string
         }
         URL string
      }
   }
}
