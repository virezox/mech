package paramount

import (
   "fmt"
   "testing"
   "time"
)

var addrs = []string{
   "paramountplus.com/movies/building-star-trek/wQH9yE_y_Dt4ekDYm3yelhhY2KXvOra_",
   "paramountplus.com/shows/aeon-flux/video/IBmGAIrEofp2vQNvSEvO5MavtZy5GOAy/aeon-flux-isthmus-crypticus",
}

func TestParamount(t *testing.T) {
   for _, addr := range addrs {
      id := GUID(addr)
      med, err := Media(id)
      if err != nil {
         t.Fatal(err)
      }
      fmt.Println(med)
      time.Sleep(time.Second)
   }
}
