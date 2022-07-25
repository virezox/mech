package dash

import (
   "encoding/xml"
   "fmt"
   "os"
   "testing"
)

func Test_DASH(t *testing.T) {
   file, err := os.Open("manifest.mpd")
   if err != nil {
      t.Fatal(err)
   }
   defer file.Close()
   var pre Presentation
   if err := xml.NewDecoder(file).Decode(&pre); err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", pre.Period)
}
