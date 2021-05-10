package youtube

import (
   "bytes"
   "encoding/json"
   "fmt"
   "io"
   "net/http"
   "net/url"
)

const (
   Origin = "https://www.youtube.com"
   chunk = 10_000_000
   invert = "\x1b[7m"
   reset = "\x1b[m"
)

func httpGet(addr string) (*bytes.Buffer, error) {
   fmt.Println(invert, "GET", reset, addr)
   res, err := http.Get(addr)
   if err != nil { return nil, err }
   defer res.Body.Close()
   buf := new(bytes.Buffer)
   buf.ReadFrom(res.Body)
   return buf, nil
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

func (v Video) NewFormat(itag int) (Format, error) {
   for _, format := range v.StreamingData.AdaptiveFormats {
      if format.Itag == itag { return format, nil }
   }
   return Format{}, fmt.Errorf("itag %v", itag)
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

func (v Video) PublishDate() string {
   return v.Microformat.PlayerMicroformatRenderer.PublishDate
}

func (v Video) Title() string { return v.VideoDetails.Title }

func (v Video) ViewCount() int { return v.VideoDetails.ViewCount }

func (f Format) Write(w io.Writer, update bool) error {
   var req *http.Request
   if f.URL != "" {
      var err error
      req, err = http.NewRequest("GET", f.URL, nil)
      if err != nil { return err }
   } else {
      val, err := url.ParseQuery(f.SignatureCipher)
      if err != nil { return err }
      baseJs, err := getBaseJs(update)
      if err != nil { return err }
      sig, err := decrypt(val.Get("s"), baseJs)
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
      io.Copy(w, res.Body)
      pos += chunk
   }
   return nil
}
