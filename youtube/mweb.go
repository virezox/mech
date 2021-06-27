package youtube
import "encoding/json"


func NewMWeb(id string) (*MWeb, error) {
   res, err := newPlayer(
      id, ClientMWeb.ClientName, ClientMWeb.ClientVersion,
   ).post()
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
