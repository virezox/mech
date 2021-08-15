package main

import (
   "fmt"
   "github.com/tdewolff/parse/v2"
   "github.com/tdewolff/parse/v2/js"
   "os"
)

func main() {
   f, err := os.Open("index.js")
   if err != nil {
      panic(err)
   }
   defer f.Close()
   ast, err := js.Parse(parse.NewInput(f))
   if err != nil {
      panic(err)
   }
   fmt.Printf("%+v\n", ast)
}
