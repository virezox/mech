package youtube

import (
   "encoding/base64"
   "testing"
)

func Test_Filter(t *testing.T) {
   filter := New_Filter()
   filter.Features(Features["Subtitles/CC"])
   param := New_Params()
   param.Filter(filter)
   buf := param.Marshal()
   if base64.StdEncoding.EncodeToString(buf) != "EgIoAQ==" {
      t.Fatal(buf)
   }
}

func Test_Sort(t *testing.T) {
   param := New_Params()
   param.Sort_By(Sort_By["Rating"])
   buf := param.Marshal()
   if base64.StdEncoding.EncodeToString(buf) != "CAE=" {
      t.Fatal(buf)
   }
}
