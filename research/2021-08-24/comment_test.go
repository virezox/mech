package comment
import "testing"

func TestProto(t *testing.T) {
   s := continuation("q5UnT4Ik6KU")
   if s != "Eg0SC3E1VW5UNElrNktVGAYyDyINIgtxNVVuVDRJazZLVQ==" {
      t.Fatal(s)
   }
}
