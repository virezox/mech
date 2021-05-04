package youtube

import (
   "encoding/json"
   "errors"
   "fmt"
   "io"
   "net/http"
   "net/url"
   "os"
   "path/filepath"
   "regexp"
   "strconv"
)

/*
Current logic is based on this input:

var uy={VP:function(a){a.reverse()},
eG:function(a,b){var c=a[0];a[0]=a[b%a.length];a[b%a.length]=c},
li:function(a,b){a.splice(0,b)}};
vy=function(a){a=a.split("");uy.eG(a,50);uy.eG(a,48);uy.eG(a,23);uy.eG(a,31);return a.join("")};

if this fails in the future, we should keep a record of all failed cases, to
keep from repeating a mistake.
*/
func decrypt(sig, body []byte) error {
   // get line
   line := regexp.MustCompile(`\.split\(""\);[^\n]+`).Find(body)
   // get swaps
   matches := regexp.MustCompile(`\d+`).FindAll(line, -1)
   for _, match := range matches {
      pos, err := strconv.Atoi(string(match))
      if err != nil { return err }
      pos %= len(sig)
      sig[0], sig[pos] = sig[pos], sig[0]
   }
   return nil
}

func getPlayer() ([]byte, error) {
   yt := newYouTube()
   yt.Path = "iframe_api"
   println("Get", yt.String())
   res, err := http.Get(yt.String())
   if err != nil { return nil, err }
   defer res.Body.Close()
   body, err := io.ReadAll(res.Body)
   if err != nil { return nil, err }
   match := regexp.MustCompile(`/player\\/(\w+)`).FindSubmatch(body)
   id := string(match[1])
   // cache
   cache, err := os.UserCacheDir()
   if err != nil { return nil, err }
   cache = filepath.Join(cache, "youtube")
   play := filepath.Join(cache, id + ".js")
   _, err = os.Stat(play)
   if os.IsNotExist(err) {
      os.Mkdir(cache, os.ModeDir)
      yt.Path = fmt.Sprintf("s/player/%v/player_ias.vflset/en_US/base.js", id)
      res, err := http.Get(yt.String())
      if err != nil { return nil, err }
      defer res.Body.Close()
      file, err := os.Create(play)
      if err != nil { return nil, err }
      defer file.Close()
      file.ReadFrom(res.Body)
   } else if err != nil {
      return nil, err
   } else {
      println("Exist", play)
   }
   return os.ReadFile(play)
}

type Video struct {
   StreamingData struct {
      AdaptiveFormats []struct {
         Bitrate int
         Height int
         Itag int
         MimeType string
         SignatureCipher string
      }
   }
   Microformat struct {
      PlayerMicroformatRenderer struct {
         PublishDate string
      }
   }
   VideoDetails struct {
      ShortDescription string
      Title string
      VideoId string
      ViewCount int `json:"viewCount,string"`
   }
}

// NewVideo fetches video metadata
func NewVideo(id string) (Video, error) {
   yt := newYouTube()
   yt.Set("eurl", yt.String())
   yt.Set("video_id", id)
   yt.RawQuery = yt.Encode()
   yt.Path = "get_video_info"
   println("Get", yt.String())
   res, err := http.Get(yt.String())
   if err != nil {
      return Video{}, err
   }
   defer res.Body.Close()
   body, err := io.ReadAll(res.Body)
   if err != nil {
      return Video{}, err
   }
   val, err := url.ParseQuery(string(body))
   if err != nil {
      return Video{}, err
   }
   var (
      play = val.Get("player_response")
      vid Video
   )
   err = json.Unmarshal([]byte(play), &vid)
   if err != nil {
      return Video{}, err
   }
   return vid, nil
}

func (v Video) Description() string { return v.VideoDetails.ShortDescription }

// GetStream returns the url for a specific format
func (v Video) GetStream(itag int) (string, error) {
   if len(v.StreamingData.AdaptiveFormats) == 0 {
      return "", errors.New("AdaptiveFormats empty")
   }
   // get cipher text
   cipher, err := v.signatureCipher(itag)
   if err != nil { return "", err }
   query, err := url.ParseQuery(cipher)
   if err != nil { return "", err }
   sig := []byte(query.Get("s"))
   // get player
   body, err := getPlayer()
   if err != nil { return "", err }
   // decrypt
   err = decrypt(sig, body)
   if err != nil { return "", err }
   return query.Get("url") + "&sig=" + string(sig), nil
}

func (v Video) PublishDate() string {
   return v.Microformat.PlayerMicroformatRenderer.PublishDate
}

func (v Video) Title() string { return v.VideoDetails.Title }

func (v Video) ViewCount() int { return v.VideoDetails.ViewCount }

func (v Video) signatureCipher(itag int) (string, error) {
   for _, format := range v.StreamingData.AdaptiveFormats {
      if format.Itag == itag { return format.SignatureCipher, nil }
   }
   return "", errors.New("itag not found")
}

type youTube struct {
   url.URL
   url.Values
}

func newYouTube() youTube {
   return youTube{
      url.URL{Scheme: "https", Host: "www.youtube.com"}, make(url.Values),
   }
}
