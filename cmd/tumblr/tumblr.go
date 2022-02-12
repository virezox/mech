package main

import (
   "fmt"
   "github.com/89z/format"
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
      for _, element := range post.Response.Timeline.Elements {
         fmt.Println(element)
      }
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
         pro := format.NewProgress(res)
         if _, err := file.ReadFrom(pro); err != nil {
            return err
         }
         time.Sleep(time.Second)
      }
   }
   return nil
}
