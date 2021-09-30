package instagram

import (
   "fmt"
   "testing"
   "time"
)

var tests = []string{
   // multiple image
   "CT-cnxGhvvO",
   // single image
   //"CUV3UsjL1WM",
   // video
   "CSL6jroDj-Q",
   //"CUWBw4TM6Np",
}

func TestSidecar(t *testing.T) {
   for _, test := range tests {
      m, err := NewMedia(test)
      if err != nil {
         t.Fatal(err)
      }
      fmt.Printf("%+v\n", m)
      time.Sleep(time.Second)
   }
}
