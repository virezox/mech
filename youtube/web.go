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
   r, err := newPlayer(id, "WEB", VersionWeb).post()
   if err != nil {
      return nil, err
   }
   defer r.Body.Close()
   w := new(Web)
   if err := json.NewDecoder(r.Body).Decode(w); err != nil {
      return nil, err
   }
   return w, nil
}
