package youtube

import (
   "bytes"
   "encoding/json"
   "errors"
   "fmt"
   "github.com/robertkrimen/otto"
   "net/http"
   "net/url"
   "os"
   "path/filepath"
   "regexp"
)

const Origin = "https://www.youtube.com"

func decrypt(sig string, js []byte) (string, error) {
   /*
May 5 2021:
var uy={bH:function(a,b){a.splice(0,b)},
Fg:function(a,b){var c=a[0];a[0]=a[b%a.length];a[b%a.length]=c},
S6:function(a){a.reverse()}};
vy=function(a){a=a.split("");uy.bH(a,3);uy.Fg(a,7);uy.Fg(a,50);uy.S6(a,71);uy.bH(a,2);uy.S6(a,80);uy.Fg(a,38);return a.join("")};

May 4 2021:
var uy={an:function(a){a.reverse()},
gN:function(a,b){a.splice(0,b)},
J4:function(a,b){var c=a[0];a[0]=a[b%a.length];a[b%a.length]=c}};
vy=function(a){a=a.split("");uy.gN(a,2);uy.J4(a,47);uy.gN(a,1);uy.an(a,49);uy.gN(a,2);uy.J4(a,4);uy.an(a,71);uy.J4(a,15);uy.J4(a,40);return a.join("")};

May 3 2021:
var uy={VP:function(a){a.reverse()},
eG:function(a,b){var c=a[0];a[0]=a[b%a.length];a[b%a.length]=c},
li:function(a,b){a.splice(0,b)}};
vy=function(a){a=a.split("");uy.eG(a,50);uy.eG(a,48);uy.eG(a,23);uy.eG(a,31);return a.join("")};
   */
   child := regexp.MustCompile(`\n[^.]+\.split\(""\);[^\n]+`).Find(js)
   // child name
   childName := regexp.MustCompile(`\w+`).Find(child)
   // parent name
   parentName := regexp.MustCompile(`;\w+`).Find(child)[1:]
   // parent
   parent := regexp.MustCompile(
      fmt.Sprintf(`var %s=\S+\n[^\n]+\n[^}]+}};`, parentName),
   ).Find(js)
   // run
   vm := otto.New()
   vm.Run(string(parent) + string(child))
   value, err := vm.Call(string(childName), nil, sig)
   if err != nil {
      return "", fmt.Errorf("parent %q %v", parent, err)
   }
   return value.String(), nil
}

func get(addr string) (*bytes.Buffer, error) {
   fmt.Println("Get", addr)
   res, err := http.Get(addr)
   if err != nil { return nil, err }
   defer res.Body.Close()
   buf := new(bytes.Buffer)
   buf.ReadFrom(res.Body)
   return buf, nil
}

func getBaseJs(update bool) ([]byte, error) {
   cache, err := os.UserCacheDir()
   if err != nil { return nil, err }
   cache = filepath.Join(cache, "youtube")
   play := filepath.Join(cache, "base.js")
   if update {
      buf, err := get(Origin + "/iframe_api")
      if err != nil { return nil, err }
      id := regexp.MustCompile(`/player\\/\w+`).Find(buf.Bytes())[9:]
      base := fmt.Sprintf("/s/player/%s/player_ias.vflset/en_US/base.js", id)
      buf, err = get(Origin + base)
      if err != nil { return nil, err }
      os.Mkdir(cache, os.ModeDir)
      file, err := os.Create(play)
      if err != nil { return nil, err }
      defer file.Close()
      file.ReadFrom(buf)
   }
   return os.ReadFile(play)
}

type Format struct {
   Bitrate int
   Height int
   Itag int
   MimeType string
   SignatureCipher string
   URL string
}

func (v Video) NewFormat(itag int) (Format, error) {
   for _, format := range v.StreamingData.AdaptiveFormats {
      if format.Itag == itag { return format, nil }
   }
   return Format{}, errors.New("itag not found")
}

func (f Format) NewRequest(update bool) (*http.Request, error) {
   var req *http.Request
   if f.URL != "" {
      var err error
      req, err = http.NewRequest("GET", f.URL, nil)
      if err != nil { return nil, err }
   } else {
      val, err := url.ParseQuery(f.SignatureCipher)
      if err != nil { return nil, err }
      baseJs, err := getBaseJs(update)
      if err != nil { return nil, err }
      sig, err := decrypt(val.Get("s"), baseJs)
      if err != nil { return nil, err }
      req, err = http.NewRequest("GET", val.Get("url"), nil)
      if err != nil { return nil, err }
      val = req.URL.Query()
      val.Set("sig", sig)
      req.URL.RawQuery = val.Encode()
   }
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
   info, err := url.Parse(Origin + "/get_video_info")
   if err != nil {
      return Video{}, err
   }
   val := info.Query()
   val.Set("eurl", Origin)
   val.Set("video_id", id)
   info.RawQuery = val.Encode()
   buf, err := get(info.String())
   if err != nil {
      return Video{}, err
   }
   info.RawQuery = buf.String()
   var (
      play = info.Query().Get("player_response")
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
