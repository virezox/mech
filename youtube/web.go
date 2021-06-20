package youtube
import "encoding/json"

type Web struct {
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

func NewWeb(id string) (Web, error) {
   res, err := post(id, "WEB", "1.19700101")
   if err != nil {
      return Web{}, err
   }
   defer res.Body.Close()
   var w Web
   json.NewDecoder(res.Body).Decode(&w)
   return w, nil
}

func (w Web) Author() string {
   return w.VideoDetails.Author
}

func (w Web) Countries() []string {
   return w.Microformat.PlayerMicroformatRenderer.AvailableCountries
}

func (w Web) Description() string {
   return w.VideoDetails.ShortDescription
}

func (w Web) PublishDate() string {
   return w.Microformat.PlayerMicroformatRenderer.PublishDate
}

func (w Web) Title() string {
   return w.VideoDetails.Title
}

func (w Web) ViewCount() int {
   return w.VideoDetails.ViewCount
}
