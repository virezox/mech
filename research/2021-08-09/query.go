package main

import (
   "fmt"
   "golang.org/x/net/html"
   "io"
   "os"
)

type scanner struct {
   *html.Tokenizer
}

func newScanner(r io.Reader) scanner {
   return scanner{
      html.NewTokenizer(r),
   }
}

func (s scanner) text() string {
   for {
      n := s.Next()
      if n == html.ErrorToken {
         break
      }
      if n == html.TextToken {
         return s.Token().Data
      }
   }
   return ""
}

func (s scanner) token(key, val string) *token {
   for {
      if s.Next() == html.ErrorToken {
         break
      }
      t := s.Token()
      for _, a := range t.Attr {
         if a.Key == key && a.Val == val {
            return &token{t}
         }
      }
   }
   return nil
}

type token struct {
   html.Token
}

func (t token) attr(key string) string {
   for _, a := range t.Attr {
      if a.Key == key {
         return a.Val
      }
   }
   return ""
}

func main() {
   f, err := os.Open("nyt.html")
   if err != nil {
      panic(err)
   }
   defer f.Close()
   s := newScanner(f)
   t := s.token("type", "application/ld+json")
   fmt.Println(t.attr("data-rh"))
   fmt.Println(s.text())
}
