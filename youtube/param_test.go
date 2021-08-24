package youtube
import "testing"

func TestParam(t *testing.T) {
   m := Params["TYPE"]["Video"]
   if s := m.Encode(); s != "EgIQAQ==" {
      t.Fatal(s)
   }
}
