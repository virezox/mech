package main

import (
   "fmt"
   "golang.org/x/net/html"
   "os"
)

func main() {
   f, err := os.Open("outer.html")
   if err != nil {
      panic(err)
   }
   defer f.Close()
   z := html.NewTokenizer(f)
   for {
      if z.Next() == html.ErrorToken {
         break
      }
      fmt.Println(z.Token())
   }
}
