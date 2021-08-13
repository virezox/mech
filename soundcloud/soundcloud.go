package soundcloud

import (
   "encoding/json"
   "fmt"
   "github.com/89z/mech"
   "net/http"
   "net/http/httputil"
   "os"
   "regexp"
   "strings"
)

const Origin = "https://api-v2.soundcloud.com"

func ReadClientID() (string, error) {
   cache, err := os.UserCacheDir()
   if err != nil {
      return "", err
   }
   b, err := os.ReadFile(cache + "/mech/soundcloud.js")
   if err != nil {
      return "", err
   }
   re := regexp.MustCompile(`\?client_id=([^&]+)&`)
   find := re.FindSubmatch(b)
   if find == nil {
      return "", fmt.Errorf("findSubmatch %v", re)
   }
   return string(find[1]), nil
}

// client_id appears to last for at least a year
func WriteClientID() error {
   script, err := getScript()
   if err != nil {
      return err
   }
   fmt.Println("GET", script)
   res, err := http.Get(script)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   cache, err := os.UserCacheDir()
   if err != nil {
      return err
   }
   cache += "/mech"
   os.Mkdir(cache, os.ModeDir)
   f, err := os.Create(cache + "/soundcloud.js")
   if err != nil {
      return err
   }
   defer f.Close()
   if _, err := f.ReadFrom(res.Body); err != nil {
      return err
   }
   return nil
}

func getScript() (string, error) {
   addr := "https://soundcloud.com"
   fmt.Println("GET", addr)
   res, err := http.Get(addr)
   if err != nil {
      return "", err
   }
   defer res.Body.Close()
   scan := mech.NewScanner(res.Body)
   for scan.ScanAttr("crossorigin", "") {
      src := scan.Attr("src")
      /*
      asset 49 works as well:
      https://a-v2.sndcdn.com/assets/49-4b976e4f.js
      client_id:"fSSdm5yTnDka1g0Fz1CO5Yx6z0NbeHAj"
      content-length: 1393972

      but asset 2 is smaller:
      https://a-v2.sndcdn.com/assets/2-b0e52b4d.js
      ?client_id=fSSdm5yTnDka1g0Fz1CO5Yx6z0NbeHAj&
      content-length: 922378
      */
      if strings.HasPrefix(src, "https://a-v2.sndcdn.com/assets/2-") {
         return src, nil
      }
   }
   return "", fmt.Errorf("%+v", res)
}

// MediaURLResponse is the JSON response of retrieving media information of a
// track
type Media struct {
   URL string
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
   req, err := http.NewRequest("GET", Origin + "/resolve", nil)
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
func (t Track) GetMedia(id string) (*Media, error) {
   var addr string
   for _, code := range t.Media.Transcodings {
      if code.Format.Protocol == "progressive" {
         addr = code.URL
      }
   }
   req, err := http.NewRequest("GET", addr, nil)
   if err != nil {
      return nil, err
   }
   q := req.URL.Query()
   q.Set("client_id", id)
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
   m := new(Media)
   if err := json.NewDecoder(res.Body).Decode(&m); err != nil {
      return nil, err
   }
   return m, nil
}
