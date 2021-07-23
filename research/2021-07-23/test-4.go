package main

import (
   "encoding/json"
   "fmt"
   "github.com/89z/mech"
   "net/http"
)

const addr =
   "http://nytimes.com/2021/07/14/podcasts/the-daily" +
   "/heat-wave-climate-change-pacific-northwest.html"

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
      var audio struct {
         ContentURL string
      }
      json.Unmarshal(text, &audio)
      if audio.ContentURL != "" {
         fmt.Println(audio.ContentURL)
      }
   }
}
