package youtube
import "encoding/json"

const VersionWeb = "1.19700101"

type Microformat struct {
   PlayerMicroformatRenderer `json:"playerMicroformatRenderer"`
}

type PlayerMicroformatRenderer struct {
   AvailableCountries []string
   PublishDate string
}

type Web struct {
   Microformat `json:"microformat"`
   VideoDetails `json:"videoDetails"`
}

func NewWeb(id string) (*Web, error) {
   res, err := video(id, "WEB", VersionWeb).post()
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   w := new(Web)
   json.NewDecoder(res.Body).Decode(w)
   return w, nil
}
