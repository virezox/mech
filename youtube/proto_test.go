package youtube
import "testing"

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
