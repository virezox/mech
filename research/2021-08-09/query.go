package main

import (
   "fmt"
   "golang.org/x/net/html"
   "os"
)

func node() {
   f, err := os.Open("nyt.html")
   if err != nil {
      panic(err)
   }
   defer f.Close()
   n, err := html.Parse(f)
   if err != nil {
      panic(err)
   }
   todo := []*html.Node{n}
   for len(todo) > 0 {
      t := todo[0]
      todo = todo[1:]
      for c := t.FirstChild; c != nil; c = c.NextSibling {
         todo = append(todo, c)
      }
      for _, a := range t.Attr {
         if a.Key == "type" && a.Val == "application/ld+json" {
            fmt.Println(t.FirstChild.Data)
         }
      }
   }
}

func token() {
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
      for _, a := range z.Token().Attr {
         if a.Key == "type" && a.Val == "application/ld+json" {
            z.Next()
            fmt.Println(z.Token().Data)
         }
      }
   }
}
