package apple

import (
   "fmt"
   "testing"
)

const addr =
   "https://podcasts.apple.com/podcast/4-on-the-horn-with-adam-kovacevich/" +
   "id1585883787?i=1000539806607"

func TestApple(t *testing.T) {
   Verbose(true)
   aud, err := NewAudio(addr)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", aud)
}
