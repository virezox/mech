package ocr

import (
   "testing"
   "time"
)

var tests = []struct{in, out string} {
   {"png/2S7BH.png", "2S7BH"},
   {"png/79B3P.png", "79B3P"},
   {"png/7UI3H.png", "7UI3H"},
   {"png/BWD6N.png", "BWD6N"},
   {"png/CZMAO.png", "CZMAO"},
   {"png/FEXBV.png", "FEXBV"},
   {"png/R2XBP.png", "R2XBP"},
}

func TestImage(t *testing.T) {
   for _, test := range tests {
      img, err := NewImage(test.in)
      if err != nil {
         t.Fatal(err)
      }
      out := img.ParsedResults[0].ParsedText
      if out != test.out {
         t.Fatal(out)
      }
      time.Sleep(time.Second)
   }
}
