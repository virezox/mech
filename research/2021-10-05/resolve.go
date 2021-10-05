package main

import (
   "fmt"
   "net/http"
   "net/url"
)

func session() ([]string, error) {
   res, err := http.Head("https://schnaussandmunk.bandcamp.com/track/amaris-2")
   if err != nil {
      return nil, err
   }
   var vals []string
   for _, cook := range res.Cookies() {
      if cook.Name == "session" {
         val, err := url.PathUnescape(cook.Value)
         if err != nil {
            return nil, err
         }
         vals = append(vals, val)
      }
   }
   return vals, nil
}

func trackID(session string) (int, error) {
   var id int
   _, err := fmt.Sscanf(session, "1\tr:[\"nilZ0t%vx", &id)
   if err != nil {
      return 0, err
   }
   return id, nil
}

func main() {
   id, err := trackID("1\tr:[\"nilZ0t2809477874x1633469972\"]\tt:1633469972")
   if err != nil {
      panic(err)
   }
   fmt.Println(id == 2809477874)
}
