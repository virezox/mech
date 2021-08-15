package main

import (
   "fmt"
   "github.com/tdewolff/parse/v2"
   "github.com/tdewolff/parse/v2/js"
   "io"
   "os"
   "strconv"
   "strings"
)

const s = `
apps = [12, 31];
reference = {"month": 12, "day": 31};
`

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
      m[bExpr.X.JS()] = bExpr.Y.JS()
      /*
      var w walker
      js.Walk(w, s)
      */
   }
   return m, nil
}

type walker struct{}

func (w walker) Enter(n js.INode) js.IVisitor {
   pName, ok := n.(*js.PropertyName)
   if ok {
      s := string(pName.Literal.Data)
      pName.Literal.Data = strconv.AppendQuote(nil, s)
   }
   return w
}

func (walker) Exit(js.INode) {}
