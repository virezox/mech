package youtube

import (
   "encoding/json"
   "fmt"
   "io"
   "net/http"
   "regexp"
)

type Player struct {
   Video
}

func NewPlayer(id string) error {
   req, err := http.NewRequest("GET", "https://www.youtube.com/watch", nil)
   if err != nil {
      return err
   }
   val := req.URL.Query()
   val.Set("v", id)
   req.URL.RawQuery = val.Encode()
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   body, err := io.ReadAll(res.Body)
   if err != nil {
      return err
   }
   re := regexp.MustCompile(">var ytInitialPlayerResponse = (.+);<")
   find := re.FindSubmatch(body)
   if find == nil {
      return fmt.Errorf("findSubmatch %v", re)
   }
   var vid Video
   json.Unmarshal(find[1], &vid)
   fmt.Printf("%+v\n", vid)
   return nil
}
