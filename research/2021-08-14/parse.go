package main

import (
   "fmt"
   "github.com/tdewolff/parse/v2"
   "github.com/tdewolff/parse/v2/js"
)

// js.ObjectExpr
func main() {
   ast, err := js.Parse(parse.NewInputString(`y={"year":2021}`))
   if err != nil {
      panic(err)
   }
   for _, is := range ast.BlockStmt.List {
      oe := is.(*js.ExprStmt).Value.(*js.BinaryExpr).Y.(*js.ObjectExpr)
      for _, p := range oe.List {
         // PropertyName
         fmt.Println(p.Name.Computed) // <nil>
         fmt.Println(p.Name.Literal) // year
         fmt.Println(p.Name.Literal.JS()) // year
      }
   }
}
