package mech

import (
   "golang.org/x/net/html"
   "io"
)

func tag(t string) func(Node) bool {
   return func(n Node) bool {
      return n.Data == t
   }
}

type Node struct {
   *html.Node
}

func NewNode(r io.Reader) (Node, error) {
   d, err := html.Parse(r)
   if err != nil {
      return Node{}, err
   }
   return Node{d}, nil
}

func (n Node) Attr(key string) string {
   for _, a := range n.Node.Attr {
      if a.Key == key {
         return a.Val
      }
   }
   return ""
}

func (n Node) ByAttrAll(key, val string) []Node {
   return n.selector(true, func(n Node) bool {
      for _, a := range n.Node.Attr {
         if a.Key == key && a.Val == val {
            return true
         }
      }
      return false
   })
}

func (n Node) ByTag(name string) Node {
   for _, n := range n.selector(false, tag(name)) {
      return n
   }
   return Node{}
}

func (n Node) ByTagAll(name string) []Node {
   return n.selector(true, tag(name))
}

func (n Node) selector(all bool, f func(n Node) bool) []Node {
   var (
      in = []Node{n}
      out []Node
   )
   for len(in) > 0 {
      n := in[0]
      if f(n) {
         out = append(out, n)
         if ! all {
            break
         }
      }
      for c := n.FirstChild; c != nil; c = c.NextSibling {
         in = append(in, Node{c})
      }
      in = in[1:]
   }
   return out
}

func (n Node) Text() string {
   for _, n := range n.selector(false, func(n Node) bool {
      return n.Type == html.TextNode
   }) {
      return n.Data
   }
   return ""
}
