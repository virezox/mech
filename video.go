package youtube

import (
   "bytes"
   "encoding/json"
   "fmt"
   "github.com/robertkrimen/otto"
   "io"
   "net/http"
   "net/url"
   "os"
   "path/filepath"
   "regexp"
)

const (
   Origin = "https://www.youtube.com"
   chunk = 10_000_000
   invert = "\x1b[7m"
   reset = "\x1b[m"
)

func decrypt(sig string, js []byte) (string, error) {
   re := `\n[^.]+\.split\(""\);[^\n]+`
   child := regexp.MustCompile(re).Find(js)
   if child == nil {
      return "", fmt.Errorf("find %v", re)
   }
   re = `\w+`
   childName := regexp.MustCompile(re).Find(child)
   if childName == nil {
      return "", fmt.Errorf("find %v", re)
   }
   re = `;\w+`
   parentName := regexp.MustCompile(re).Find(child)
   if parentName == nil {
      return "", fmt.Errorf("find %v", re)
   }
   parentName = parentName[1:]
   re = fmt.Sprintf(`var %s=[^\n]+\n[^\n]+\n[^}]+}};`, parentName)
   parent := regexp.MustCompile(re).Find(js)
   if parent == nil {
      return "", fmt.Errorf("find %v", re)
   }
   vm := otto.New()
   vm.Run(string(parent) + string(child))
   value, err := vm.Call(string(childName), nil, sig)
   if err != nil {
      return "", fmt.Errorf("parent %q %v", parent, err)
   }
   return value.String(), nil
   /*
May 7 2021:
var uy={wd:function(a,b){var c=a[0];a[0]=a[b%a.length];a[b%a.length]=c},
kI:function(a){a.reverse()},
NY:function(a,b){a.splice(0,b)}};wy.prototype.set=function(a,b){this.i[a]!==b&&(this.i[a]=b,this.url="")};
vy=function(a){a=a.split("");uy.wd(a,41);uy.NY(a,3);uy.kI(a,41);uy.NY(a,2);uy.kI(a,5);uy.wd(a,62);uy.NY(a,3);uy.wd(a,69);uy.NY(a,2);return a.join("")};

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
}

func httpGet(addr string) (*bytes.Buffer, error) {
   fmt.Println(invert, "GET", reset, addr)
   res, err := http.Get(addr)
   if err != nil { return nil, err }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      return nil, fmt.Errorf("StatusCode %v", res.StatusCode)
   }
   buf := new(bytes.Buffer)
   buf.ReadFrom(res.Body)
   return buf, nil
}

type BaseJS struct {
   Cache string
   Create string
}

func NewBaseJS() (BaseJS, error) {
   var (
      b BaseJS
      err error
   )
   b.Cache, err = os.UserCacheDir()
   if err != nil {
      return BaseJS{}, err
   }
   b.Cache = filepath.Join(b.Cache, "youtube")
   b.Create = filepath.Join(b.Cache, "base.js")
   return b, nil
}

func (b BaseJS) Get() error {
   buf, err := httpGet(Origin + "/iframe_api")
   if err != nil { return err }
   re := regexp.MustCompile(`/player\\/\w+`)
   id := re.Find(buf.Bytes())
   if id == nil {
      return fmt.Errorf("find %v", re)
   }
   get := fmt.Sprintf("/s/player/%s/player_ias.vflset/en_US/base.js", id[9:])
   buf, err = httpGet(Origin + get)
   if err != nil { return err }
   os.Mkdir(b.Cache, os.ModeDir)
   file, err := os.Create(b.Create)
   if err != nil { return err }
   defer file.Close()
   _, err = file.ReadFrom(buf)
   return err
}

type Format struct {
   Bitrate int
   ContentLength int64 `json:"contentLength,string"`
   Height int
   Itag int
   MimeType string
   SignatureCipher string
   URL string
}

func (f Format) Write(w io.Writer) error {
   var req *http.Request
   if f.URL != "" {
      var err error
      req, err = http.NewRequest("GET", f.URL, nil)
      if err != nil { return err }
   } else {
      val, err := url.ParseQuery(f.SignatureCipher)
      if err != nil { return err }
      baseJS, err := NewBaseJS()
      if err != nil { return err }
      create, err := os.ReadFile(baseJS.Create)
      if err != nil { return err }
      sig, err := decrypt(val.Get("s"), create)
      if err != nil { return err }
      req, err = http.NewRequest("GET", val.Get("url"), nil)
      if err != nil { return err }
      val = req.URL.Query()
      val.Set("sig", sig)
      req.URL.RawQuery = val.Encode()
   }
   var pos int64
   fmt.Println(invert, "GET", reset, req.URL)
   for pos < f.ContentLength {
      bytes := fmt.Sprintf("bytes=%v-%v", pos, pos+chunk-1)
      req.Header.Set("Range", bytes)
      fmt.Println(bytes)
      res, err := new(http.Client).Do(req)
      if err != nil { return err }
      defer res.Body.Close()
      if res.StatusCode != http.StatusPartialContent {
         return fmt.Errorf("StatusCode %v", res.StatusCode)
      }
      io.Copy(w, res.Body)
      pos += chunk
   }
   return nil
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
   buf, err := httpGet(info.String())
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

func (v Video) Formats() []Format {
   var formats []Format
   for _, format := range v.StreamingData.AdaptiveFormats {
      if format.ContentLength > 0 {
         formats = append(formats, format)
      }
   }
   return formats
}

func (v Video) NewFormat(itag int) (Format, error) {
   for _, format := range v.Formats() {
      if format.Itag == itag { return format, nil }
   }
   return Format{}, fmt.Errorf("itag %v", itag)
}

func (v Video) PublishDate() string {
   return v.Microformat.PlayerMicroformatRenderer.PublishDate
}

func (v Video) Title() string { return v.VideoDetails.Title }

func (v Video) ViewCount() int { return v.VideoDetails.ViewCount }
