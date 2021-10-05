package resolve

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

func tralbum(session string) (rune, int, error) {
   var (
      typ rune
      id int
   )
   _, err := fmt.Sscanf(session, "1\tr:[\"nilZ0%c%vx", &typ, &id)
   if err != nil {
      return 0, 0, err
   }
   return typ, id, nil
}
