package main

import (
   "fmt"
   "github.com/tdewolff/parse/v2"
   "github.com/tdewolff/parse/v2/js"
   "os"
)

func main() {
   b, err := os.ReadFile("index.js")
   if err != nil {
      panic(err)
   }
   l := js.NewLexer(parse.NewInputBytes(b))
   for {
      var key string
      for {
         tt, data := l.Next()
         if tt == js.ErrorToken {
            return
         } else if tt == js.IdentifierToken {
            key = string(data)
            break
         }
      }
      for {
         tt, data := l.Next()
         if tt == js.EqToken {
            break
         } else if tt != js.WhitespaceToken {
            key += string(data)
         }
      }
      fmt.Printf("%q\n", key)
      var val string
      for {
         tt, data := l.Next()
         if tt == js.SemicolonToken {
            break
         } else if tt != js.WhitespaceToken {
            val += string(data)
         }
      }
      fmt.Print(val, "\n\n")
   }
}
