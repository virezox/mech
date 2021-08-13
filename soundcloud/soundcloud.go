package soundcloud

import (
   "encoding/json"
   "fmt"
   "github.com/89z/mech"
   "io"
   "net/http"
   "net/http/httputil"
   "os"
   "regexp"
)

const (
   resolveURL = "https://api-v2.soundcloud.com/resolve"
   trackURL = "https://api-v2.soundcloud.com/tracks"
)

// Fetch a SoundCloud client ID. The basic notion of how this function works is
// that SoundCloud provides a client ID so its web app can make API requests.
// This client ID (along with other intialization data for the web app) is
// provided in a JavaScript file imported through a <script> tag in the HTML.
// This function scrapes the HTML and tries to find the URL to that JS file,
// and then scrapes the JS file to find the client ID.
func ClientID() (string, error) {
   script, err := getScript()
   if err != nil {
      return "", err
   }
   fmt.Println("GET", script)
   res, err := http.Get(script)
   if err != nil {
      return "", err
   }
   defer res.Body.Close()
   body, err := io.ReadAll(res.Body)
   if err != nil {
      return "", err
   }
   // Extract the client ID
   re := regexp.MustCompile(`\bclient_id:"([^"]+)"`)
   find := re.FindSubmatch(body)
   if find == nil {
      return "", fmt.Errorf("findSubmatch %v", re)
   }
   return string(find[1]), nil
}

func getScript() (string, error) {
   addr := "https://soundcloud.com"
   fmt.Println("GET", addr)
   res, err := http.Get(addr)
   if err != nil {
      return "", err
   }
   scan := mech.NewScanner(res.Body)
   var src string
   // The link to the JS file with the client ID looks like this:
   // <script crossorigin src="https://a-v2.sndcdn.com/assets/sdfhkjhsdkf.js">
   // Extract all the URLS that match our pattern. It seems like our desired
   // URL is imported last
   for scan.ScanAttr("crossorigin", "") {
      src = scan.Attr("src")
   }
   return src, nil
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

func NewTrack(id, addr string) (*Track, error) {
   req, err := http.NewRequest("GET", resolveURL, nil)
   if err != nil {
      return nil, err
   }
   q := req.URL.Query()
   q.Set("client_id", id)
   q.Set("url", addr)
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
   t := new(Track)
   if err := json.NewDecoder(res.Body).Decode(&t); err != nil {
      return nil, err
   }
   return t, nil
}

// The media URL is the actual link to the audio file for the track. "addr" is
// Track.Media.Transcodings[0].URL
func getMediaURL(id, addr string) (string, error) {
   req, err := http.NewRequest("GET", addr, nil)
   if err != nil {
      return "", err
   }
   q := req.URL.Query()
   q.Set("client_id", id)
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
