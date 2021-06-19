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
   fmt.Println(invert, "GET", reset, req.URL)
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

func (p Player) Author() string {
   return p.VideoDetails.Author
}

func (p Player) Countries() []string {
   return p.Microformat.PlayerMicroformatRenderer.AvailableCountries
}

func (p Player) Description() string {
   return p.VideoDetails.ShortDescription
}

func (p Player) Formats() []Format {
   var formats []Format
   for _, format := range p.StreamingData.AdaptiveFormats {
      if format.ContentLength > 0 {
         formats = append(formats, format)
      }
   }
   return formats
}

func (p Player) NewFormat(itag int) (Format, error) {
   for _, format := range p.Formats() {
      if format.Itag == itag {
         return format, nil
      }
   }
   return Format{}, fmt.Errorf("itag %v", itag)
}

func (p Player) PublishDate() string {
   return p.Microformat.PlayerMicroformatRenderer.PublishDate
}

func (p Player) Title() string {
   return p.VideoDetails.Title
}

func (p Player) ViewCount() int {
   return p.VideoDetails.ViewCount
}
