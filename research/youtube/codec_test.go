package youtube

import (
   "fmt"
   "testing"
)

var tests = []string{
   "video/mp4; codecs=\"av01.0.05M.08\"",
   "video/mp4; codecs=\"avc1.4d401f\"",
   "video/webm; codecs=\"vp9\"",
}

func TestCodec(t *testing.T) {
   for _, test := range tests {
      codec, err := getCodec(test)
      if err != nil {
         t.Fatal(err)
      }
      fmt.Println(codec)
   }
}
