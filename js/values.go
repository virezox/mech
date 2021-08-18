package js

import (
   "bytes"
   "encoding/json"
   "github.com/tdewolff/parse/v2"
   "github.com/tdewolff/parse/v2/js"
)

type Values map[string]string

func Parse(b []byte) (Values, error) {
   ast, err := js.Parse(parse.NewInputBytes(b))
   if err != nil {
      return nil, err
   }
   v := make(Values)
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
      v[bExpr.X.JS()] = bExpr.Y.JS()
   }
   return v, nil
}

func (v Values) Get(key string) []byte {
   val := v[key]
   return []byte(val)
}

func (v Values) Has(key string) bool {
   _, ok := v[key]
   return ok
}

type quote struct{}

func (q quote) Enter(n js.INode) js.IVisitor {
   pName, ok := n.(*js.PropertyName)
   if ok {
      b := bytes.Trim(pName.Literal.Data, `'"`)
      pName.Literal.Data, _ = json.Marshal(string(b))
   }
   return q
}

func (quote) Exit(js.INode) {}
