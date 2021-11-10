package youtube

import (
   "testing"
)

func TestFilter(t *testing.T) {
   p := Params{
      Filter: &Filter{UploadDate: UploadDateLastHour},
   }
   s, err := p.Encode()
   if err != nil {
      t.Fatal(err)
   }
   if s != "EgIIAQ==" {
      t.Fatal(s)
   }
}

func TestSort(t *testing.T) {
   p := Params{SortBy: SortByRating}
   s, err := p.Encode()
   if err != nil {
      t.Fatal(err)
   }
   if s != "CAE=" {
      t.Fatal(err)
   }
}
