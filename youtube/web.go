package youtube

import (
   "encoding/json"
   "fmt"
   "net/http"
   "strings"
)

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
   body := fmt.Sprintf(`
   {
      "videoId": %q, "context": {
         "client": {"clientName": "WEB", "clientVersion": "1.19700101"}
      }
   }
   `, id)
   req, err := http.NewRequest(
      "POST", PlayerAPI, strings.NewReader(body),
   )
   if err != nil {
      return Web{}, err
   }
   val := req.URL.Query()
   val.Set("key", "AIzaSyAO_FJ2SlqU8Q4STEHLGCilw_Y9_11qcW8")
   req.URL.RawQuery = val.Encode()
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return Web{}, err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      return Web{}, fmt.Errorf("status %v", res.Status)
   }
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
