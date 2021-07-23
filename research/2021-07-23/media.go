package main

import (
   "fmt"
   "github.com/89z/mech"
   "net/http"
)

func image(doc *mech.Node) string {
   img := doc.ByAttr("property", "og:image")
   img.Scan()
   return img.Attr("content")
}

var tests = []string{
   "http://hrp.org.uk" +
   "/tower-of-london/history-and-stories/tower-of-london-prison",
   "http://youtube.com/watch?v=LxK5Ocehj10",
}

func main() {
   for _, test := range tests {
      fmt.Println(test)
      res, err := http.Get(test)
      if err != nil {
         panic(err)
      }
      defer res.Body.Close()
      doc, err := mech.Parse(res.Body)
      if err != nil {
         panic(err)
      }
      fmt.Println(image(doc))
   }
}
