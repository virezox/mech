package main

import (
   "encoding/json"
   "fmt"
   "github.com/89z/mech"
   "net/http"
)

const addr =
   "http://independent.co.uk" +
   "/news/world/americas/us-politics/b1873357.html"

func main() {
   fmt.Println(addr)
   res, err := http.Get(addr)
   if err != nil {
      panic(err)
   }
   defer res.Body.Close()
   doc, err := mech.Parse(res.Body)
   if err != nil {
      panic(err)
   }
   script := doc.ByAttr("type", "application/ld+json")
   for script.Scan() {
      text := []byte(script.Text())
      var article struct {
         Video struct { ContentURL string }
      }
      json.Unmarshal(text, &article)
      if article.Video.ContentURL != "" {
         fmt.Println(article.Video.ContentURL)
      }
   }
}
