package youtube

import (
   "fmt"
   "mime"
   "testing"
)

var mimeTypes = []string{
   "audio/mp4; codecs=\"mp4a.40.2\"",
   "audio/mp4; codecs=\"mp4a.40.5\"",
   "audio/webm; codecs=\"opus\"",
   "video/3gpp; codecs=\"mp4v.20.3, mp4a.40.2\"",
   "video/mp4; codecs=\"av01.0.00M.08\"",
   "video/mp4; codecs=\"av01.0.01M.08\"",
   "video/mp4; codecs=\"av01.0.04M.08\"",
   "video/mp4; codecs=\"av01.0.05M.08\"",
   "video/mp4; codecs=\"av01.0.08M.08\"",
   "video/mp4; codecs=\"av01.0.12M.08\"",
   "video/mp4; codecs=\"avc1.42001E, mp4a.40.2\"",
   "video/mp4; codecs=\"avc1.4d400c\"",
   "video/mp4; codecs=\"avc1.4d4015\"",
   "video/mp4; codecs=\"avc1.4d401e\"",
   "video/mp4; codecs=\"avc1.4d401f\"",
   "video/mp4; codecs=\"avc1.64001F, mp4a.40.2\"",
   "video/mp4; codecs=\"avc1.640028\"",
   "video/webm; codecs=\"vp9\"",
}

func TestFormat(t *testing.T) {
   for _, mimeType := range mimeTypes {
      exts, err := mime.ExtensionsByType(mimeType)
      if err != nil {
         t.Fatal(err)
      }
      fmt.Println(mimeType, exts)
   }
}
