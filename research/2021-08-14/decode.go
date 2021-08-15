package main

import (
   "fmt"
   "github.com/tdewolff/parse/v2"
   "github.com/tdewolff/parse/v2/js"
   "io"
   "os"
   "strconv"
)

type quote struct{}

func (q quote) Enter(n js.INode) js.IVisitor {
   pName, ok := n.(*js.PropertyName)
   if ok {
      s := string(pName.Literal.Data)
      pName.Literal.Data = strconv.AppendQuote(nil, s)
   }
   return q
}

func (quote) Exit(js.INode) {}

func decode(r io.Reader) (map[string]string, error) {
   ast, err := js.Parse(parse.NewInput(r))
   if err != nil {
      return nil, err
   }
   m := make(map[string]string)
   for _, iStmt := range ast.BlockStmt.List {
      eStmt, ok := iStmt.(*js.ExprStmt)
      if !ok {
         continue
      }
      bExpr, ok := eStmt.Value.(*js.BinaryExpr)
      if !ok {
         continue
      }
      var q quote
      js.Walk(q, bExpr.Y)
      m[bExpr.X.JS()] = bExpr.Y.JS()
   }
   return m, nil
}

func main() {
   f, err := os.Open("index.js")
   if err != nil {
      panic(err)
   }
   defer f.Close()
   m, err := decode(f)
   if err != nil {
      panic(err)
   }
   for k, v := range m {
      fmt.Printf("%q\n%v\n", k, v)
   }
}
