package youtube

import (
   "encoding/json"
   "errors"
   "fmt"
   "io"
   "net/http"
   "net/url"
)

const API = "https://www.youtube.com/get_video_info"

func readAll(addr string) ([]byte, error) {
   println("Get", addr)
   res, err := http.Get(addr)
   if err != nil { return nil, err }
   defer res.Body.Close()
   return io.ReadAll(res.Body)
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
      ViewCount int `json:",string"`
   }
}

// NewVideo fetches video metadata
func NewVideo(id string) (Video, error) {
   val := make(url.Values)
   val.Set("eurl", API)
   val.Set("video_id", id)
   body, err := readAll(API + "?" + val.Encode())
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

func (v Video) Description() string { return v.VideoDetails.ShortDescription }

// GetStream returns the url for a specific format
func (v Video) GetStream(itag int) (string, error) {
   if len(v.StreamingData.AdaptiveFormats) == 0 {
      return "", errors.New("AdaptiveFormats empty")
   }
   var cipher string
   for _, format := range v.StreamingData.AdaptiveFormats {
      if format.Itag == itag { cipher = format.SignatureCipher }
   }
   query, err := url.ParseQuery(cipher)
   if err != nil { return "", err }
   operations, err := parseDecipherOps(v.VideoDetails.VideoId)
   if err != nil { return "", err }
   // apply operations
   bs := []byte(query.Get("s"))
   for _, op := range operations {
      bs = op(bs)
   }
   return fmt.Sprintf("%s&%s=%s", query.Get("url"), query.Get("sp"), bs), nil
}

func (v Video) PublishDate() string {
   return v.Microformat.PlayerMicroformatRenderer.PublishDate
}

func (v Video) Title() string { return v.VideoDetails.Title }

func (v Video) ViewCount() int { return v.VideoDetails.ViewCount }
