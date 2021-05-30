package youtube

import (
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

type BaseJS struct {
   Cache string
   Create string
}

func NewBaseJS() (BaseJS, error) {
   cache, err := os.UserCacheDir()
   if err != nil {
      return BaseJS{}, err
   }
   cache = filepath.Join(cache, "youtube")
   return BaseJS{
      cache, filepath.Join(cache, "base.js"),
   }, nil
}

func (b BaseJS) Get() error {
   fmt.Println(invert, "GET", reset, Origin + "/iframe_api")
   res, err := http.Get(Origin + "/iframe_api")
   if err != nil { return err }
   defer res.Body.Close()
   body, err := io.ReadAll(res.Body)
   if err != nil { return err }
   re := regexp.MustCompile(`/player\\/(\w+)`)
   id := re.FindSubmatch(body)
   if id == nil {
      return fmt.Errorf("FindSubmatch %v", re)
   }
   os.Mkdir(b.Cache, os.ModeDir)
   file, err := os.Create(b.Create)
   if err != nil { return err }
   defer file.Close()
   get := fmt.Sprintf("/s/player/%s/player_ias.vflset/en_US/base.js", id[1])
   fmt.Println(invert, "GET", reset, Origin + get)
   if res, err := http.Get(Origin + get); err != nil {
      return err
   } else {
      defer res.Body.Close()
      if _, err := file.ReadFrom(res.Body); err != nil {
         return err
      }
   }
   return nil
}

type Format struct {
   Bitrate int64
   ContentLength int64 `json:"contentLength,string"`
   Height int
   Itag int
   MimeType string
   SignatureCipher string
   URL string
}

func (f Format) Write(w io.Writer) error {
   req, err := f.request()
   if err != nil { return err }
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
      if _, err := io.Copy(w, res.Body); err != nil {
         return err
      }
      pos += chunk
   }
   return nil
}

func (f Format) request() (*http.Request, error) {
   if f.URL != "" {
      return http.NewRequest("GET", f.URL, nil)
   }
   baseJS, err := NewBaseJS()
   if err != nil { return nil, err }
   js, err := os.ReadFile(baseJS.Create)
   if err != nil { return nil, err }
   re := `\n[^.]+\.split\(""\);.+`
   child := regexp.MustCompile(re).Find(js)
   if child == nil {
      return nil, fmt.Errorf("find %v", re)
   }
   re = `\w+`
   childName := regexp.MustCompile(re).Find(child)
   if childName == nil {
      return nil, fmt.Errorf("find %v", re)
   }
   re = `;(\w+)`
   parentName := regexp.MustCompile(re).FindSubmatch(child)
   if parentName == nil {
      return nil, fmt.Errorf("find %v", re)
   }
   re = fmt.Sprintf(`var %s=.+\n.+\n[^}]+}};`, parentName[1])
   parent := regexp.MustCompile(re).Find(js)
   if parent == nil {
      return nil, fmt.Errorf("find %v", re)
   }
   val, err := url.ParseQuery(f.SignatureCipher)
   if err != nil { return nil, err }
   vm := otto.New()
   if _, err := vm.Run(string(parent) + string(child)); err != nil {
      return nil, err
   }
   sig, err := vm.Call(string(childName), nil, val.Get("s"))
   if err != nil { return nil, err }
   req, err := http.NewRequest("GET", val.Get("url"), nil)
   if err != nil { return nil, err }
   val = req.URL.Query()
   val.Set("sig", sig.String())
   req.URL.RawQuery = val.Encode()
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
   addr, err := url.Parse(Origin + "/get_video_info")
   if err != nil {
      return Video{}, err
   }
   val := addr.Query()
   val.Set("eurl", Origin)
   val.Set("html5", "1")
   val.Set("video_id", id)
   addr.RawQuery = val.Encode()
   fmt.Println(invert, "GET", reset, addr)
   res, err := http.Get(addr.String())
   if err != nil {
      return Video{}, err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      return Video{}, fmt.Errorf("StatusCode %v", res.StatusCode)
   }
   body, err := io.ReadAll(res.Body)
   if err != nil {
      return Video{}, err
   }
   addr.RawQuery = string(body)
   var (
      play = addr.Query().Get("player_response")
      vid Video
   )
   json.Unmarshal([]byte(play), &vid)
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
