package youtube

import (
   "testing"
)

func TestFilter(t *testing.T) {
   p := param{
      Filter: &filter{UploadDate: uploadDateLastHour},
   }
   s, err := p.encode()
   if err != nil {
      t.Fatal(err)
   }
   if s != "EgIIAQ==" {
      t.Fatal(s)
   }
}

func TestSort(t *testing.T) {
   p := param{SortBy: sortByRating}
   s, err := p.encode()
   if err != nil {
      t.Fatal(err)
   }
   if s != "CAE=" {
      t.Fatal(err)
   }
}
