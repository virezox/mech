package main

import (
   "encoding/json"
   "fmt"
   "github.com/89z/mech"
   "net/http"
   "os"
)

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
   if len(os.Args) != 2 {
      fmt.Println("media [URL]")
      return
   }
   addr := os.Args[1]
   res, err := http.Get(addr)
   if err != nil {
      panic(err)
   }
   defer res.Body.Close()
   doc, err := mech.Parse(res.Body)
   if err != nil {
      panic(err)
   }
   if img := image(doc); img != "" {
      fmt.Println(img)
   }
   if aud := audio(doc); aud != "" {
      fmt.Println(aud)
   }
   if vid := video(doc); vid != "" {
      fmt.Println(vid)
   }
}
