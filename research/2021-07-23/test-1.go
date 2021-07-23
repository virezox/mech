package main

import (
   "fmt"
   "github.com/89z/mech"
   "net/http"
)

const addr =
   "http://hrp.org.uk" +
   "/tower-of-london/history-and-stories/tower-of-london-prison"

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
   img := doc.ByAttr("property", "og:image")
   img.Scan()
   fmt.Println(img.Attr("content"))
}
