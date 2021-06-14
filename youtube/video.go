package youtube

import (
   "encoding/json"
   "fmt"
   "io"
   "net/http"
   "net/url"
)

type Video struct {
   StreamingData struct {
      AdaptiveFormats []Format
   }
   Microformat struct {
      PlayerMicroformatRenderer struct {
         AvailableCountries []string
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
   if err := json.Unmarshal([]byte(play), &vid); err != nil {
      return Video{}, err
   }
   return vid, nil
}

func (v Video) Author() string {
   return v.VideoDetails.Author
}

func (v Video) Countries() []string {
   return v.Microformat.PlayerMicroformatRenderer.AvailableCountries
}

func (v Video) Description() string {
   return v.VideoDetails.ShortDescription
}

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
      if format.Itag == itag {
         return format, nil
      }
   }
   return Format{}, fmt.Errorf("itag %v", itag)
}

func (v Video) PublishDate() string {
   return v.Microformat.PlayerMicroformatRenderer.PublishDate
}

func (v Video) Title() string {
   return v.VideoDetails.Title
}

func (v Video) ViewCount() int {
   return v.VideoDetails.ViewCount
}
