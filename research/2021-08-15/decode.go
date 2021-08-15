package main

import (
   "fmt"
   "github.com/tdewolff/parse/v2"
   "github.com/tdewolff/parse/v2/js"
   "io"
   "strconv"
   "strings"
)

type walker struct{}

func (w walker) Enter(n js.INode) js.IVisitor {
   switch n := n.(type) {
   case *js.PropertyName:
      s := string(n.Literal.Data)
      n.Literal.Data = strconv.AppendQuote(nil, s)
   }
   return w
}

func (walker) Exit(js.INode) {}

func decode(r io.Reader) (map[string]string, error) {
   t, err := js.Parse(parse.NewInput(r))
   if err != nil {
      return nil, err
   }
   for _, s := range t.BlockStmt.List {
      var w walker
      js.Walk(w, s)
      fmt.Println(s)
   }
   return nil, nil
}

const s = `
apps = [12, 31];
reference = {"month": 12, "day": 31};
`

func main() {
   _, err := decode(strings.NewReader(s))
   if err != nil {
      panic(err)
   }
}
