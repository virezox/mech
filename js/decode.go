package js

import (
   "bytes"
   "encoding/json"
   "github.com/tdewolff/parse/v2"
   "github.com/tdewolff/parse/v2/js"
   "io"
)

type Decoder struct {
   *js.AST
}

func NewDecoder(r io.Reader) (*Decoder, error) {
   ast, err := js.Parse(parse.NewInput(r))
   if err != nil {
      return nil, err
   }
   return &Decoder{ast}, nil
}

func (d *Decoder) Decode(be map[string]string) {
   for _, iStmt := range d.BlockStmt.List {
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
      be[bExpr.X.JS()] = bExpr.Y.JS()
   }
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
