package youtube
import "encoding/json"

var ClientMWeb = Client{"MWEB", "2.19700101"}

type MWeb struct {
   Microformat `json:"microformat"`
   VideoDetails `json:"videoDetails"`
}

func NewMWeb(id string) (*MWeb, error) {
   res, err := ClientMWeb.newPlayer(id).post()
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

type Microformat struct {
   PlayerMicroformatRenderer `json:"playerMicroformatRenderer"`
}

type PlayerMicroformatRenderer struct {
   AvailableCountries []string
   PublishDate string
}
