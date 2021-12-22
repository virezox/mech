package apple

import (
   "fmt"
   "testing"
   "time"
)

const origin = "https://podcasts.apple.com"

var podcasts = []string{
   "/gb/podcast/theology-in-the-raw/id1018952191?i=1000544808500",
   "/podcast/4-on-the-horn-with-adam-kovacevich/id1585883787?i=1000539806607",
}

func TestApple(t *testing.T) {
   for _, pod := range podcasts {
      aud, err := NewAudio(origin + pod)
      if err != nil {
         t.Fatal(err)
      }
      fmt.Printf("%+v\n", aud)
      time.Sleep(time.Second)
   }
}
