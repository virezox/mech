package youtube

import (
   "testing"
)

func TestProtoFilter(t *testing.T) {
   filter := NewFilter()
   filter.UploadDate(UploadDate["Last hour"])
   param := NewParams()
   param.Filter(filter)
   enc := param.Encode()
   if enc != "EgIIAQ==" {
      t.Fatal(enc)
   }
}

func TestProtoSort(t *testing.T) {
   param := NewParams()
   param.SortBy(SortBy["Rating"])
   enc := param.Encode()
   if enc != "CAE=" {
      t.Fatal(enc)
   }
}
