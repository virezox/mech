package youtube
import "testing"

func TestSignatureCipher(t *testing.T) {
   b, err := NewBaseJS()
   if err != nil {
      t.Fatal(err)
   }
   if err := b.Get(); err != nil {
      t.Fatal(err)
   }
}
