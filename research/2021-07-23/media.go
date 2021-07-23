package media

import (
   "encoding/json"
   "github.com/89z/mech"
)

func audio(doc *mech.Node) string {
   script := doc.ByAttr("type", "application/ld+json")
   for script.Scan() {
      text := []byte(script.Text())
      var audio struct {
         ContentURL string
      }
      json.Unmarshal(text, &audio)
      if audio.ContentURL != "" {
         return audio.ContentURL
      }
   }
   return ""
}

func image(doc *mech.Node) string {
   img := doc.ByAttr("property", "og:image")
   img.Scan()
   return img.Attr("content")
}
