package youtube
import "testing"

func TestContinue(t *testing.T) {
   s, err := NewContinuation("q5UnT4Ik6KU").Encode()
   if err != nil {
      t.Fatal(err)
   }
   if s != "Eg0SC3E1VW5UNElrNktVGAYyDyINIgtxNVVuVDRJazZLVQ==" {
      t.Fatal(s)
   }
}

type test struct {
   name string
   fn func(*Param)
   want string
}

var tests = []test{
   {"Channel", (*Param).Channel, "EgIQAg=="},
   {"FourK", (*Param).FourK, "EgJwAQ=="},
   {"FourToTwentyMinutes", (*Param).FourToTwentyMinutes, "EgIYAw=="},
   {"Today", (*Param).Today, "EgIIAg=="},
   {"UploadDate", (*Param).UploadDate, "CAISAA=="},
   {"Video", (*Param).Video, "EgIQAQ=="},
}

func TestParam(t *testing.T) {
   for _, each := range tests {
      var p Param
      each.fn(&p)
      s, err := p.Encode()
      if err != nil {
         t.Fatal(err)
      }
      if s != each.want {
         t.Fatal(s)
      }
   }
}
