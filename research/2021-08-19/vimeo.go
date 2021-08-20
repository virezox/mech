package main

import (
   "fmt"
   "github.com/89z/mech/html"
   "github.com/89z/mech/js"
   "net/http"
)

func main() {
   addr := "https://vimeo.com/66531465"
   fmt.Println("GET", addr)
   res, err := http.Get(addr)
   if err != nil {
      panic(err)
   }
   defer res.Body.Close()
   lex := html.NewLexer(res.Body)
   lex.NextAttr("class", "app_banner_container")
   lex.NextTag("script")
   for k, v := range js.NewLexer(lex.Bytes()).Values() {
      fmt.Printf("%v %s\n", k, v)
   }
}
