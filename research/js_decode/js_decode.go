package main

import (
   "bytes"
   "encoding/json"
   "fmt"
   "github.com/tdewolff/parse/v2"
   "github.com/tdewolff/parse/v2/js"
   "io"
   "strings"
)

type quote struct{}

func (quote) Exit(js.INode) {}

func (q quote) Enter(n js.INode) js.IVisitor {
   pName, ok := n.(*js.PropertyName)
   if ok {
      b := bytes.Trim(pName.Literal.Data, `'"`)
      pName.Literal.Data, _ = json.Marshal(string(b))
   }
   return q
}

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
   m, err := decode(strings.NewReader(`d={ab:9,'cd':9,'c"d':9,"ef":9,"e'f":9}`))
   if err != nil {
      panic(err)
   }
   for k, v := range m {
      fmt.Printf("%q\n%v\n", k, v)
   }
}
