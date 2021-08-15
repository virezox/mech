package main

import (
   "encoding/json"
   "fmt"
   "github.com/89z/mech"
   "net/http"
   "net/http/httputil"
   "os"
)

func main() {
   if len(os.Args) != 2 {
      fmt.Println("media [URL]")
      return
   }
   req, err := http.NewRequest("GET", os.Args[1], nil)
   if err != nil {
      panic(err)
   }
   // instagram.com
   req.Header.Set("User-Agent", "Mozilla")
   d, err := httputil.DumpRequest(req, false)
   if err != nil {
      panic(err)
   }
   os.Stdout.Write(d)
   res, err := new(http.Client).Do(req)
   if err != nil {
      panic(err)
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      panic(res.Status)
   }
   s := mech.NewDecoder(res.Body)
   // This is going to kill audio and video if the page is missing og:image.
   // However that is unlikely, so we will cross that bridge when we come to it.
   s.AttrSelector("property", "og:image")
   fmt.Println(s.Attr("content"))
   // audio video
   for s.AttrSelector("type", "application/ld+json") {
      s.TextSelector()
      data := []byte(s.Data)
      // audio
      var audio struct {
         ContentURL string
      }
      json.Unmarshal(data, &audio)
      if audio.ContentURL != "" {
         fmt.Println(audio.ContentURL)
      }
      // video
      var article struct {
         Video struct {
            ContentURL string
         }
      }
      json.Unmarshal(data, &article)
      if article.Video.ContentURL != "" {
         fmt.Println(article.Video.ContentURL)
      }
   }
}
