package paramount

import (
   "fmt"
   "testing"
   "time"
)

var guids = []string{
   // paramountplus.com/movies/building-star-trek/wQH9yE_y_Dt4ekDYm3yelhhY2KXvOra_
   "wQH9yE_y_Dt4ekDYm3yelhhY2KXvOra_",
   // paramountplus.com/shows/aeon-flux/video/IBmGAIrEofp2vQNvSEvO5MavtZy5GOAy/aeon-flux-isthmus-crypticus
   "IBmGAIrEofp2vQNvSEvO5MavtZy5GOAy",
}

func TestParamount(t *testing.T) {
   for _, guid := range guids {
      media, err := NewMedia(guid)
      if err != nil {
         t.Fatal(err)
      }
      video, err := media.Video()
      if err != nil {
         t.Fatal(err)
      }
      for _, param := range video.Param {
         fmt.Printf("%+v\n", param)
      }
      time.Sleep(time.Second)
   }
}
