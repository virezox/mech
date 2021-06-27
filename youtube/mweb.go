package youtube
import "encoding/json"

const VersionMWeb = "2.19700101"

type Microformat struct {
   PlayerMicroformatRenderer `json:"playerMicroformatRenderer"`
}

type PlayerMicroformatRenderer struct {
   AvailableCountries []string
   PublishDate string
}

type MWeb struct {
   Microformat `json:"microformat"`
   VideoDetails `json:"videoDetails"`
}

func NewMWeb(id string) (*MWeb, error) {
   res, err := newPlayer(id, "MWEB", VersionMWeb).post()
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   mw := new(MWeb)
   if err := json.NewDecoder(res.Body).Decode(mw); err != nil {
      return nil, err
   }
   return mw, nil
}
