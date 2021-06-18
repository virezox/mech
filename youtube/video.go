package youtube

import (
   "encoding/json"
   "fmt"
   "io"
   "net/http"
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
   req, err := http.NewRequest("GET", Origin + "/get_video_info", nil)
   if err != nil {
      return Video{}, err
   }
   val := req.URL.Query()
   val.Set("c", "TVHTML5")
   val.Set("cver", "4.19700101")
   val.Set("eurl", Origin)
   val.Set("html5", "1")
   val.Set("video_id", id)
   req.URL.RawQuery = val.Encode()
   fmt.Println(invert, "GET", reset, req.URL)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return Video{}, err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      return Video{}, fmt.Errorf("status %v", res.Status)
   }
   body, err := io.ReadAll(res.Body)
   if err != nil {
      return Video{}, err
   }
   req.URL.RawQuery = string(body)
   play := req.URL.Query().Get("player_response")
   var vid Video
   json.Unmarshal([]byte(play), &vid)
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
