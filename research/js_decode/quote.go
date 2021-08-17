package main

import (
   "fmt"
   "github.com/tdewolff/parse/v2"
   "github.com/tdewolff/parse/v2/js"
)

type quote struct{}

func (q quote) Enter(n js.INode) js.IVisitor {
   pName, ok := n.(*js.PropertyName)
   if ok {
      fmt.Println(pName)
      fmt.Println(pName.JS())
      fmt.Println(pName.Literal)
      fmt.Println(pName.Literal.JS())
      fmt.Println(pName.Computed)
   }
   return q
}

func (quote) Exit(js.INode) {}

func main() {
   ast, err := js.Parse(parse.NewInputString(`d={"month":12,"1day":31}`))
   if err != nil {
      panic(err)
   }
   for _, iStmt := range ast.BlockStmt.List {
      var q quote
      js.Walk(q, iStmt)
   }
}

/*
month
month
month
month
<nil>
"1day"
"1day"
"1day"
"1day"
<nil>
*/
