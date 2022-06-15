package youtube

import (
   "encoding/base64"
   "testing"
)

func TestFilter(t *testing.T) {
   filter := NewFilter()
   filter.Features(Features["Subtitles/CC"])
   param := NewParams()
   param.Filter(filter)
   buf, err := param.MarshalBinary()
   if err != nil {
      t.Fatal(err)
   }
   if base64.StdEncoding.EncodeToString(buf) != "EgIoAQ==" {
      t.Fatal(buf)
   }
}

func TestSort(t *testing.T) {
   param := NewParams()
   param.SortBy(SortBy["Rating"])
   buf, err := param.MarshalBinary()
   if err != nil {
      t.Fatal(err)
   }
   if base64.StdEncoding.EncodeToString(buf) != "CAE=" {
      t.Fatal(buf)
   }
}
