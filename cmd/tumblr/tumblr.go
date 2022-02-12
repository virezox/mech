package main

import (
   "fmt"
   "github.com/89z/mech/tumblr"
   "net/http"
   "os"
   "path"
   "time"
)

func blogPost(link *tumblr.Permalink, info bool) error {
   post, err := link.BlogPost()
   if err != nil {
      return err
   }
   if info {
      fmt.Printf("%+v\n", post)
   } else {
      for _, elem := range post.Response.Timeline.Elements {
         fmt.Println("GET", elem.Video_URL)
         res, err := http.Get(elem.Video_URL)
         if err != nil {
            return err
         }
         defer res.Body.Close()
         file, err := os.Create(path.Base(elem.Video_URL))
         if err != nil {
            return err
         }
         defer file.Close()
         if _, err := file.ReadFrom(res.Body); err != nil {
            return err
         }
         time.Sleep(time.Second)
      }
   }
   return nil
}
