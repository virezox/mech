package main

import (
   "encoding/json"
   "fmt"
   "github.com/89z/mech"
   "net/http"
)

var tests = []string{
   "http://hrp.org.uk" +
   "/tower-of-london/history-and-stories/tower-of-london-prison",
   "http://independent.co.uk" +
   "/news/world/americas/us-politics/b1873357.html",
   "http://nytimes.com/2021/07/14/podcasts/the-daily" +
   "/heat-wave-climate-change-pacific-northwest.html",
   "http://youtube.com/watch?v=LxK5Ocehj10",
}

func video(doc *mech.Node) string {
   script := doc.ByAttr("type", "application/ld+json")
   for script.Scan() {
      text := []byte(script.Text())
      var article struct {
         Video struct {
            ContentURL string
         }
      }
      json.Unmarshal(text, &article)
      if article.Video.ContentURL != "" {
         return article.Video.ContentURL
      }
   }
   return ""
}

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

func main() {
   for _, test := range tests {
      fmt.Println(test)
      res, err := http.Get(test)
      if err != nil {
         panic(err)
      }
      defer res.Body.Close()
      doc, err := mech.Parse(res.Body)
      if err != nil {
         panic(err)
      }
      img := image(doc)
      if img != "" {
         fmt.Println(img)
      }
      aud := audio(doc)
      if aud != "" {
         fmt.Println(aud)
      }
      vid := video(doc)
      if vid != "" {
         fmt.Println(vid)
      }
      fmt.Println()
   }
}
