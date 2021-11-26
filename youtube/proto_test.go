package youtube

import (
   "testing"
)

type test struct {
   in Params
   out string
}

var tests = []test{
   {Params{Filter: &Filter{UploadDate: UploadDateLastHour}}, "EgIIAQ=="},
   {Params{SortBy: SortByRating}, "CAE="},
}

func TestProto(t *testing.T) {
   for _, test := range tests {
      s, err := test.in.Encode()
      if err != nil {
         t.Fatal(err)
      }
      if s != test.out {
         t.Fatal(s)
      }
   }
}
