package youtube

import (
   "encoding/json"
   "fmt"
   "io"
   "net/http"
   "regexp"
)

type Player struct {
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

func NewPlayer(id string) (Player, error) {
   req, err := http.NewRequest("GET", "https://www.youtube.com/watch", nil)
   if err != nil {
      return Player{}, err
   }
   val := req.URL.Query()
   val.Set("v", id)
   req.URL.RawQuery = val.Encode()
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return Player{}, err
   }
   defer res.Body.Close()
   body, err := io.ReadAll(res.Body)
   if err != nil {
      return Player{}, err
   }
   re := regexp.MustCompile(">var ytInitialPlayerResponse = (.+);<")
   find := re.FindSubmatch(body)
   if find == nil {
      return Player{}, fmt.Errorf("findSubmatch %v", re)
   }
   var play Player
   json.Unmarshal(find[1], &play)
   return play, nil
}

func (v Player) Author() string {
   return v.VideoDetails.Author
}

func (v Player) Countries() []string {
   return v.Microformat.PlayerMicroformatRenderer.AvailableCountries
}

func (v Player) Description() string {
   return v.VideoDetails.ShortDescription
}

func (v Player) Formats() []Format {
   var formats []Format
   for _, format := range v.StreamingData.AdaptiveFormats {
      if format.ContentLength > 0 {
         formats = append(formats, format)
      }
   }
   return formats
}

func (v Player) NewFormat(itag int) (Format, error) {
   for _, format := range v.Formats() {
      if format.Itag == itag {
         return format, nil
      }
   }
   return Format{}, fmt.Errorf("itag %v", itag)
}

func (v Player) PublishDate() string {
   return v.Microformat.PlayerMicroformatRenderer.PublishDate
}

func (v Player) Title() string {
   return v.VideoDetails.Title
}

func (v Player) ViewCount() int {
   return v.VideoDetails.ViewCount
}
