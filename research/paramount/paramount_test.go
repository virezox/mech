package paramount

import (
   "fmt"
   "testing"
)

const issue = "paramountplus.com/shows/aeon-flux/video/IBmGAIrEofp2vQNvSEvO5MavtZy5GOAy/aeon-flux-isthmus-crypticus"

func TestParamount(t *testing.T) {
   id := VideoID(issue)
   med, err := Media(id)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Println(med)
}
