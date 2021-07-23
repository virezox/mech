package main

import (
   "fmt"
   "github.com/89z/mech"
   "net/http"
)

func main() {
   res, err := http.Get("http://youtube.com/watch?v=LxK5Ocehj10")
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
