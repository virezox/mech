package media

import (
   "fmt"
   "github.com/89z/mech"
   "net/http"
   "testing"
)

const addr =
   "http://nytimes.com/2021/07/14/podcasts/the-daily" +
   "/heat-wave-climate-change-pacific-northwest.html"

func TestMedia(t *testing.T) {
   fmt.Println(addr)
   res, err := http.Get(addr)
   if err != nil {
      t.Fatal(err)
   }
   defer res.Body.Close()
   doc, err := mech.Parse(res.Body)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Println(image(doc))
   fmt.Println(audio(doc))
}
