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

const (
   base = "https://www.youtube.com/s/player/%v/player_ias.vflset/en_US/base.js"
   iframe = "https://www.youtube.com/iframe_api"
   info = "https://www.youtube.com/get_video_info"
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
func decrypt(sig, baseJs []byte) error {
   // get line
   line := regexp.MustCompile(`\.split\(""\);[^\n]+`).Find(baseJs)
   // get swaps
   for _, match := range regexp.MustCompile(`\d+`).FindAll(line, -1) {
      pos, err := strconv.Atoi(string(match))
      if err != nil { return err }
      pos %= len(sig)
      sig[0], sig[pos] = sig[pos], sig[0]
   }
   return nil
}

func getBaseJs() ([]byte, error) {
   fmt.Println("Get", iframe)
   res, err := http.Get(iframe)
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
      fmt.Println("Get", fmt.Sprintf(base, id))
      res, err := http.Get(fmt.Sprintf(base, id))
      if err != nil { return nil, err }
      defer res.Body.Close()
      file, err := os.Create(play)
      if err != nil { return nil, err }
      defer file.Close()
      file.ReadFrom(res.Body)
   } else if err != nil {
      return nil, err
   } else {
      fmt.Println("Exist", play)
   }
   return os.ReadFile(play)
}

type Format struct {
   Bitrate int
   Height int
   Itag int
   MimeType string
   SignatureCipher string
}

func (v Video) NewFormat(itag int) (Format, error) {
   for _, format := range v.StreamingData.AdaptiveFormats {
      if format.Itag == itag { return format, nil }
   }
   return Format{}, errors.New("itag not found")
}

// NewRequest returns the url for a specific format
func (f Format) NewRequest() (*http.Request, error) {
   val, err := url.ParseQuery(f.SignatureCipher)
   if err != nil { return nil, err }
   sig := []byte(val.Get("s"))
   // get player
   baseJs, err := getBaseJs()
   if err != nil { return nil, err }
   // decrypt
   err = decrypt(sig, baseJs)
   if err != nil { return nil, err }
   req, err := http.NewRequest("GET", val.Get("url"), nil)
   if err != nil { return nil, err }
   val = req.URL.Query()
   val.Set("sig", string(sig))
   req.URL.RawQuery = val.Encode()
   req.Header.Set("Range", "bytes=0-")
   return req, nil
}

type Video struct {
   StreamingData struct {
      AdaptiveFormats []Format
   }
   Microformat struct {
      PlayerMicroformatRenderer struct {
         PublishDate string
      }
   }
   VideoDetails struct {
      Author string
      ShortDescription string
      Title string
      ViewCount int `json:"viewCount,string"`
   }
}

// NewVideo fetches video metadata
func NewVideo(id string) (Video, error) {
   req, err := http.NewRequest(http.MethodGet, info, nil)
   if err != nil {
      return Video{}, err
   }
   val := req.URL.Query()
   val.Set("video_id", id)
   val.Set("eurl", info)
   req.URL.RawQuery = val.Encode()
   fmt.Println("Get", req.URL)
   res, err := new(http.Client).Do(req)
   if err != nil {
      return Video{}, err
   }
   defer res.Body.Close()
   body, err := io.ReadAll(res.Body)
   if err != nil {
      return Video{}, err
   }
   val, err = url.ParseQuery(string(body))
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

func (v Video) Author() string { return v.VideoDetails.Author }

func (v Video) Description() string { return v.VideoDetails.ShortDescription }

func (v Video) PublishDate() string {
   return v.Microformat.PlayerMicroformatRenderer.PublishDate
}

func (v Video) Title() string { return v.VideoDetails.Title }

func (v Video) ViewCount() int { return v.VideoDetails.ViewCount }
