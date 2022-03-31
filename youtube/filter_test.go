package youtube

import (
   "testing"
)

func TestFilter(t *testing.T) {
   filter := NewFilter()
   filter.Features(Features["Subtitles/CC"])
   param := NewParams()
   param.Filter(filter)
   enc := param.Encode()
   if enc != "EgIoAQ==" {
      t.Fatal(enc)
   }
}

func TestSort(t *testing.T) {
   param := NewParams()
   param.SortBy(SortBy["Rating"])
   enc := param.Encode()
   if enc != "CAE=" {
      t.Fatal(enc)
   }
}
