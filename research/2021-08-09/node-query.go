package main

import (
   "fmt"
   "golang.org/x/net/html"
   "os"
)

func main() {
   f, err := os.Open("nyt.html")
   if err != nil {
      panic(err)
   }
   defer f.Close()
   z := html.NewTokenizer(f)
   for {
      if z.Next() == html.ErrorToken {
         break
      }
      t := z.Token()
      for _, a := range t.Attr {
         if a.Key == "type" && a.Val == "application/ld+json" {
            fmt.Println(t)
            z.Next()
            t := z.Token()
            fmt.Println(t.Data)
         }
      }
   }
}
