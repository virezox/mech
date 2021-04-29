package youtube
import "testing"

func TestVideo(t *testing.T) {
   s := "https://www.youtube.com/watch?v=BnEn7X3Pr7o"
   v, err := NewVideo(s)
   if err != nil {
      t.Error(err)
   }
   if v.DASHManifestURL == "" {
      t.Error()
   }
}
